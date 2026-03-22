package set_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/set"
	"sort"
)

func ExampleSet() {
	s := set.New(1, 2, 3)
	s.Add(4)
	fmt.Println(s.Contains(2))
	fmt.Println(len(s))
	// Output:
	// true
	// 4
}

func ExampleUnion() {
	s1 := set.New(1, 2)
	s2 := set.New(2, 3)
	u := set.Union(s1, s2)
	fmt.Println(len(u)) // 1, 2, 3
	// Output: 3
}

func ExampleIntersection() {
	s1 := set.New(1, 2)
	s2 := set.New(2, 3)
	i := set.Intersection(s1, s2)
	values := i.Values()
	sort.Ints(values)
	fmt.Println(values)
	// Output: [2]
}
