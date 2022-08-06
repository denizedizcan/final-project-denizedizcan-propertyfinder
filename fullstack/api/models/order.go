package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	OrderNumber uint64    `gorm:"primary_key;auto_increment" json:"Sku"`
	OrderedBy   string    `gorm:"size:255;not null;" json:"OrderedBy"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *User) FindAllOrders(db *gorm.DB) (*[]Order, error) {
	var err error
	Orders := []Order{}
	err = db.Debug().Model(&Order{}).Limit(100).Find(&Orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &Orders, nil
}
