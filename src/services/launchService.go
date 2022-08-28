package services

import (
	"github.com/maykealisson/buy-and-hold/src/dtos"
)

type launchService struct{}

func LaunchService() *launchService {
	return &launchService{}
}

func (service *launchService) Create(userId uint32, assertId uint32, dto dtos.LaunchDto) (dtos.LaunchDto, error) {

	// var err error
	// pegar o id do usuario
	// pegar o id do assert
	// busca o assert
	// e criar um launch para este assert

	// se for uma compra soma a qtd ao total do assert e faz preço medio
	// se for uma venda diminui a qtd do assert
	// atualiza valor do assert com o valor informado no lancamento
	// atualizar o assert

	// retornar launchDto
	return dtos.LaunchDto{}, nil

}

func (service *launchService) FindByMonth(userId uint32, month int) ([]dtos.LaunchDto, error) {

	// var err error
	// pegar o id do usuario
	// pega o numero mes pegar primeiro dia e ultimo dia do mes e fazer query
	// busca todos os lancamentos do mes informado para este usuario agrupando por data

	// retornar lista de launchDto
	return []dtos.LaunchDto{}, nil

}

func (service *launchService) FindByAssert(userId uint32, assertId uint32) ([]dtos.LaunchDto, error) {

	// var err error
	// pegar o id do usuario
	// pega o assertId
	// busca todos os lancamentos do assert informado este usuario agrupando por data

	// retornar lista de launchDto
	return []dtos.LaunchDto{}, nil

}

func (service *launchService) Update(userId uint32, assertId uint32, launchId uint32, dto dtos.LaunchDto) (dtos.LaunchDto, error) {

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
	return dtos.LaunchDto{}, nil

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
