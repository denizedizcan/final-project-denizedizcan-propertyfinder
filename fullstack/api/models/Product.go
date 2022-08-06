package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Sku       uint64    `gorm:"primary_key;auto_increment" json:"Sku"`
	Name      string    `gorm:"size:255;not null;" json:"Name"`
	Base_Code string    `gorm:"size:255;not null;" json:"Base_Code"`
	Is_active bool      `gorm:"type:bool;default:false" json:"Is_active"`
	Language  string    `gorm:"size:255;not null;" json:"Language"`
	VAT       uint8     `gorm:"size:255;not null;" json:"VAT"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func FindAllProducts(db *gorm.DB) ([]Product, error) {

	var products []Product

	if result := db.Find(&products); result.Error != nil {
		return []Product{}, result.Error
	}
	return products, nil
}

func (p *[]Product) InsertProducts(db *gorm.DB) error{
	if result := db.
}