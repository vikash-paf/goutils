# `str`

The `str` package provides string manipulation and formatting utilities.

## Usage

```go
str.IsBlank("   ") // true

str.Reverse("こんにちは") // はちにんこ

str.Truncate("hello world", 8, "...") // hello...

str.ToCamelCase("hello_world") // helloWorld
str.ToSnakeCase("HelloWorld") // hello_world
```

## String Algorithms

### JaroWinkler
Calculates the Jaro-Winkler similarity between two strings, returning a score from 0.0 (no match) to 1.0 (exact match). Excellent for phonetic matching and typo-tolerance.
```go
sim := str.JaroWinkler("martha", "marhta") // ~0.96

if sim > 0.9 {
    fmt.Println("High similarity!")
}
```
