package models

import (
	"errors"
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

func InsertStocks(db *gorm.DB, s []Stock) error {
	if result := db.Create(&s); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateStocks(db *gorm.DB, s []Stock) error {
	for i := 0; i < len(s); i++ {
		if result := db.Model(&Stock{}).Where("sku = ?", s[i].Sku).Update("quantity", s[i].Quantity*3); result.Error != nil {
			return result.Error
		}
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
	if stock.Quantity < b.Quantity {
		return errors.New("not enough stock")
	}
	return nil
}

func (b *BasketItems) DropStock(db *gorm.DB) error {
	var stock Stock
	if result := db.Where("sku = ?", b.Sku).First(&stock); result.Error != nil {
		return result.Error
	}
	if stock.Quantity < b.Quantity {
		return errors.New("not enough stock")
	}
	if result := db.Model(&stock).Where("sku = ?", b.Sku).Update("quantity", stock.Quantity-b.Quantity); result.Error != nil {
		return result.Error
	}
	return nil
}
