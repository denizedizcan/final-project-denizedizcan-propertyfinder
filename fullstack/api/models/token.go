package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Token struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Token     string    `gorm:"size:255;not null;unique" json:"Token"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func RandStringBytes(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (u *User) CreateToken(db *gorm.DB) error {
	token := "Bearer " + RandStringBytes(20)
	var t Token
	t.Token = token
	t.User = *u

	if result := db.Create(&t); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) ValidateToken(db *gorm.DB, token string) error {
	var t Token
	if result := db.Joins("tokens").Joins("users").Where("name = ?", u.Name).Take(&t); result.Error != nil {
		return result.Error
	}
	fmt.Print(db.Joins("tokens").Joins("users").First("name = ?", u.Name).Take(&t).RowsAffected)
	if t.Token != token {
		return errors.New("auth fail")
	}
	return nil
}
