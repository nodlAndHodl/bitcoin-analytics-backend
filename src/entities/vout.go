package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScriptPubKey struct {
	Asm     *string `json:"asm"`
	Hex     *string `json:"hex"`
	Type    *string `json:"type"`
	Address *string `json:"address,omitempty"`
}

type Vout struct {
	ID            string       `gorm:"type:char(36);primary_key;"`
	TransactionID string       `json:"transaction_id"`
	Transaction   Transaction  `gorm:"foreignkey:TransactionID"`
	Txid          *string      `gorm:"index"`
	Value         float64      `json:"value"`
	N             int          `json:"n"`
	ScriptPubKey  ScriptPubKey `gorm:"embedded;embeddedPrefix:script_pub_key_" json:"scriptPubKey"`
}

func (v *Vout) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New().String()
	v.ScriptPubKey.Asm = SetNilIfEmpty(v.ScriptPubKey.Asm)
	v.ScriptPubKey.Hex = SetNilIfEmpty(v.ScriptPubKey.Hex)
	v.ScriptPubKey.Type = SetNilIfEmpty(v.ScriptPubKey.Type)
	v.ScriptPubKey.Address = SetNilIfEmpty(v.ScriptPubKey.Address)
	return
}

func SetNilIfEmpty(field *string) *string {
	if field != nil && *field == "" {
		return nil
	}
	return field
}
