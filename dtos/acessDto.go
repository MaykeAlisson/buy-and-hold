package dtos

type AcessDto struct {
	UserId uint32 `json:"user_id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
