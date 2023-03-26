package calendar_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/algo-ds-study/bits_manipulation/calendar"
)

func TestMonthAvailability_NumberOfDays(t *testing.T) {
	const (
		_31DaysInDecimal int32 = 2147483647
		_30DaysInDecimal int32 = 1073741823
		_29DaysInDecimal int32 = 536870911
		_28DaysInDecimal int32 = 268435455
		FullUnavailable        = 0
	)
	assert.Equal(t, calendar.Month_With31Days, calendar.MonthAvailability(_31DaysInDecimal))
	assert.Equal(t, calendar.Month_With30Days, calendar.MonthAvailability(_30DaysInDecimal))
	assert.Equal(t, calendar.Month_With29Days, calendar.MonthAvailability(_29DaysInDecimal))
	assert.Equal(t, calendar.Month_With28Days, calendar.MonthAvailability(_28DaysInDecimal))
	assert.Equal(t, calendar.Month_FullUnavailable, calendar.MonthAvailability(FullUnavailable))
}

func TestMonthAvailability_CheckIfDayIsAvailable(t *testing.T) {
	tests := []struct {
		Month         calendar.MonthAvailability
		CheckedDay    int
		ExpectedValue bool
	}{
		{
			Month:         calendar.Month_With31Days,
			CheckedDay:    31,
			ExpectedValue: true,
		},
		{
			Month:         calendar.Month_With30Days,
			CheckedDay:    31,
			ExpectedValue: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.ExpectedValue, test.Month.IsDayAvailable(test.CheckedDay), fmt.Sprintf("month %032b checking day %v", test.Month, test.CheckedDay))
	}
}

func TestMonthAvailability_SetDayUnavailable(t *testing.T) {
	m := calendar.Month_With31Days
	m.SetDayUnavailable(15)
	assert.False(t, m.IsDayAvailable(15))
}

func TestMonthAvailability_SetDayAvailable(t *testing.T) {
	m := calendar.Month_FullUnavailable
	assert.False(t, m.IsDayAvailable(15))

	m.SetDayAvailable(15)
	assert.True(t, m.IsDayAvailable(15))
}

func TestMonthAvailability_CheckAvailableDaysInOneShoot(t *testing.T) {
	m := calendar.Month_With31Days

	//checked days
	availability := calendar.DayToBitRepresentation(3) | calendar.DayToBitRepresentation(5) | calendar.DayToBitRepresentation(15)
	unavailableDays := calendar.MonthAvailability(availability)

	//Check all days using one shoot
	assert.True(t, m.IsAvailable(unavailableDays))

	m.SetDayUnavailable(15)
	assert.False(t, m.IsAvailable(unavailableDays))
}

func TestMonthAvailability_BitRepresentationString(t *testing.T) {
	m := calendar.MonthAvailability(1)
	assert.Equal(t, "0000000 00000000 00000000 00000001", m.String())
	assert.Equal(t, "1111111 11111111 11111111 11111111", calendar.Month_With31Days.String())
}
