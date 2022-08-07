package models

import (
	"time"

	"gorm.io/gorm"
)

type Price struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Value     uint32    `gorm:"default:0" json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func InsertPrices(db *gorm.DB, p []Price) error {
	if result := db.Create(&p); result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllPrices(db *gorm.DB) ([]Price, error) {

	var prices []Price

	if result := db.Find(&prices); result.Error != nil {
		return []Price{}, result.Error
	}
	return prices, nil
}

func (p *BasketItems) FindOnePrices(db *gorm.DB) (Price, error) {

	var prices Price
	if result := db.Where("sku = ?", p.Sku).Find(&prices); result.Error != nil {
		return Price{}, result.Error
	}
	return prices, nil
}
