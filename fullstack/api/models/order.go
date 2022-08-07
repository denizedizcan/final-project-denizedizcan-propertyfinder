package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	OrderNumber uint64       `gorm:"primary_key;auto_increment" json:"ordernumber"`
	UserID      uint64       `gorm:"unique" json:"user_id"`
	Value       uint32       `json:"value"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	OrderItems  []OrderItems `gorm:"foreignKey:OrderNumber;references:OrderNumber"`
}

func (o *Order) CreateOrder(db *gorm.DB) (uint64, error) {
	if result := db.Create(&o); result.Error != nil {
		return 0, result.Error
	}
	return o.OrderNumber, nil
}

func (o *Order) FindOrder(db *gorm.DB) (*Order, error) {

	if o.OrderNumber != 0 {
		if result := db.Model(Order{}).Preload(clause.Associations).Where("order_number = ?", o.OrderNumber).First(&o); result.Error != nil {
			return &Order{}, result.Error
		}
	}
	return o, nil
}
