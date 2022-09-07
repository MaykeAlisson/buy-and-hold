package services

import (
	"errors"

	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/models"
	"github.com/maykealisson/buy-and-hold/src/providers"
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

func AuthService() *authService {
	return &authService{}
}

func (a *authService) Longin(dto dtos.UserDto) (dtos.AcessDto, error) {
	user := models.User{}
	user.ToDomain(dto)
	user.Prepare()

	acess, err := a.SingIn(user.Email, user.Password)
	if err != nil {
		return dtos.AcessDto{}, err
	}

	return acess, nil

}

func (a *authService) SingIn(email string, password string) (dtos.AcessDto, error) {

	var err error
	user := models.User{}

	err = database.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return dtos.AcessDto{}, errors.New("User with unregistered email " + email)
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return dtos.AcessDto{}, errors.New("password not valid ")
	}

	token, err := providers.JwtProvider().CreateToken(user.Id)
	if err != nil {
		return dtos.AcessDto{}, err
	}

	return dtos.AcessDto{
		UserId: user.Id,
		Name:   user.Name,
		Token:  token,
	}, nil

}
