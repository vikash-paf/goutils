package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestStartOfDay(t *testing.T) {
	now := time.Now()
	sod := StartOfDay(now)

	if sod.Hour() != 0 || sod.Minute() != 0 || sod.Second() != 0 || sod.Nanosecond() != 0 {
		t.Errorf("StartOfDay() = %v, want 00:00:00", sod)
	}
}

func TestEndOfDay(t *testing.T) {
	now := time.Now()
	eod := EndOfDay(now)

	if eod.Hour() != 23 || eod.Minute() != 59 || eod.Second() != 59 {
		t.Errorf("EndOfDay() = %v, want 23:59:59", eod)
	}
}

func TestStartOfWeek(t *testing.T) {
	// 2023-10-18 is a Wednesday
	dt := time.Date(2023, 10, 18, 12, 0, 0, 0, time.UTC)
	sow := StartOfWeek(dt, time.Monday)

	if sow.Weekday() != time.Monday {
		t.Errorf("StartOfWeek() day = %v, want Monday", sow.Weekday())
	}
	if sow.Day() != 16 {
		t.Errorf("StartOfWeek() date = %d, want 16", sow.Day())
	}
}

func TestWeekendWeekday(t *testing.T) {
	// 2023-10-21 is a Saturday
	sat := time.Date(2023, 10, 21, 12, 0, 0, 0, time.UTC)
	// 2023-10-20 is a Friday
	fri := time.Date(2023, 10, 20, 12, 0, 0, 0, time.UTC)

	if !IsWeekend(sat) {
		t.Errorf("IsWeekend(sat) should be true")
	}
	if IsWeekend(fri) {
		t.Errorf("IsWeekend(fri) should be false")
	}

	if !IsWeekday(fri) {
		t.Errorf("IsWeekday(fri) should be true")
	}
	if IsWeekday(sat) {
		t.Errorf("IsWeekday(sat) should be false")
	}
}

func TestAddBusinessDays(t *testing.T) {
	// 2023-10-20 is a Friday
	fri := time.Date(2023, 10, 20, 12, 0, 0, 0, time.UTC)

	// Add 2 business days -> should be Tuesday (10-24)
	res := AddBusinessDays(fri, 2)
	if res.Day() != 24 || res.Weekday() != time.Tuesday {
		t.Errorf("AddBusinessDays(+2) = %v, want Tuesday 24th", res)
	}

	// Subtract 1 business day from Friday -> should be Thursday (10-19)
	resBack := AddBusinessDays(fri, -1)
	if resBack.Day() != 19 || resBack.Weekday() != time.Thursday {
		t.Errorf("AddBusinessDays(-1) = %v, want Thursday 19th", resBack)
	}
}

func ExampleStartOfDay() {
	dt := time.Date(2023, 10, 18, 12, 34, 56, 0, time.UTC)
	sod := StartOfDay(dt)
	fmt.Println(sod.Format("15:04:05"))
	// Output: 00:00:00
}

func ExampleAddBusinessDays() {
	// 2023-10-20 is a Friday
	fri := time.Date(2023, 10, 20, 12, 0, 0, 0, time.UTC)

	// Add 2 business days -> Tuesday
	tue := AddBusinessDays(fri, 2)
	fmt.Println(tue.Weekday())
	// Output: Tuesday
}
