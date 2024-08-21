package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScriptSig struct {
	Asm *string `json:"asm"`
	Hex *string `json:"hex"`
}

type Vin struct {
	ID            string      `gorm:"type:char(36);primary_key;"`
	TransactionID string      `json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignkey:TransactionID"`
	Txid          *string     `gorm:"index"`
	Vout          int         `json:"vout"`
	ScriptSig     ScriptSig   `gorm:"embedded;embeddedPrefix:script_sig_" json:"scriptSig"`
	Sequence      int         `json:"sequence"`
	Coinbase      string      `json:"coinbase,omitempty"`
}

func (v *Vin) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New().String()
	v.ScriptSig.Asm = SetNilIfEmpty(v.ScriptSig.Asm)
	v.ScriptSig.Hex = SetNilIfEmpty(v.ScriptSig.Hex)
	return
}
