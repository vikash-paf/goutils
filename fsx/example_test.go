package fsx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/fsx"
	"os"
)

func ExampleExists() {
	fmt.Println(fsx.Exists("/tmp/non-existent-file"))
	// Output: false
}

func ExampleReadFileLines() {
	path := "/tmp/test.txt"
	_ = os.WriteFile(path, []byte("line1\nline2"), 0644)
	defer os.Remove(path)

	lines, _ := fsx.ReadFileLines(path)
	fmt.Println(lines)
	// Output: [line1 line2]
}

func ExampleEnsureDir() {
	path := "/tmp/test-dir"
	err := fsx.EnsureDir(path)
	fmt.Println(err == nil)
	os.RemoveAll(path)
	// Output: true
}
