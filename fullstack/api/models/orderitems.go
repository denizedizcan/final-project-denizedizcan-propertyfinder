package models

import (
	"time"
)

type OrderItems struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	OrderNumber uint64    `gorm:"unique" json:"ordernumber"`
	Sku         uint64    `gorm:"unique" json:"sku"`
	Quantity    uint8     `gorm:"default:0" json:"quantity"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
