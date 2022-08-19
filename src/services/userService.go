package services

import (
	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
)

func CreateUser(dto dtos.UserDto) (dtos.AcessDto, error) {

	user := dto.ToDomain()

	// verifica se email ja nao foi cadastrado

	userCreated, err := user.SaveUser(database.DB)
	if err != nil {

		//formattedError := formaterror.FormatError(err.Error())

		//responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return dtos.AcessDto{}, err
	}

	// gerar token

	var acess = dtos.AcessDto{
		UserId: userCreated.Id,
		Name:   userCreated.Name,
		Token:  "",
	}

	return acess, nil

}

func UpdateUser(userId uint32, dto dtos.UserDto) error {

	// verifica se existe usuario com id
	// altera o nome, email e senha

	return nil

}

func DeleteUser(userId uint32) error {

	// verifica se existe usuario com id
	// deleta todos os registros relacionado ao usuario

	return nil

}
