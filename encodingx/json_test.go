package encodingx

import (
	"fmt"
	"testing"
)

func TestEncodingx(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	u := User{Name: "Alice", Age: 30}

	// MustMarshalJSON
	jsonStr := MustMarshalJSON(u)
	if jsonStr != `{"name":"Alice","age":30}` {
		t.Errorf("Unexpected JSON: %s", jsonStr)
	}

	// ToMap
	m, err := ToMap(u)
	if err != nil {
		t.Fatalf("ToMap failed: %v", err)
	}
	if m["name"] != "Alice" || fmt.Sprintf("%v", m["age"]) != "30" {
		t.Errorf("Unexpected map: %v", m)
	}
}

func ExampleMustMarshalJSON() {
	m := map[string]int{"a": 1}
	fmt.Println(MustMarshalJSON(m))
	// Output: {"a":1}
}
