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
func (u *User) FindUserBasket(db *gorm.DB) (Basket, error) {

	var Baskets Basket

	if result := db.Preload(clause.Associations).Find(&Baskets); result.Error != nil {
		return Basket{}, result.Error
	}
	return Baskets, nil
}

func (b *Basket) InsertBasket(db *gorm.DB) error {
	if result := db.Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}
