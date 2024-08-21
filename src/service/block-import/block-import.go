package blockimport

import (
	"encoding/json"
	"fmt"

	dbconfig "github.com/nodlandhodl/bitcoin-analytics-backend/src/db-config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/entities"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/service/bitcoind"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var db *gorm.DB = dbconfig.ConnectDB()

type ImportOptions struct {
	BlockHash string
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ImportBlocksToDb(options ImportOptions) {
	bitcoindService, err := bitcoind.NewBitcoindService()
	if err != nil {
		panic(err)
	}

	var hash string
	if len(options.BlockHash) == 0 {
		var count int64
		if err := db.Model(&entities.Block{}).Count(&count).Error; err != nil {
			panic(err)
		}

		var errd error
		hash, errd = bitcoindService.GetBlockHash(int(count))
		if errd != nil {
			panic(errd)
		}
	} else {
		hash = options.BlockHash
	}

	blockd, erd := bitcoindService.GetBlock(hash)
	if erd != nil {
		panic(erd.Error)
	}

	fmt.Println(blockd)

	newBlock := &entities.Block{
		Height:            blockd.Height,
		Hash:              blockd.Hash,
		Version:           blockd.Version,
		VersionHex:        blockd.VersionHex,
		MerkleRoot:        blockd.MerkleRoot,
		Time:              blockd.Time,
		Mediantime:        blockd.MedianTime,
		Nonce:             blockd.Nonce,
		Bits:              blockd.Bits,
		Difficulty:        blockd.Difficulty,
		Chainwork:         blockd.Chainwork,
		NTx:               blockd.NTx,
		PreviousBlockHash: blockd.PreviousBlockHash,
		NextBlockHash:     blockd.NextBlockHash,
		StrippedSize:      blockd.StrippedSize,
		Size:              blockd.Size,
		Weight:            blockd.Weight,
		Tx:                datatypes.JSON{},
	}

	txJSON, err := json.Marshal(blockd.Tx)
	if err != nil {
		panic(err)
	}
	newBlock.Tx = datatypes.JSON(txJSON)

	result := db.Create(&newBlock)
	if result.Error != nil {
		panic(result.Error)
	}

	for _, tx := range blockd.Tx {
		transactiond, err4 := bitcoindService.GetRawTransaction(tx, true)
		if err4 != nil {
			panic(err4)
		}
		fmt.Println(transactiond)
		newTransaction := &entities.Transaction{
			BlockID:       newBlock.ID,
			Hex:           transactiond.Hex,
			Confirmations: transactiond.Confirmations,
			Time:          transactiond.Time,
			Blockhash:     transactiond.Blockhash,
			Txid:          transactiond.Txid,
			Hash:          transactiond.Hash,
			Size:          transactiond.Size,
			Vsize:         transactiond.Vsize,
			Version:       transactiond.Version,
			Locktime:      transactiond.Locktime,
			Weight:        transactiond.Weight,
		}

		res := db.Create(&newTransaction)
		if res.Error != nil {
			panic(res.Error)
		}

		for _, vout := range transactiond.Vout {
			newVout := entities.Vout{
				Txid:          stringToPointer(transactiond.Txid),
				TransactionID: newTransaction.ID,
				Value:         vout.Value,
				N:             vout.N,
				ScriptPubKey: entities.ScriptPubKey{
					Asm:     stringToPointer(vout.ScriptPubKey.Asm),
					Hex:     stringToPointer(vout.ScriptPubKey.Hex),
					Type:    stringToPointer(vout.ScriptPubKey.Type),
					Address: stringToPointer(vout.ScriptPubKey.Address),
				},
			}
			res := db.Create(&newVout)
			if res.Error != nil {
				panic(res.Error)
			}
		}

		for _, vin := range transactiond.Vin {
			newVin := entities.Vin{
				Txid:          stringToPointer(vin.Txid),
				TransactionID: newTransaction.ID,
				Vout:          vin.Vout,
				Coinbase:      vin.Coinbase,
				ScriptSig: entities.ScriptSig{
					Asm: stringToPointer(vin.ScriptSig.Asm),
					Hex: stringToPointer(vin.ScriptSig.Hex),
				},
				Sequence: vin.Sequence,
			}
			res := db.Create(&newVin)
			if res.Error != nil {
				panic(res.Error)
			}
		}
	}

	if blockd.NextBlockHash != "" {
		ImportBlocksToDb(ImportOptions{BlockHash: options.BlockHash})
	}
}
