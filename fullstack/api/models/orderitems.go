package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItems struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	OrderNumber uint64    `json:"ordernumber"`
	Sku         uint64    `json:"sku"`
	Quantity    uint32    `json:"quantity"`
	Value       uint32    `json:"value"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func CreateOrderItems(db *gorm.DB, b []OrderItems) error {
	if result := db.Model(&OrderItems{}).Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteOrderItem(db *gorm.DB, o OrderItems) error {
	if result := db.Model(&OrderItems{}).Where("id = ?", o.ID).Delete(&o); result.Error != nil {
		return result.Error
	}
	return nil
}
