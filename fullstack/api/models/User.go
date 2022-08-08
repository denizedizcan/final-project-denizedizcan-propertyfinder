package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/badoux/checkmail"
)

//user struct fields used in db
type User struct {
	UserID    uint64    `gorm:"primary_key;auto_increment" json:"user_id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Basket    Basket    `gorm:"foreignKey:UserID;references:UserID"`
	Order     []Order   `gorm:"foreignKey:UserID;references:user_id"`
}

//user prepare values to insert or update
func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.CreatedAt = time.Now()
}

// validate values
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "create":
		if u.Name == "" {
			return errors.New("required Name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("required Name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

//Crate user
func (u *User) SaveUser(db *gorm.DB) error {

	if result := db.Create(&u); result.Error != nil {
		return result.Error
	}
	return nil
}

//login user
func (u *User) LoginUser(db *gorm.DB) error {

	var user User

	if result := db.Model(User{}).Where("email = ?", u.Email).Take(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

//find user from db
func (u *User) FindUser(db *gorm.DB) error {
	if u.UserID != 0 {
		if result := db.Model(&u).Preload(clause.Associations).Find(&u); result.Error != nil {
			return result.Error
		}
	}
	if result := db.Preload(clause.Associations).Model(&u).Where("email = ?", u.Email).Find(&u); result.Error != nil {
		return result.Error
	}
	return nil
}

// find user data and return it
func (u *User) FindUserData(db *gorm.DB) (*User, error) {
	if u.UserID != 0 {
		if result := db.Model(User{}).Preload(clause.Associations).Find(&u); result.Error != nil {
			return &User{}, result.Error
		}
	}
	if result := db.Model(User{}).Preload(clause.Associations).Where("email = ?", u.Email).Find(&u); result.Error != nil {
		return &User{}, result.Error
	}
	return u, nil
}

// find users basket
func (b *User) FindUserBasketbyUser(db *gorm.DB) (Basket, error) {

	b, err := b.FindUserData(db)
	if err != nil {
		return Basket{}, err
	}
	var basket Basket

	if result := db.Model(Basket{}).Preload(clause.Associations).Where("basket_id = ?", b.Basket.BasketID).First(&basket); result.Error != nil {
		return Basket{}, result.Error
	}
	return basket, nil
}
