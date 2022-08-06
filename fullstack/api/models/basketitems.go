package models

import (
	"time"
)

type BasketItems struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	BasketID  uint64    `gorm:"unique" json:"basket_id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `json:"quantity"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
