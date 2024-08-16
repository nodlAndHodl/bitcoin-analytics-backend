package blockimport

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/models"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/service/bitcoind"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

func ImportBlocksToDb() {

	bitcoindService, err := bitcoind.NewBitcoindService()
	if err != nil {
		panic(err)
	}

	var count int64

	if err := db.Model(&models.Block{}).Count(&count).Error; err != nil {
		log.Fatal(err)
	}

	var blockHash, errd = bitcoindService.GetBlockHash(int(count))
	if errd != nil {
		panic(errd)
	}
	var blockd, erd = bitcoindService.GetBlock(blockHash)
	if erd != nil {
		panic(erd.Error)
	}

	fmt.Println(blockd)

	newBlock := &models.Block{
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
	// Get the latest block height from the bitcoind service
	// Loop through the blocks from the latest block height in the database to the latest block height from the bitcoind service
	// Get the block hash from the bitcoind service
	// Get the block from the bitcoind service

	result := db.Create(&newBlock)
	if result.Error != nil {
		panic(result.Error)
	}

	var trasactiond, err4 = bitcoindService.GetRawTransaction(blockd.Tx[0], true)
	if err4 != nil {
		panic(err4)
	}
	fmt.Println(trasactiond)
	var newTransaction = &models.Transaction{
		BlockID:       newBlock.ID,
		Hex:           trasactiond.Hex,
		Confirmations: trasactiond.Confirmations,
		Time:          trasactiond.Time,
		Blockhash:     trasactiond.Blockhash,
		Txid:          trasactiond.Txid,
		Hash:          trasactiond.Hash,
		Size:          trasactiond.Size,
		Vsize:         trasactiond.Vsize,
		Version:       trasactiond.Version,
		Locktime:      trasactiond.Locktime,
		Weight:        trasactiond.Weight,
	}

	voutJSON, err := json.Marshal(trasactiond.Vout)
	if err != nil {
		panic(err)
	}
	newTransaction.Vout = string(voutJSON)

	vinJSON, err := json.Marshal(trasactiond.Vin)
	if err != nil {
		panic(err)
	}

	newTransaction.Vin = string(vinJSON)

	res := db.Create(&newTransaction)
	if res.Error != nil {
		panic(res.Error)
	}

}
