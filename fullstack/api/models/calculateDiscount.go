package models

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const discount_amount uint32 = 3000

func (u *Basket) CalculateLastForOrder(db *gorm.DB) error {
	var user User
	if u.UserID != 0 {
		if result := db.Model(User{}).Preload(clause.Associations).Where("email = ?", u.UserID).First(&user); result.Error != nil {
			return result.Error
		}
	} else {
		return errors.New("invalid user id")
	}

	if len(user.Order) < 4 {
		return errors.New("not enough order")
	}
	var total uint32 = 0
	for i := len(user.Order) - 1; i > len(user.Order)-4; i-- {
		total += user.Order[i].Value
	}

	if total < discount_amount {
		return errors.New("not enough order payment")
	}
	return nil
}
