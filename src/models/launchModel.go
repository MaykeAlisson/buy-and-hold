package models

import "time"

type Launch struct {
	Id           uint32
	Operation    string
	Amount       int32
	Price        int64
	DateOperacao time.Time
	Broker       string
	AccertId     uint32
}
