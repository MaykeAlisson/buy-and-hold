package dtos

import (
	"errors"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/models"
)

type AssertDto struct {
	Name string `json:"name"`
}

func (dto *AssertDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if dto.Name == "" {
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
	}
}
