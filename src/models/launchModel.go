package models

import (
	"time"

	"gorm.io/gorm"
)

type Launch struct {
	Id           uint32
	Operation    string
	Amount       int32
	Price        float64
	DateOperacao time.Time
	Broker       string
	AccertId     uint32
}

func (l *Launch) DeleteByAccert(db *gorm.DB, assertId uint32) error {
	var err error
	err = db.Debug().Where("assert_id = ? ", assertId).Delete(&l).Error
	if err != nil {
		return err
	}
	return nil
}
