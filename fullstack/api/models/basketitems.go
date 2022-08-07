package models

import (
	"time"

	"gorm.io/gorm"
)

type BasketItems struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	BasketID  uint64    `gorm:"unique" json:"basket_id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `gorm:"default:0" json:"quantity"`
	Value     uint32    `gorm:"default:0" json:"value"`
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

func (b *Basket) FindBasketItems(db *gorm.DB) ([]BasketItems, error) {
	var basketItems []BasketItems

	if result := db.Where("basket_id = ?", b.BasketID).Find(&basketItems); result.Error != nil {
		return []BasketItems{}, result.Error
	}
	return basketItems, nil
}

func (b *BasketItems) UpdateBasketItemsValue(db *gorm.DB) error {
	var prices Price
	if result := db.Where("sku = ?", b.Sku).First(&prices); result.Error != nil {
		return result.Error
	}
	if result := db.Model(&b).Update("value", b.Quantity*prices.Value); result.Error != nil {
		return result.Error
	}
	return nil
}
