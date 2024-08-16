package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID            string `gorm:"type:char(36);primary_key;"`
	Txid          string `gorm:"not null"`
	Hash          string
	Version       int
	Size          int
	Vsize         int
	Weight        int
	Locktime      int
	Vin           string
	Vout          string
	Hex           string
	Blockhash     string
	Confirmations int
	Time          int
	BlockID       string
	Block         Block `gorm:"foreignkey:BlockID"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
