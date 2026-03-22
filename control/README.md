# `control`

The `control` package provides flow control and error handling helpers designed to drastically reduce boilerplate in Go.

## Usage

### If (Ternary Operator)
```go
status := control.If(statusCode == 200, "OK", "Error")
```

### Must
Standardizes the "panic on error" initialization pattern.
```go
// Panics if db.Open returns an error
db := control.Must(sql.Open("postgres", dsn))
```

### Coalesce
Returns the first non-zero value.
```go
name := control.Coalesce(user.NickName, user.FullName, "Anonymous")
```

### Try
Executes a function and returns a fallback value if it errors.
```go
val := control.Try(100, func() (int, error) {
    return strconv.Atoi("invalid")
})
// val is 100
```
