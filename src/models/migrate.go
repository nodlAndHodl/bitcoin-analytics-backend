package models

import (
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func AutoMigrate() {
	db.AutoMigrate(&Block{}, &Transaction{}, &Vin{})
}
