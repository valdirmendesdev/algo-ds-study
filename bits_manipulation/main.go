package main

import (
	"fmt"
	"time"

	"github.com/valdirmendesdev/algo-ds-study/bits_manipulation/calendar"
)

func main() {
	start := time.Now()
	m := calendar.Month_With31Days
	PrintMonthAvailability(m)

	m.SetDayUnavailable(15)
	PrintMonthAvailability(m)

	m.SetDayUnavailable(23)
	PrintMonthAvailability(m)

	fmt.Println("\n*** Check if days are available ***\n")
	PrintDayIsAvailable(m, 1)
	PrintDayIsAvailable(m, 15)
	PrintDayIsAvailable(m, 23)
	PrintDayIsAvailable(m, 30)

	m = calendar.Month_FullUnavailable
	fmt.Println("\n*** Month without available days ***\n")
	PrintMonthAvailability(m)

	fmt.Println("\n*** Check many days availability in one shoot ***\n")
	m = calendar.Month_With30Days
	//Represents October month with Brazilian holidays
	m.SetDayUnavailable(12).SetDayUnavailable(28)

	checkedDays := calendar.Month_FullUnavailable
	//Days that I want check if available
	checkedDays.SetDayAvailable(3).SetDayAvailable(12).SetDayAvailable(28)

	PrintMonthAvailability(m)
	fmt.Printf("Are days 3,12,15 available: %v\n", m.IsAvailable(checkedDays))

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time %s", elapsed)
}

func PrintMonthAvailability(m calendar.MonthAvailability) {
	fmt.Printf("Month availability: %s\n", m)
}

func PrintDayIsAvailable(m calendar.MonthAvailability, day int) {
	PrintMonthAvailability(m)
	fmt.Printf("Is Day %v available: %v\n", day, m.IsDayAvailable(day))
}
