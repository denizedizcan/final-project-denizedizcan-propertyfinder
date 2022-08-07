package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItems struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	OrderNumber uint64    `json:"ordernumber"`
	Sku         uint64    `gorm:"unique" json:"sku"`
	Quantity    uint32    `json:"quantity"`
	Value       uint32    `json:"value"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func CreateOrderItems(db *gorm.DB, b []OrderItems) error {
	if result := db.Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}
