# `tuple`

The `tuple` package introduces a standard two-value generic `Pair`.

## Usage

```go
p := tuple.NewPair(1, "hello")
fmt.Println(p.Left, p.Right)

// Combine two slices into pairs
names := []string{"Alice", "Bob"}
ages := []int{30, 25}
zipped := tuple.Zip(names, ages)
// [{Alice, 30}, {Bob, 25}]

// Split them back apart
unzippedNames, unzippedAges := tuple.Unzip(zipped)
```
