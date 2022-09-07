package dtos

import (
	"errors"
	"strings"
)

type LauncheDto struct {
	Operation    string  `json:"operation"`
	Amount       int32   `json:"amount"`
	Price        float64 `json:"price"`
	DateOperacao string  `json:"date_operation"`
	Broker       string  `json:"broker"`
	Assert       string  `json:"assert"`
}

func (l *LauncheDto) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if l.Operation == "" {
			return errors.New("Required operation[BUY or SELL]")
		}
		if l.Operation != "" {
			if strings.ToUpper(l.Operation) != "BUY" || strings.ToUpper(l.Operation) != "SELL" {
				return errors.New("Invalid operation [BUY or SELL]")
			}
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
			return errors.New("Required operation [BUY or SELL]")
		}
		if l.Operation != "" {
			if strings.ToUpper(l.Operation) != "BUY" && strings.ToUpper(l.Operation) != "SELL" {
				return errors.New("Invalid operation [BUY or SELL]")
			}
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
