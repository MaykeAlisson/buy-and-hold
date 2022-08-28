package services

import (
	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
)

type accertService struct{}

func AssertService() *accertService {
	return &accertService{}
}

func (service *accertService) CreateAssert(userId uint32, dto dtos.AssertDto) (dtos.AssertDto, error) {

	var err error
	assert := dto.ToDomain()
	assert.UserId = userId
	assert.Prepare()

	assertCreated, err := assert.Save(database.DB)
	if err != nil {
		return dtos.AssertDto{}, err
	}

	return dtos.AssertDto{
		Id:           assertCreated.Id,
		Name:         assertCreated.Name,
		Amount:       assertCreated.Amount,
		Price:        assertCreated.Price,
		AveragePrice: assertCreated.AveragePrice,
	}, nil

}

func (service *accertService) FindByName(userId uint32, name string) ([]dtos.AssertDto, error) {

	// var err error
	// pegar o id do usuario
	// busca por like para o este usuario

	// retornar lista assertDto
	return []dtos.AssertDto{}, nil

}

func (service *accertService) Update(assertId uint32, userId uint32, dto dtos.AssertDto) (dtos.AssertDto, error) {

	// var err error
	// pegar o id do usuario
	// busca assert com id e idUser
	// atualiza assert

	// retornar assertDto
	return dtos.AssertDto{}, nil

}

func (service *accertService) Delete(assertId uint32, userId uint32) error {

	// var err error
	// pegar o id do usuario
	// busca assert com id e idUser
	// deleta lançamentos para este assert
	// deleta assert

	// retornar erro se tiver
	return nil

}
