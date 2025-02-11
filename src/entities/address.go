package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID             string `gorm:"type:char(36);primary_key;"`
	Address        string `gorm:"not null"`
	TotalAmountIn  float64
	TotalAmountOut float64
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
