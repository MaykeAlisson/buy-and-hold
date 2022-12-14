package models

import (
	"errors"
	"html"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/dtos"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint32
	Name     string
	Email    string
	Password string
}

func (u *User) ToDomain(dto dtos.UserDto) {
	u.Name = dto.Name
	u.Email = dto.Email
	u.Password = dto.Password
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) ExistsEmail(db *gorm.DB, email string) (bool, error) {
	var err error
	result := map[string]interface{}{}
	err = db.Debug().Model(User{}).Where("email = ?", email).Take(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, err
}

func (u *User) Update(db *gorm.DB) error {
	var err error
	err = db.Debug().Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}
