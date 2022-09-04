package utils

import "time"

type datesctuct struct {
}

func DateUtils() *datesctuct {
	return &datesctuct{}
}

func (d *datesctuct) ParseDate(date string) (time.Time, error) {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time, err
	}
	return time, nil
}

func (d *datesctuct) MonthInterval(y int, m time.Month) (firstDay, lastDay string) {
	firstDay = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC).Format("2006-01-02")
	return firstDay, lastDay
}
