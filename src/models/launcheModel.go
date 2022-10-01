package models

import (
	"html"
	"strings"
	"time"

	"github.com/maykealisson/buy-and-hold/src/dtos"
	"github.com/maykealisson/buy-and-hold/src/utils"
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

func (l *Launche) ToDomain(dto dtos.LauncheDto) error {
	data, err := utils.DateUtils().ParseDate(dto.DateOperacao)
	if err != nil {
		return err
	}
	l.Operation = dto.Operation
	l.Amount = dto.Amount
	l.Price = dto.Price
	l.DateOperation = data
	l.Broker = dto.Broker
	return nil
}

func (l *Launche) Prepare() {
	l.Broker = html.EscapeString(strings.ToUpper(strings.TrimSpace(l.Broker)))
	l.Operation = html.EscapeString(strings.ToUpper(strings.TrimSpace(l.Operation)))
}

func (l *Launche) Save(db *gorm.DB) (*Launche, error) {

	var err error
	err = db.Debug().Create(&l).Error
	if err != nil {
		return &Launche{}, err
	}
	return l, nil
}

func (l *Launche) FindById(db *gorm.DB, id uint32) (*Launche, error) {
	var err error
	err = db.Debug().Where("id = ? ", id).Take(&l).Error
	if err != nil {
		return &Launche{}, err
	}

	return l, nil

}

func (l *Launche) FindByAssertId(db *gorm.DB, assertId uint32) ([]Launche, error) {
	var err error
	var results []Launche
	err = db.Debug().Where("assert_id = ? ", assertId).Order("date_operation desc").Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil

}

func (l *Launche) FindByMonth(db *gorm.DB, userId uint32, startDate string, endDate string) ([]dtos.LauncheDto, error) {

	var err error
	type Result struct {
		Id            uint32
		Operation     string
		Amount        int32
		Price         float64
		DateOperation time.Time
		Broker        string
		AssertId      uint32
		Assert        string
	}
	var results []Result
	var launches []dtos.LauncheDto
	err = db.Debug().Table("launches l").Select("l.*, a.name as assert ").Joins("inner join asserts a on l.assert_id = a.id AND a.user_id = ? ", userId).Where("l.date_operation between ? and ?", startDate, endDate).Order("l.date_operation desc").Find(&results).Error
	if err != nil {
		return []dtos.LauncheDto{}, err
	}

	for _, value := range results {
		launches = append(launches, dtos.LauncheDto{
			Id:           value.Id,
			Operation:    value.Operation,
			Amount:       value.Amount,
			Price:        value.Price,
			DateOperacao: value.DateOperation.Format("2006-01-02"),
			Broker:       value.Broker,
			Assert:       value.Assert,
		})
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

func (l *Launche) DeleteById(db *gorm.DB) error {
	var err error
	err = db.Debug().Where("id = ? ", l.Id).Delete(&l).Error
	if err != nil {
		return err
	}
	return nil
}
