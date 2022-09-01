package dtos

import (
	"errors"
	"strings"
	"time"
)

type LaunchDto struct {
	Operation    string    `json:"operation"`
	Amount       int32     `json:"amount"`
	Price        float64   `json:"price"`
	DateOperacao time.Time `json:"date_operation"`
	Broker       string    `json:"broker"`
}

func (l *LaunchDto) Validate(action string) error {
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
		// if l.DateOperacao == 0 {
		// 	return errors.New("Required date_operation")
		// }
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
		// if l.DateOperacao == "" {
		// 	return errors.New("Required date_operation")
		// }
		if l.Broker == "" {
			return errors.New("Required broker")
		}
		return nil
	}

}
