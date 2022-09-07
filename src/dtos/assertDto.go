package dtos

import (
	"errors"
	"strings"
)

type AssertDto struct {
	Id             uint32  `json:"id"`
	Name           string  `json:"name"`
	Amount         int32   `json:"amount"`
	Price          float64 `json:"price"`
	AveragePrice   float64 `json:"average_price"`
	InvestedAmount float64 `json:"invested_amount"`
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
		if dto.InvestedAmount != 0 {
			if dto.InvestedAmount < 0 {
				return errors.New("Invalid invested_amount")
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
