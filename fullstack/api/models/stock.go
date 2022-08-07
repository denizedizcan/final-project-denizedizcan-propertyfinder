package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `gorm:"default:0" json:"stock"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func InsertStocks(db *gorm.DB, s []Stock) error {
	if result := db.Create(&s); result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllStocks(db *gorm.DB) ([]Stock, error) {

	var stocks []Stock

	if result := db.Find(&stocks); result.Error != nil {
		return []Stock{}, result.Error
	}
	return stocks, nil
}

func (b *BasketItems) CheckStock(db *gorm.DB) error {

	var stock Stock
	if result := db.Where("sku = ?", b.Sku).First(&stock); result.Error != nil {
		return result.Error
	}
	if stock.Quantity == 0 {
		return errors.New("no stock left")
	}
	return nil
}
