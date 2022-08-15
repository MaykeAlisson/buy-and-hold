package models

type Assert struct {
	Id           uint32 `json:"id"`
	Name         string `json:"name"`
	Amount       int32  `json:"amount"`
	Price        int64  `json:"price"`
	AvegerePrice int64  `json:"avegere_price"`
}
