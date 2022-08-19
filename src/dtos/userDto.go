package dtos

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/maykealisson/buy-and-hold/src/models"
)

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserDto) Validate() error {
	if u.Name == "" {
		return errors.New("Required name")
	}
	if u.Email == "" {
		return errors.New("Required email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}
	if u.Password == "" {
		return errors.New("Required password")
	}
	return nil
}

func (u *UserDto) ToDomain() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
