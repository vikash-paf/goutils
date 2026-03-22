# `ptr`

The `ptr` package makes it safe and easy to work with pointers in Go, especially useful when dealing with JSON bodies or ORMs.

## Usage

```go
// Get a pointer to a primitive instantly without creating a dummy variable
p := ptr.Of("hello")

// Safely dereference a pointer. If nil, returns the zero-value ("")
val := ptr.Val(p)

// Provide a fallback value if the pointer is nil
fallback := ptr.ValOrDefault(nilPtr, "default")

// Safely check if two pointers contain the same value, even if one/both are nil
isEqual := ptr.Equal(p1, p2)
```
