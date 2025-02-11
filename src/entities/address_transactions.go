package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressTransactions struct {
	ID            string `gorm:"type:char(36);primary_key;"`
	AddressID     string `gorm:"not null;index:idx_address_id"`
	TransactionID string `gorm:"not null;index:idx_transaction_id"`
}

func (at *AddressTransactions) BeforeCreate(tx *gorm.DB) (err error) {
	at.ID = uuid.New().String()
	return
}
