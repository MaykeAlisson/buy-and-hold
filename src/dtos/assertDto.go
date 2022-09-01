package dtos

import (
	"errors"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/models"
)

type AssertDto struct {
	Id           uint32  `json:"id"`
	Name         string  `json:"name"`
	Amount       int32   `json:"amount"`
	Price        float64 `json:"price"`
	AveragePrice float64 `json:"average_price"`
}

func (dto *AssertDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if dto.Name == "" {
			return errors.New("Required name")
		}
		if dto.Amount != 0 {
			if dto.Amount < 0 {
				return errors.New("Invalid amount")
			}
		}
		if dto.Price != 0 {
			if dto.Price < 0 {
				return errors.New("Invalid price")
			}
		}
		if dto.AveragePrice != 0 {
			if dto.AveragePrice < 0 {
				return errors.New("Invalid average_price")
			}
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
		Name:         dto.Name,
		Amount:       dto.Amount,
		Price:        dto.Price,
		AveragePrice: dto.AveragePrice,
	}
}
