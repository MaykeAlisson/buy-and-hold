package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id       uint32
	Name     string
	Email    string
	Password string
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
