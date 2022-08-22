package dtos

import (
	"errors"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/models"
)

type AssertDto struct {
	Id           uint32 `json:"id"`
	Name         string `json:"name"`
	Amount       int32 `json:"amount"`
	Price        int64 `json:"price"`
	AveragePrice int64 `json:"average_price"`
}

func (dto *AssertDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if dto.Name == "" {
			return errors.New("Required name")
		}
		if dto.Amount != "" {
			// verificar se nao e negativo
			return errors.New("Required name")
		}
		if dto.Price != "" {
			// verificar se nao e negativo
			return errors.New("Required name")
		}
		if dto.AveragePrice != "" {
			// verificar se nao e negativo
			return errors.New("Required name")
		}
		return nil

	default:
		if dto.Name == "" {
			return errors.New("Required name")
		}
		return nil
	}
}

func (dto *AssertDto) ToDomain() models.Assert {
	return models.Assert{
		Name: dto.Name,
		Amount: dto.Amount,     
		Price: dto.Price,     
		AveragePrice: dto.AveragePrice,
	}
}
