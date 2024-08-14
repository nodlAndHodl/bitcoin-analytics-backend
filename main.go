package main

import (
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/models"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	//run all migrations
	// bitcoindService, err := bitcoind.NewBitcoindService()
	// if err != nil {
	// 	panic(err)
	// }
	// blockhash, err := bitcoindService.GetBlockHash(0)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Block: %v\n", blockhash)
	// block, err := bitcoindService.GetBlock(blockhash)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Block: %v\n", block)
	models.AutoMigrate()
	//run all routes
	routes.Routes()
}
