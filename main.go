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
	models.AutoMigrate()
	//run all routes
	routes.Routes()
}
