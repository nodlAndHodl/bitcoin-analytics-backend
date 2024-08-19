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
	Vout          string
	Hex           string
	Blockhash     string
	Confirmations int
	Time          int
	BlockID       string
	Block         Block `gorm:"foreignkey:BlockID"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type ScriptPubKey struct {
	Asm    string `json:"asm"`
	Hex    string `json:"hex"`
	Desc   string `json:"desc,omitempty"`
	Type   string `json:"type"`
	Adress string `json:"address,omitempty"`
}

type Vin struct {
	ID        string    `gorm:"type:char(36);primary_key;"`
	Txid      string    `json:"txid" gorm:"index"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `gorm:"embedded" json:"scriptSig"`
	Sequence  int       `json:"sequence"`
	Coinbase  string    `json:"coinbase,omitempty"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (v *Vin) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New().String()
	return
}
