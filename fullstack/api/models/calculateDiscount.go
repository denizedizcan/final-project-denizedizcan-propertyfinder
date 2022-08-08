package models

import (
	"gorm.io/gorm"
)

// discont value to check
const discount_check float64 = 3000

// check last 4 order and check if orders total value is higher then discount check if higher return the discount amount to do
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

// check if basket has 4 or more same item if it is the case make a discount and return the discount value
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

//check for last month orders total value if its higher then discount check return the value for the discount
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

// check for all discount avaible and pick the highest discount value
func (u *Basket) FindDiscount(db *gorm.DB) float64 {
	last_for_order := u.CalculateLastForOrder(db)
	check_if_4 := u.CheckIf4item(db)
	last_month := u.CalculateLastMonthOrder(db)

	if last_for_order > check_if_4 && last_for_order > last_month {
		return last_for_order
	} else if check_if_4 > last_for_order && check_if_4 > last_month {
		return check_if_4
	} else {
		return last_month
	}
}
