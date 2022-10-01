package services

import (
	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/models"
)

type accertService struct{}

func AssertService() *accertService {
	return &accertService{}
}

func (service *accertService) CreateAssert(userId uint32, dto dtos.AssertDto) (dtos.AssertDto, error) {

	var err error
	assert := models.Assert{}
	assert.ToDomain(dto)
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

	var err error
	var assert = models.Assert{}
	asserts, err := assert.FindByName(database.DB, name, userId)
	if err != nil {
		return []dtos.AssertDto{}, err
	}

	var results = []dtos.AssertDto{}

	for _, value := range asserts {
		results = append(results, dtos.AssertDto{
			Id:             value.Id,
			Name:           value.Name,
			Amount:         value.Amount,
			Price:          value.Price,
			AveragePrice:   value.AveragePrice,
			InvestedAmount: value.InvestedAmount,
		})
	}

	return results, nil

}

func (service *accertService) FindAllByUser(userId uint32) ([]dtos.AssertDto, error) {

	var err error
	var assert = models.Assert{}
	asserts, err := assert.FindAllByUser(database.DB, userId)
	if err != nil {
		return []dtos.AssertDto{}, err
	}

	var results = []dtos.AssertDto{}

	for _, value := range asserts {
		results = append(results, dtos.AssertDto{
			Id:             value.Id,
			Name:           value.Name,
			Amount:         value.Amount,
			Price:          value.Price,
			AveragePrice:   value.AveragePrice,
			InvestedAmount: value.InvestedAmount,
		})
	}

	return results, nil

}

func (service *accertService) Update(assertId uint32, userId uint32, dto dtos.AssertDto) (dtos.AssertDto, error) {

	var err error
	assert := models.Assert{}
	assert.ToDomain(dto)
	assertReturn, err := assert.FindByID(database.DB, assertId, userId)
	if err != nil {
		return dtos.AssertDto{}, err
	}
	assertReturn.Amount = assert.Amount
	assertReturn.Name = assert.Name
	assertReturn.AveragePrice = assert.AveragePrice
	assertReturn.Price = assert.Price

	assertReturn.Prepare()

	err = assertReturn.Update(database.DB)
	if err != nil {
		return dtos.AssertDto{}, err
	}

	return dtos.AssertDto{
		Id:             assertReturn.Id,
		Name:           assertReturn.Name,
		Price:          assertReturn.Price,
		Amount:         assertReturn.Amount,
		AveragePrice:   assertReturn.AveragePrice,
		InvestedAmount: assertReturn.InvestedAmount,
	}, nil

}

func (service *accertService) Delete(assertId uint32, userId uint32) error {

	var err error
	assert := models.Assert{}
	launch := models.Launche{}
	assertReturn, err := assert.FindByID(database.DB, assertId, userId)
	if err != nil {
		return err
	}
	launch.DeleteByAccert(database.DB, assertReturn.Id)
	if err != nil {
		return err
	}

	assert.Delete(database.DB, assertReturn.Id)
	if err != nil {
		return err
	}
	return nil

}
