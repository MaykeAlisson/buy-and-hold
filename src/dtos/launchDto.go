package dtos

import "time"

type LaunchDto struct {
	Operation    string    `json:"operation"`
	Amount       int32     `json:"amount"`
	Price        int64     `json:"price"`
	DateOperacao time.Time `json:"date_operation"`
	Broker       string    `json:"broker"`
}
