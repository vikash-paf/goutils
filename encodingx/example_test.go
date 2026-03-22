package encodingx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/encodingx"
)

func ExampleMustMarshalJSON() {
	m := map[string]int{"a": 1}
	fmt.Println(encodingx.MustMarshalJSON(m))
	// Output: {"a":1}
}

func ExampleToMap() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	u := User{Name: "Alice", Age: 30}
	m, _ := encodingx.ToMap(u)
	fmt.Println(m["name"])
	// Output: Alice
}

func ExamplePrettyPrint() {
	m := map[string]int{"a": 1}
	// Note: output check is tricky due to indentation, but we test it works.
	fmt.Println(encodingx.PrettyPrint(m))
	// Output:
	// {
	//   "a": 1
	// }
}
