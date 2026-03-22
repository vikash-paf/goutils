package id_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/id"
)

func ExampleUUID() {
	uuid := id.UUID()
	fmt.Println(len(uuid))
	// Output: 36
}

func ExampleRandomString() {
	str := id.RandomString(10)
	fmt.Println(len(str))
	// Output: 10
}

func ExampleNanoID() {
	nid := id.NanoID(21)
	fmt.Println(len(nid))
	// Output: 21
}
