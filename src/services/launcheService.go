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

	// e criar um launch para este assert

	// se for uma compra soma a qtd ao total do assert e faz preço medio
	if launche.Operation == "BUY" {
		assertReturn.Amount += launche.Amount
	} else {
		assertReturn.Amount -= launche.Amount
	}

	assertReturn.Price = launche.Price
	// se for uma venda diminui a qtd do assert
	// atualiza valor do assert com o valor informado no lancamento
	// atualizar o assert
	err = assertReturn.Update(database.DB)
	if err != nil {
		return dtos.LauncheDto{}, err
	}

	// retornar launchDto

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

	// pegar o id do usuario
	// pega o numero mes pegar primeiro dia e ultimo dia do mes e fazer query
	// busca todos os lancamentos do mes informado para este usuario agrupando por data

	// retornar lista de launchDto
	return results, nil

}

func (service *launchService) FindByAssert(userId uint32, assertId uint32) ([]dtos.LauncheDto, error) {

	// var err error
	// pegar o id do usuario
	// pega o assertId
	// busca todos os lancamentos do assert informado este usuario agrupando por data

	// retornar lista de launchDto
	return []dtos.LauncheDto{}, nil

}

func (service *launchService) Update(userId uint32, assertId uint32, launchId uint32, dto dtos.LauncheDto) (dtos.LauncheDto, error) {

	// var err error
	// pegar o id do usuario
	// pegar o id do assert
	// pega o id do launch
	// busca o assert
	// atualiza o launch

	// se for uma compra soma a qtd ao total do assert e faz preço medio
	// se for uma venda diminui a qtd do assert
	// atualiza valor do assert com o valor informado no lancamento
	// atualizar o assert

	// retornar launchDto
	return dtos.LauncheDto{}, nil

}

func (service *launchService) Delete(userId uint32, assertId uint32, launchtId uint32) error {

	// var err error
	// pegar o id do usuario
	// pegar o id do assert
	// pega o id do launch
	// busca o assert
	// buscar o launch

	// se for uma compra diminiir a qtd ao total do assert e faz preço medio
	// se for uma venda almentar a qtd do assert

	// deletar o launch
	// atualizar o assert

	// retornar launchDto
	return nil

}
