package dtos

import (
	"errors"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/models"
	"github.com/maykealisson/buy-and-hold/src/utils"
)

type LauncheDto struct {
	Operation    string  `json:"operation"`
	Amount       int32   `json:"amount"`
	Price        float64 `json:"price"`
	DateOperacao string  `json:"date_operation"`
	Broker       string  `json:"broker"`
}

func (l *LauncheDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if l.Operation == "" {
			return errors.New("Required operation")
		}
		if l.Amount == 0 {
			return errors.New("Required amount")
		}
		if l.Price == 0 {
			return errors.New("Required price")
		}
		if l.DateOperacao == "" {
			return errors.New("Required date_operation - YYYY-MM-dd")
		}
		if l.Broker == "" {
			return errors.New("Required broker")
		}
		return nil

	default:
		if l.Operation == "" {
			return errors.New("Required operation")
		}
		if l.Amount == 0 {
			return errors.New("Required amount")
		}
		if l.Price == 0 {
			return errors.New("Required price")
		}
		if l.DateOperacao == "" {
			return errors.New("Required date_operation - YYYY-MM-dd")
		}
		if l.Broker == "" {
			return errors.New("Required broker")
		}
		return nil
	}

}

func (l *LauncheDto) ToDomain() (models.Launche, error) {
	data, err := utils.DateUtils().ParseDate(l.DateOperacao)
	if err != nil {
		return models.Launche{}, err
	}
	return models.Launche{
		Operation:     l.Operation,
		Amount:        l.Amount,
		Price:         l.Price,
		DateOperation: data,
		Broker:        l.Broker,
	}, nil
}
