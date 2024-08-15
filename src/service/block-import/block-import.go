package blockimport

import (
	"fmt"

	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/models"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/service/bitcoind"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

func ImportBlocksToDb() {

	bitcoindService, err := bitcoind.NewBitcoindService()
	if err != nil {
		panic(err)
	}

	var block *models.Block
	// Get the latest block height from the database
	er := db.Find(&block).Order("height desc").First(&block)
	if er.Error != nil {
		if er.Error.Error() != "record not found" {
			panic(er.Error)
		}
	}

	if block == nil || block.Height == 0 {
		block = &models.Block{Height: 0}
	}

	savedBlockHeight := block.Height

	var blockHash, errd = bitcoindService.GetBlockHash(savedBlockHeight)
	if errd != nil {
		panic(errd)
	}
	var blockd, erd = bitcoindService.GetBlock(blockHash)
	if erd != nil {
		panic(er.Error)
	}

	fmt.Println(blockd)

	newBlock := models.Block{
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
		//Tx:                blockd.Tx,
		StrippedSize: blockd.StrippedSize,
		Size:         blockd.Size,
		Weight:       blockd.Weight,
		//Transactions: nil,
	}
	// Get the latest block height from the bitcoind service
	// Loop through the blocks from the latest block height in the database to the latest block height from the bitcoind service
	// Get the block hash from the bitcoind service
	// Get the block from the bitcoind service

	// Save the block to the database
	// db.Create(&newBlock)
	result := db.Create(&newBlock)
	if result.Error != nil {
		panic(result.Error)
	}

}
