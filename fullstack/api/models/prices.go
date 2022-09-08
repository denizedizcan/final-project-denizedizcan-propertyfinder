package models

import (
	"time"

	"gorm.io/gorm"
)

type Price struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Value     float64   `json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
// insert price list
func InsertPrices(db *gorm.DB, p []Price) error {
	if result := db.Create(&p); result.Error != nil {
		return result.Error
	}
	return nil
}

// get price list
func FindAllPrices(db *gorm.DB) ([]Price, error) {

	var prices []Price

	if result := db.Find(&prices); result.Error != nil {
		return []Price{}, result.Error
	}
	return prices, nil
}
