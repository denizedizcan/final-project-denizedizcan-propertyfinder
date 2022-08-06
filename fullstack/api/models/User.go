package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/badoux/checkmail"
)

//user struct fields used in db
type User struct {
	UserID    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"Name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
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

func (u *User) SaveUser(db *gorm.DB) error {

	if result := db.Create(&u); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) LoginUser(db *gorm.DB) error {

	var user User

	if result := db.Model(User{}).Where("email = ?", u.Email).Take(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
