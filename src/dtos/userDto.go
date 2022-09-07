package dtos

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("Required email")
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid Email")
			}
		}
		return nil
	case "login":
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

	default:
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
}
