package main

import (
	dbconfig "github.com/nodlandhodl/bitcoin-analytics-backend/src/db-config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/entities"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/routes"
	blockimport "github.com/nodlandhodl/bitcoin-analytics-backend/src/service/block-import"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = dbconfig.ConnectDB()
)

func main() {
	defer dbconfig.DisconnectDB(db)
	entities.AutoMigrate()
	blockimport.ImportBlocksToDb(blockimport.ImportOptions{})
	routes.Routes()
}
