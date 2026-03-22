package mathx

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([]int) = %v, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4.0 {
		t.Errorf("Sum([]float64) = %v, want 4.0", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum(empty) = %v, want 0", got)
	}
}

func TestAverage(t *testing.T) {
	if got := Average([]int{1, 2, 3, 4}); got != 2.5 {
		t.Errorf("Average() = %v, want 2.5", got)
	}
	if got := Average([]int{}); got != 0.0 {
		t.Errorf("Average(empty) = %v, want 0", got)
	}
}

func TestMinMaxBy(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	users := []user{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 40},
	}

	minAge := MinBy(users, func(u user) int { return u.age })
	if minAge == nil || minAge.name != "Bob" {
		t.Errorf("MinBy() = %v, want Bob", minAge)
	}

	maxAge := MaxBy(users, func(u user) int { return u.age })
	if maxAge == nil || maxAge.name != "Charlie" {
		t.Errorf("MaxBy() = %v, want Charlie", maxAge)
	}

	var empty []user
	if got := MinBy(empty, func(u user) int { return u.age }); got != nil {
		t.Errorf("MinBy(empty) = %v, want nil", got)
	}
}

func TestClamp(t *testing.T) {
	if got := Clamp(5, 1, 10); got != 5 {
		t.Errorf("Clamp(5, 1, 10) = %v, want 5", got)
	}
	if got := Clamp(0, 1, 10); got != 1 {
		t.Errorf("Clamp(0, 1, 10) = %v, want 1", got)
	}
	if got := Clamp(15, 1, 10); got != 10 {
		t.Errorf("Clamp(15, 1, 10) = %v, want 10", got)
	}
}

func ExampleClamp() {
	val1 := Clamp(-5, 0, 100)
	val2 := Clamp(50, 0, 100)
	val3 := Clamp(150, 0, 100)
	
	fmt.Println(val1, val2, val3)
	// Output: 0 50 100
}
