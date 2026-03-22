// Package timex provides utility functions for time manipulation.
package timex

import "time"

// StartOfDay returns the time corresponding to the start of the given day (00:00:00).
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the time corresponding to the very end of the given day (23:59:59.999999999).
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns the start of the week for the given time, based on the specified start day (e.g., time.Monday).
func StartOfWeek(t time.Time, weekStart time.Weekday) time.Time {
	offset := int(t.Weekday()) - int(weekStart)
	if offset < 0 {
		offset += 7
	}
	start := t.AddDate(0, 0, -offset)
	return StartOfDay(start)
}

// IsWeekend returns true if the time falls on a Saturday or Sunday.
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	return wd == time.Saturday || wd == time.Sunday
}

// IsWeekday returns true if the time falls on a Monday through Friday.
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// AddBusinessDays adds a given number of business days (skipping weekends) to the time.
// Note: It does not skip holidays as that requires complex localization data.
func AddBusinessDays(t time.Time, days int) time.Time {
	if days == 0 {
		return t
	}

	step := 1
	if days < 0 {
		step = -1
		days = -days
	}

	for i := 0; i < days; {
		t = t.AddDate(0, 0, step)
		if IsWeekday(t) {
			i++
		}
	}
	return t
}
