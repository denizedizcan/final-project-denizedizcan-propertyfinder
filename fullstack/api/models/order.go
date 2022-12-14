package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	OrderNumber uint64       `gorm:"primary_key;auto_increment" json:"ordernumber"`
	UserID      uint64       `json:"user_id"`
	Value       float64      `json:"value"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	OrderItems  []OrderItems `gorm:"foreignKey:OrderNumber;references:OrderNumber"`
}

// create order
func (o *Order) CreateOrder(db *gorm.DB) (uint64, error) {
	if result := db.Create(&o); result.Error != nil {
		return 0, result.Error
	}
	return o.OrderNumber, nil
}

// create old order
func (o *Order) CreateOrderOld(db *gorm.DB) (uint64, error) {
	o.CreatedAt = time.Now().AddDate(0, -3, 0)

	if result := db.Create(&o); result.Error != nil {
		return 0, result.Error
	}
	return o.OrderNumber, nil
}

// find order data and return it
func (o *Order) FindOrder(db *gorm.DB) (*Order, error) {

	if o.OrderNumber != 0 {
		if result := db.Model(Order{}).Preload(clause.Associations).Where("order_number = ?", o.OrderNumber).First(&o); result.Error != nil {
			return &Order{}, result.Error
		}
	}
	return o, nil
}

// delete order
func (o *Order) DeleteOrder(db *gorm.DB) error {
	var orderItems []OrderItems
	if result := db.Model(Order{}).Preload(clause.Associations).Where("order_number = ?", o.OrderNumber).Find(&orderItems); result.Error != nil {
		return result.Error
	}
	for i := 0; i < len(orderItems); i++ {
		if result := DeleteOrderItem(db, orderItems[i]); result != nil {
			return result
		}
	}
	if result := db.Model(Order{}).Preload(clause.Associations).Where("order_number = ?", o.OrderNumber).Delete(&o); result.Error != nil {
		return result.Error
	}
	return nil
}
