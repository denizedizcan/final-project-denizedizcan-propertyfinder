package models

import (
	"time"
)

type Basket struct {
	BasketID    uint64        `gorm:"primary_key;auto_increment" json:"basket_id"`
	UserID      uint64        `gorm:"unique" json:"user_id"`
	Value       uint32        `json:"value"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	BasketItems []BasketItems `gorm:"foreignKey:BasketID;references:BasketID"`
}
