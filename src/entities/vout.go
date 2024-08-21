package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Hex     string `json:"hex"`
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
}

type Vout struct {
	ID           string       `gorm:"type:char(36);primary_key;"`
	Txid         string       `gorm:"index"`
	Transaction  Transaction  `gorm:"foreignkey:Txid"`
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `gorm:"embedded" json:"scriptPubKey"`
}

func (v *Vout) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New().String()
	return
}
