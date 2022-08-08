package models

import (
	"fmt"

	"gorm.io/gorm"
)

const discount_check float64 = 3000

func (u *Basket) CalculateLastForOrder(db *gorm.DB) float64 {
	var user User

	if u.UserID == 0 {
		return 0
	}
	user, err := u.FindUser(db)

	if err != nil {
		return 0
	}

	if len(user.Order) == 0 {
		return 0
	}
	var total float64 = 0

	for i := len(user.Order) - 1; i > len(user.Order)-4; i-- {
		if i == -1 {
			break
		}
		total += user.Order[i].Value
	}

	if total < discount_check {
		return 0
	}
	var discount float64 = 0

	for i := 0; i < len(u.BasketItems); i++ {
		vat, err := u.BasketItems[i].FindProductVat(db)

		if err != nil {
			return 0
		}
		if vat == 8 {
			discount += u.BasketItems[i].Value * 0.10
		}
		if vat == 18 {
			discount += u.BasketItems[i].Value * 0.15
		}
	}
	return discount
}

func (u *Basket) CheckIf4item(db *gorm.DB) float64 {
	var discount float64

	for i := 0; i < len(u.BasketItems); i++ {
		if u.BasketItems[i].Quantity > 3 {
			price, err := u.BasketItems[i].FindOnePrices(db)
			if err != nil {
				return 0
			}
			discount += float64(price.Value) * float64((u.BasketItems[i].Quantity - 3)) * 0.08
		}
	}
	return discount
}

func (u *Basket) CalculateLastMonthOrder(db *gorm.DB) float64 {
	orders, err := u.FindLastMonthOrders(db)

	if err != nil || len(orders) == 0 {
		return 0
	}
	var total float64 = 0

	for i := 0; i < len(orders); i++ {
		total += orders[i].Value
	}

	if total < discount_check {
		return 0
	}

	discount := u.Value * 0.1
	return discount
}

func (u *Basket) FindDiscount(db *gorm.DB) float64 {
	last_for_order := u.CalculateLastForOrder(db)
	check_if_4 := u.CheckIf4item(db)
	last_month := u.CalculateLastMonthOrder(db)
	fmt.Println("last_for_order: ", last_for_order)
	fmt.Println("check_if_4: ", check_if_4)
	fmt.Println("last_month: ", last_month)

	if last_for_order > check_if_4 && last_for_order > last_month {
		return last_for_order
	} else if check_if_4 > last_for_order && check_if_4 > last_month {
		return check_if_4
	} else {
		return last_month
	}
}
