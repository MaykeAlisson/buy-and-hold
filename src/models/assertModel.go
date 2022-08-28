package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Assert struct {
	Id           uint32
	Name         string
	Amount       int32
	Price        int64
	AveragePrice int64
	UserId       uint32
}

func (a *Assert) Prepare() {
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
}

func (a *Assert) Save(db *gorm.DB) (*Assert, error) {

	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Assert{}, err
	}
	return a, nil
}
