package models

type Assert struct {
	Id           uint32
	Name         string
	Amount       int32
	Price        int64
	AveragePrice int64
	UserId       uint32
}
