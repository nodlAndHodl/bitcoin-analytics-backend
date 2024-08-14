package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8&loc=UTC", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
