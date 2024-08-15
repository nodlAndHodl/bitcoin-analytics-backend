package main

import (
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/models"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/routes"
	blockimport "github.com/nodlandhodl/bitcoin-analytics-backend/src/service/block-import"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	models.AutoMigrate()
	blockimport.ImportBlocksToDb()
	routes.Routes()
}
