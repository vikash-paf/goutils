# `timex`

The `timex` package extends the standard `time` package with common business-logic timing functions.

## Usage

```go
now := time.Now()

// Standard boundaries
start := timex.StartOfDay(now)
end := timex.EndOfDay(now)
weekStart := timex.StartOfWeek(now, time.Monday)

// Business logic
isWknd := timex.IsWeekend(now)

// Add days, skipping Saturdays and Sundays
deadline := timex.AddBusinessDays(now, 5)
```
