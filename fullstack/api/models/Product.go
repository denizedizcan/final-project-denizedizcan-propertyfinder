package models

import (
	"time"

	"github.com/jinzhu/gorm"
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

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	Products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&Products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &Products, nil
}
