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
