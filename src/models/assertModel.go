package models

import (
	"errors"
	"html"
	"strings"

	"github.com/maykealisson/buy-and-hold/src/dtos"
	"gorm.io/gorm"
)

type Assert struct {
	Id           uint32
	Name         string
	Amount       int32
	Price        float64
	AveragePrice float64
	UserId       uint32
}

func (a *Assert) ToDomain(dto dtos.AssertDto) {
	a.Name = dto.Name
	a.Amount = dto.Amount
	a.Price = dto.Price
	a.AveragePrice = dto.AveragePrice
}

func (a *Assert) Prepare() {
	a.Name = html.EscapeString(strings.ToUpper(strings.TrimSpace(a.Name)))
}

func (a *Assert) Save(db *gorm.DB) (*Assert, error) {

	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Assert{}, err
	}
	return a, nil
}

func (a *Assert) FindByName(db *gorm.DB, name string, userId uint32) ([]Assert, error) {

	var err error
	var asserts []Assert
	err = db.Debug().Where("name LIKE ? AND user_id = ?", name+"%", userId).Find(&asserts).Error
	if err != nil {
		return []Assert{}, err
	}
	return asserts, nil
}

func (a *Assert) FindAllByUser(db *gorm.DB, userId uint32) ([]Assert, error) {

	var err error
	var asserts []Assert
	err = db.Debug().Where("user_id = ?", userId).Find(&asserts).Error
	if err != nil {
		return []Assert{}, err
	}
	return asserts, nil
}

func (a *Assert) FindByID(db *gorm.DB, uid uint32, userId uint32) (Assert, error) {
	var err error
	var result = Assert{}
	err = db.Debug().Where("id = ? AND user_id = ?", uid, userId).Take(&result).Error
	if err != nil {
		return Assert{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Assert{}, errors.New("Asser Not Found for to user")
	}
	return result, err
}

func (a *Assert) Update(db *gorm.DB) error {
	var err error
	err = db.Debug().Save(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Assert) Delete(db *gorm.DB, assertId uint32) error {
	var err error
	err = db.Debug().Where("id = ? ", assertId).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}
