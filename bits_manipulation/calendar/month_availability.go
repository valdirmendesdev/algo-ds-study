package calendar

import (
	"fmt"
)

type MonthAvailability int32

const (
	Month_With31Days      MonthAvailability = 0x7fffffff
	Month_With30Days      MonthAvailability = Month_With31Days >> 1
	Month_With29Days      MonthAvailability = Month_With30Days >> 1
	Month_With28Days      MonthAvailability = Month_With29Days >> 1
	Month_FullUnavailable MonthAvailability = 0
)

func (m MonthAvailability) IsDayAvailable(day int) bool {
	dayInBinary := DayToBitRepresentation(day)
	return m.IsAvailable(dayInBinary)
}

func DayToBitRepresentation(day int) MonthAvailability {
	return MonthAvailability(1 << (day - 1))
}

func (m *MonthAvailability) SetDayUnavailable(day int) *MonthAvailability {
	dayInBinary := DayToBitRepresentation(day)
	*m = *m ^ dayInBinary
	return m
}

func (m *MonthAvailability) SetDayAvailable(day int) *MonthAvailability {
	dayInBinary := DayToBitRepresentation(day)
	*m = *m | dayInBinary
	return m
}

func (m MonthAvailability) IsAvailable(month MonthAvailability) bool {
	result := m&month >= month
	return result
}

func (m MonthAvailability) String() string {
	s := fmt.Sprintf("%031b", m)
	return fmt.Sprintf("%s %s %s %s", s[0:7], s[7:15], s[15:23], s[23:31])
}
