package fsx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/fsx"
	"os"
	"path/filepath"
)

func ExampleExists() {
	path := filepath.Join(os.TempDir(), "non-existent-file-goutils")
	fmt.Println(fsx.Exists(path))
	// Output: false
}

func ExampleReadFileLines() {
	path := filepath.Join(os.TempDir(), "test-goutils.txt")
	_ = os.WriteFile(path, []byte("line1\nline2"), 0644)
	defer os.Remove(path)

	lines, _ := fsx.ReadFileLines(path)
	fmt.Println(lines)
	// Output: [line1 line2]
}

func ExampleEnsureDir() {
	path := filepath.Join(os.TempDir(), "test-dir-goutils")
	err := fsx.EnsureDir(path)
	fmt.Println(err == nil)
	os.RemoveAll(path)
	// Output: true
}
