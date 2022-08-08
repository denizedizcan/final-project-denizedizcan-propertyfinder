package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BasketItems struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	BasketID  uint64    `json:"basket_id"`
	Sku       uint64    `gorm:"unique" json:"sku"`
	Quantity  uint32    `json:"quantity"`
	Value     float64   `json:"value"`
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

func (b *BasketItems) UpdateBasketItem(db *gorm.DB) error {
	if b.Quantity <= 0 {
		return errors.New("wrong update quantity")
	} else if result := db.Model(&b).Where("sku = ?", b.Sku).Update("quantity", b.Quantity); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Basket) FindBasketItems(db *gorm.DB) ([]BasketItems, error) {
	var basketItems []BasketItems

	if result := db.Model(&BasketItems{}).Where("basket_id = ?", b.BasketID).Find(&basketItems); result.Error != nil {
		return []BasketItems{}, result.Error
	}
	return basketItems, nil
}

func (b *BasketItems) UpdateBasketItemsValue(db *gorm.DB) error {
	var prices Price
	if result := db.Model(&prices).Where("sku = ?", b.Sku).First(&prices); result.Error != nil {
		return result.Error
	}
	new := float64(b.Quantity) * prices.Value
	if result := db.Model(&b).Where("sku = ?", b.Sku).Update("value", new); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BasketItems) DeleteBasketItem(db *gorm.DB) error {
	if result := db.Model(BasketItems{}).Where("sku = ?", b.Sku).Delete(&b); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BasketItems) FindProductVat(db *gorm.DB) (uint8, error) {
	var product Product

	if result := db.Model(Product{}).Where("sku = ?", b.Sku).First(&product); result.Error != nil {
		return 0, result.Error
	}

	return product.VAT, nil
}

func (b *BasketItems) UpdateBasketValue(db *gorm.DB) error {

	var basket Basket
	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).First(&basket); result.Error != nil {
		return result.Error
	}
	var val float64 = 0
	for i := 0; len(basket.BasketItems) > i; i++ {
		val += basket.BasketItems[i].Value
	}

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).Update("value", val); result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BasketItems) CheckStock(db *gorm.DB) error {

	var stock Stock
	if result := db.Model(&stock).Where("sku = ?", b.Sku).First(&stock); result.Error != nil {
		return result.Error
	}
	if stock.Quantity < b.Quantity {
		return errors.New("not enough stock")
	}
	return nil
}

func (b *BasketItems) DropStock(db *gorm.DB) error {
	var stock Stock
	if result := db.Model(&stock).Where("sku = ?", b.Sku).First(&stock); result.Error != nil {
		return result.Error
	}
	if stock.Quantity < b.Quantity {
		return errors.New("not enough stock")
	}
	new_stock := stock.Quantity - b.Quantity
	if result := db.Model(&stock).Where("sku = ?", b.Sku).Update("quantity", new_stock); result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *BasketItems) FindOnePrices(db *gorm.DB) (Price, error) {

	var prices Price
	if result := db.Model(&prices).Where("sku = ?", p.Sku).Find(&prices); result.Error != nil {
		return Price{}, result.Error
	}
	return prices, nil
}

func (b *BasketItems) FindUserBasketbyBasketitem(db *gorm.DB) (Basket, error) {
	var basket Basket

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).First(&basket); result.Error != nil {
		return Basket{}, result.Error
	}
	return basket, nil
}
