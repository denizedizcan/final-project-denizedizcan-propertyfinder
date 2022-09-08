package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `json:"quantity"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// insert stock list to db
func InsertStocks(db *gorm.DB, s []Stock) error {
	if result := db.Create(&s); result.Error != nil {
		return result.Error
	}
	return nil
}

// update stocks with slice
func UpdateStocks(db *gorm.DB, s []Stock) error {
	for i := 0; i < len(s); i++ {
		if result := db.Model(&Stock{}).Where("sku = ?", s[i].Sku).Update("quantity", s[i].Quantity); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// get stock list
func FindAllStocks(db *gorm.DB) ([]Stock, error) {

	var stocks []Stock

	if result := db.Model(&Stock{}).Find(&stocks); result.Error != nil {
		return []Stock{}, result.Error
	}
	return stocks, nil
}
