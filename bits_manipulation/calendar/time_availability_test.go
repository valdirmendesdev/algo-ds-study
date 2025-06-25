package calendar_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/valdirmendesdev/algo-ds-study/bits_manipulation/calendar"
)

func TestDayTimeAvailability(t *testing.T) {
	t.Run("get full availability - bit representation", func(t *testing.T) {
		a := calendar.NewDayTimeAvailability()
		assert.Equal(t, "11111111 11111111 11111111 11111111 11111111 11111111", a.String())
	})
	t.Run("time list representation", func(t *testing.T) {
		t.Run("full available", func(t *testing.T) {
			a := calendar.NewDayTimeAvailability()
			got := a.TimeList()

			start := time.Date(2024, 4, 3, 0, 0, 0, 0, time.UTC)
			assert.Len(t, got, calendar.NumberOfHalfHoursInADay)
			for i := 0; i < calendar.NumberOfHalfHoursInADay; i++ {
				assert.Equal(t, start.Format(time.TimeOnly), got[i].Time)
				assert.Equal(t, true, got[i].Available)
				start = start.Add(time.Minute * 30)
			}
		})
	})
}

func TestTimeAvailability(t *testing.T) {

}
