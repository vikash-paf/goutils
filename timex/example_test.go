package timex_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/timex"
	"time"
)

func ExampleStartOfDay() {
	t := time.Date(2024, 1, 1, 15, 30, 0, 0, time.UTC)
	fmt.Println(timex.StartOfDay(t).Format("15:04:05"))
	// Output: 00:00:00
}

func ExampleEndOfDay() {
	t := time.Date(2024, 1, 1, 15, 30, 0, 0, time.UTC)
	fmt.Println(timex.EndOfDay(t).Format("15:04:05"))
	// Output: 23:59:59
}

func ExampleIsWeekend() {
	t := time.Date(2024, 1, 6, 12, 0, 0, 0, time.UTC) // Saturday
	fmt.Println(timex.IsWeekend(t))
	// Output: true
}

func ExampleAddBusinessDays() {
	t := time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC) // Friday
	future := timex.AddBusinessDays(t, 1)
	fmt.Println(future.Weekday())
	// Output: Monday
}
