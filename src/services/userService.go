package services

import (
	"errors"

	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/providers"
)

type userService struct{}

func UserService() *userService {
	return &userService{}
}

func (service *userService) CreateUser(dto dtos.UserDto) (dtos.AcessDto, error) {

	var err error

	user := dto.ToDomain()
	user.Prepare()
	user.BeforeSave()

	exists, err := user.ExistsEmail(database.DB)
	if err != nil {
		return dtos.AcessDto{}, err
	}
	if exists {
		return dtos.AcessDto{}, errors.New("email already registered!")
	}

	userCreated, err := user.SaveUser(database.DB)
	if err != nil {
		return dtos.AcessDto{}, err
	}

	token, err := providers.JwtProvider().CreateToken(user.Id)
	if err != nil {
		return dtos.AcessDto{}, err
	}

	var acess = dtos.AcessDto{
		UserId: userCreated.Id,
		Name:   userCreated.Name,
		Token:  token,
	}

	return acess, nil

}

func (service *userService) UpdateUser(userId uint32, dto dtos.UserDto) error {

	// verifica se existe usuario com id
	// altera o nome, email e senha

	return nil

}

func (service *userService) DeleteUser(userId uint32) error {

	// verifica se existe usuario com id
	// deleta todos os registros relacionado ao usuario

	return nil

}
