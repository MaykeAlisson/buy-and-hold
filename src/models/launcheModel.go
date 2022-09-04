package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Launche struct {
	Id            uint32
	Operation     string
	Amount        int32
	Price         float64
	DateOperation time.Time
	Broker        string
	AssertId      uint32
}

func (l *Launche) Prepare() {
	l.Broker = html.EscapeString(strings.ToUpper(strings.TrimSpace(l.Broker)))
}

func (l *Launche) Save(db *gorm.DB) (*Launche, error) {

	var err error
	err = db.Debug().Create(&l).Error
	if err != nil {
		return &Launche{}, err
	}
	return l, nil
}

func (l *Launche) FindByMonth(db *gorm.DB, userId uint32, startDate string, endDate string) ([]Launche, error) {

	var err error
	var launches []Launche
	err = db.Debug().Table("launches l").Select("l.*, a.name as assert ").Joins("inner join asserts a on l.assert_id = a.id AND a.user_id = ? ", userId).Where("l.date_operation between ? and ?", startDate, endDate).Find(&launches).Error
	if err != nil {
		return []Launche{}, err
	}
	return launches, nil
}

func (l *Launche) DeleteByAccert(db *gorm.DB, assertId uint32) error {
	var err error
	err = db.Debug().Where("assert_id = ? ", assertId).Delete(&l).Error
	if err != nil {
		return err
	}
	return nil
}
