package entities

import (
	dbconfig "github.com/nodlandhodl/bitcoin-analytics-backend/src/db-config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = dbconfig.ConnectDB()
)

func AutoMigrate() {
	db.AutoMigrate(&Block{}, &Transaction{})
	db.AutoMigrate(&Vin{}, &Vout{})
}
