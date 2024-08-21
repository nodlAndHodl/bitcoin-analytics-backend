package entities

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Block struct {
	ID                string `gorm:"type:char(36);primary_key;"`
	Hash              string `gorm:"not null"`
	Height            int    `gorm:"not null"`
	Version           int
	VersionHex        string
	MerkleRoot        string
	Time              int
	Mediantime        int
	Nonce             int
	Bits              string
	Difficulty        float64
	Chainwork         string
	NTx               int
	PreviousBlockHash string
	NextBlockHash     string
	Tx                datatypes.JSON `gorm:"type:json"`
	StrippedSize      int
	Size              int
	Weight            int
}

func (b *Block) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
