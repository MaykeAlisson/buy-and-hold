package dtos

import (
	"errors"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
