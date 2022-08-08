package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Basket struct {
	BasketID    uint64        `gorm:"primary_key;auto_increment" json:"basket_id"`
	UserID      uint64        `gorm:"unique" json:"user_id"`
	Value       float64       `json:"value"`
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

func (b *Basket) InsertBasket(db *gorm.DB) error {
	if result := db.Create(&b); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Basket) UpdateBasketValueAfterOrder(db *gorm.DB) error {

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).Update("value", 0); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Basket) FindLastMonthOrders(db *gorm.DB) ([]Order, error) {
	today := time.Now()
	lastMonth := today.AddDate(0, -1, 0)

	var orders []Order
	if result := db.Model(Order{}).Preload(clause.Associations).Where("created_at BETWEEN ? AND ?", lastMonth, today).Find(&orders); result.Error != nil {
		return []Order{}, result.Error
	}
	return orders, nil

}

func (b *Basket) UpdateBasketValue(db *gorm.DB) error {

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.BasketID).Update("value", b.Value); result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Basket) FindUser(db *gorm.DB) (User, error) {
	var user User
	if result := db.Model(User{}).Preload(clause.Associations).Where("user_id = ?", b.UserID).First(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
