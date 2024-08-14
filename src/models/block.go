package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Block struct {
	ID                string `gorm:"type:char(36);primary_key;"`
	Hash              string `gorm:"not null"`
	Height            int    `gorm:"not null"`
	Version           int
	VersionHex        string
	Merkleroot        string
	Time              int
	Mediantime        int
	Nonce             int
	Bits              string
	Difficulty        float64
	Chainwork         string
	NTx               int
	Previousblockhash string
	Nextblockhash     string
	Tx                string
	Strippedsize      int
	Size              int
	Weight            int
	Transactions      []Transaction `gorm:"foreignkey:BlockID"`
}

func (b *Block) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
