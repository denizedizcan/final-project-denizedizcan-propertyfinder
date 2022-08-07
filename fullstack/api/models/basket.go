package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Basket struct {
	BasketID    uint64        `gorm:"primary_key;auto_increment" json:"basket_id"`
	UserID      uint64        `gorm:"unique" json:"user_id"`
	Value       uint32        `gorm:"default:0" json:"value"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	BasketItems []BasketItems `gorm:"foreignKey:BasketID;references:BasketID"`
}

func FindAllBaskets(db *gorm.DB) ([]Basket, error) {

	var Baskets []Basket

	if result := db.Preload(clause.Associations).Find(&Baskets); result.Error != nil {
		return []Basket{}, result.Error
	}
	return Baskets, nil
}

func (b *BasketItems) FindUserBasketbyBasketitem(db *gorm.DB) (Basket, error) {
	var basket Basket

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).First(&basket); result.Error != nil {
		return Basket{}, result.Error
	}
	return basket, nil
}

func (b *User) FindUserBasketbyUser(db *gorm.DB) (Basket, error) {

	b, err := b.FindUserData(db)
	if err != nil {
		return Basket{}, err
	}
	var basket Basket

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.Basket.BasketID).First(&basket); result.Error != nil {
		return Basket{}, result.Error
	}
	return basket, nil
}

func (b *Basket) InsertBasket(db *gorm.DB) error {
	if result := db.Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BasketItems) UpdateBasketValue(db *gorm.DB) error {

	var basket Basket
	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).First(&basket); result.Error != nil {
		return result.Error
	}
	var val uint32 = 0
	for i := 0; len(basket.BasketItems) > i; i++ {
		val += basket.BasketItems[i].Value
	}

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).Update("value", val); result.Error != nil {
		return result.Error
	}

	return nil
}
