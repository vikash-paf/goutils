package mathx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/mathx"
)

func ExampleSum() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(mathx.Sum(nums))
	// Output: 10
}

func ExampleAverage() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(mathx.Average(nums))
	// Output: 2.5
}

func ExampleMinBy() {
	type User struct {
		Name string
		Age  int
	}
	users := []User{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
	youngest := mathx.MinBy(users, func(u User) int { return u.Age })
	fmt.Println(youngest.Name)
	// Output: Bob
}

func ExampleMaxBy() {
	type User struct {
		Name string
		Age  int
	}
	users := []User{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
	oldest := mathx.MaxBy(users, func(u User) int { return u.Age })
	fmt.Println(oldest.Name)
	// Output: Charlie
}

func ExampleClamp() {
	fmt.Println(mathx.Clamp(150, 0, 100))
	fmt.Println(mathx.Clamp(-50, 0, 100))
	fmt.Println(mathx.Clamp(50, 0, 100))
	// Output:
	// 100
	// 0
	// 50
}
