package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	Sku         uint64        `gorm:"primary_key;auto_increment" json:"sku"`
	Name        string        `json:"name"`
	Base_Code   string        `json:"base_code"`
	Is_active   bool          `json:"is_active"`
	Language    string        `json:"language"`
	VAT         uint8         `json:"vat"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Price       Price         `gorm:"foreignKey:Sku;references:sku"`
	Stock       Stock         `gorm:"foreignKey:Sku;references:sku"`
	OrderItems  []OrderItems  `gorm:"foreignKey:Sku;references:sku"`
	BasketItems []BasketItems `gorm:"foreignKey:Sku;references:Sku"`
}

func FindAllProducts(db *gorm.DB) ([]Product, error) {

	var products []Product

	if result := db.Preload(clause.Associations).Find(&products); result.Error != nil {
		return []Product{}, result.Error
	}
	return products, nil
}

func InsertProducts(db *gorm.DB, p []Product) error {
	if result := db.Create(&p); result.Error != nil {
		return result.Error
	}
	return nil
}

func InsertOneProduct(db *gorm.DB, p Product) error {
	if result := db.Create(&p); result.Error != nil {
		return result.Error
	}
	return nil
}
