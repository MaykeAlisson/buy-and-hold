package services

import (
	"time"

	"github.com/maykealisson/buy-and-hold/src/database"
	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/models"
	"github.com/maykealisson/buy-and-hold/src/utils"
)

type launchService struct{}

func LaunchService() *launchService {
	return &launchService{}
}

func (service *launchService) Create(userId uint32, assertId uint32, dto dtos.LauncheDto) (dtos.LauncheDto, error) {

	var err error
	assert := models.Assert{}
	assertReturn, err := assert.FindByID(database.DB, assertId, userId)
	if err != nil {
		return dtos.LauncheDto{}, err
	}

	launche := models.Launche{}
	err = launche.ToDomain(dto)
	if err != nil {
		return dtos.LauncheDto{}, err
	}

	launche.AssertId = assertReturn.Id
	launche.Prepare()

	if launche.Operation == "BUY" {
		assertReturn.Amount += launche.Amount
		assertReturn.InvestedAmount += (launche.Price * float64(launche.Amount))
		assertReturn.AveragePrice = (assertReturn.InvestedAmount / float64(assertReturn.Amount))
	} else {
		assertReturn.Amount -= launche.Amount
	}

	assertReturn.Price = launche.Price
	err = assertReturn.Update(database.DB)
	if err != nil {
		return dtos.LauncheDto{}, err
	}

	launchSave, err := launche.Save(database.DB)
	if err != nil {
		return dtos.LauncheDto{}, err
	}

	return dtos.LauncheDto{
		Operation:    launchSave.Operation,
		Amount:       launchSave.Amount,
		Price:        launchSave.Price,
		DateOperacao: launchSave.DateOperation.Format("2006-01-02"),
		Broker:       launchSave.Broker,
		Assert:       assertReturn.Name,
	}, nil

}

func (service *launchService) FindByMonth(userId uint32, month int) ([]dtos.LauncheDto, error) {

	var err error
	launche := models.Launche{}
	mes := time.Month(month)
	year := time.Now().Year()
	firstDay, lastDay := utils.DateUtils().MonthInterval(year, mes)
	results, err := launche.FindByMonth(database.DB, userId, firstDay, lastDay)
	if err != nil {
		return []dtos.LauncheDto{}, err
	}

	return results, nil

}

func (service *launchService) FindByAssert(userId uint32, assertId uint32) ([]dtos.LauncheDto, error) {

	var err error
	assert := models.Assert{}
	assertReturn, err := assert.FindByID(database.DB, assertId, userId)
	if err != nil {
		return []dtos.LauncheDto{}, err
	}
	launche := models.Launche{}
	results, err := launche.FindByAssertId(database.DB, assertReturn.Id)
	if err != nil {
		return []dtos.LauncheDto{}, err
	}

	launches := []dtos.LauncheDto{}

	for _, value := range results {
		launches = append(launches, dtos.LauncheDto{
			Operation:    value.Operation,
			Amount:       value.Amount,
			Price:        value.Price,
			DateOperacao: value.DateOperation.Format("2006-01-02"),
			Broker:       value.Broker,
			Assert:       assertReturn.Name,
		})
	}

	return launches, nil

}

func (service *launchService) Delete(userId uint32, launchtId uint32) error {

	var err error
	launche := models.Launche{}
	assert := models.Assert{}
	result, err := launche.FindById(database.DB, launchtId)
	if err != nil {
		return err
	}
	assert, err = assert.FindByID(database.DB, result.AssertId, userId)
	if err != nil {
		return err
	}

	if launche.Operation == "BUY" {
		assert.Amount += launche.Amount
		assert.InvestedAmount += (launche.Price * float64(launche.Amount))
		assert.AveragePrice = (assert.InvestedAmount / float64(assert.Amount))
	} else {
		assert.Amount -= launche.Amount
		assert.InvestedAmount -= (launche.Price * float64(launche.Amount))
		assert.AveragePrice = (assert.InvestedAmount / float64(assert.Amount))
	}

	err = assert.Update(database.DB)
	if err != nil {
		return err
	}

	err = launche.DeleteById(database.DB)
	if err != nil {
		return err
	}

	return nil

}
