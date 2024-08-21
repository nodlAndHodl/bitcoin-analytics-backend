package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vin struct {
	ID          string      `gorm:"type:char(36);primary_key;"`
	Txid        string      `gorm:"index"`
	Transaction Transaction `gorm:"foreignkey:Txid"`
	Vout        int         `json:"vout"`
	ScriptSig   ScriptSig   `gorm:"embedded" json:"scriptSig"`
	Sequence    int         `json:"sequence"`
	Coinbase    string      `json:"coinbase,omitempty"`
}

func (v *Vin) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New().String()
	return
}
