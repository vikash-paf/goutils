package cache_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/cache"
)

func ExampleLRU() {
	c := cache.NewLRU[string, int](2)
	c.Set("alice", 30)
	c.Set("bob", 25)
	c.Set("charlie", 28) // Evicts "alice"

	_, ok := c.Get("alice")
	fmt.Println("Alice in cache:", ok)

	age, ok := c.Get("bob")
	fmt.Println("Bob age:", age, ok)
	// Output:
	// Alice in cache: false
	// Bob age: 25 true
}

func ExampleLFU() {
	c := cache.NewLFU[string, int](2)
	c.Set("alice", 30)
	c.Set("bob", 25)

	// Access "alice" to increase frequency
	c.Get("alice")

	c.Set("charlie", 28) // Evicts "bob" because it has lower frequency

	_, ok := c.Get("bob")
	fmt.Println("Bob in cache:", ok)

	age, ok := c.Get("alice")
	fmt.Println("Alice age:", age, ok)
	// Output:
	// Bob in cache: false
	// Alice age: 30 true
}
