package calendar

import (
	"fmt"
	"strconv"
	"time"
)

const (
	fullAvailable           = 281474976710655 // 2**48
	NumberOfHalfHoursInADay = 48
	halfHour                = time.Minute * 30
)

type TimeAvailability struct {
	Time      string
	Available bool
}

type DayTimeAvailability interface {
	String() string
	TimeList() []TimeAvailability
}

type dayTimeAvailability struct {
	availability int64
}

func (d *dayTimeAvailability) TimeList() []TimeAvailability {
	date := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	result := make([]TimeAvailability, 0, NumberOfHalfHoursInADay)
	for i := 0; i < NumberOfHalfHoursInADay; i++ {
		available, err := strconv.ParseBool(d.getBit(i))
		if err != nil {
			available = false
		}
		result = append(result, TimeAvailability{
			Time:      date.Format(time.TimeOnly),
			Available: available,
		})
		date = date.Add(halfHour)
	}
	return result
}

func NewDayTimeAvailability() DayTimeAvailability {
	return &dayTimeAvailability{
		availability: fullAvailable,
	}
}

func (d *dayTimeAvailability) getBit(index int) string {
	s := d.getBitString()
	pos := NumberOfHalfHoursInADay - index - 1
	return s[pos : pos+1]
}

func (d *dayTimeAvailability) getBitString() string {
	return fmt.Sprintf("%048b", d.availability)
}

func (d *dayTimeAvailability) String() string {
	s := d.getBitString()
	return fmt.Sprintf("%s %s %s %s %s %s", s[0:8], s[8:16], s[16:24], s[24:32], s[32:40], s[40:48])
}
