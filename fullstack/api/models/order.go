package models

import (
	"time"
)

type Order struct {
	OrderNumber uint64       `gorm:"primary_key;auto_increment" json:"ordernumber"`
	UserID      uint64       `gorm:"unique" json:"user_id"`
	Value       uint32       `gorm:"default:0" json:"value"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	OrderItems  []OrderItems `gorm:"foreignKey:OrderNumber;references:OrderNumber"`
}
