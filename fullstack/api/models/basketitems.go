package models

import (
	"time"

	"gorm.io/gorm"
)

type BasketItems struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	BasketID  uint64    `gorm:"unique" json:"basket_id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `json:"quantity"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func InsertBasketItems(db *gorm.DB, p []BasketItems) error {

	if result := db.Create(&p); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BasketItems) InsertOneBasketItem(db *gorm.DB) error {

	if result := db.Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BasketItems) UpdateOrDeleteBasketItem(db *gorm.DB) error {
	if b.Quantity <= 0 {
		if result := db.Delete(&b); result.Error != nil {
			return result.Error
		}
	} else if result := db.Model(&b).Update("quantity", b.Quantity); result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllBasketItems(db *gorm.DB) (BasketItems, error) {
	var BasketItems BasketItems

	if result := db.Find(&BasketItems); result.Error != nil {
		return BasketItems, result.Error
	}
	return BasketItems, nil
}
