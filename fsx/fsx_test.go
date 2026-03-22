package fsx

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFsx(t *testing.T) {
	tmpDir := os.TempDir()
	testDir := filepath.Join(tmpDir, "fsx_test")
	testFile := filepath.Join(testDir, "test.txt")

	defer os.RemoveAll(testDir)

	if err := EnsureDir(testDir); err != nil {
		t.Fatalf("Failed to ensure dir: %v", err)
	}

	if !Exists(testDir) {
		t.Error("testDir should exist")
	}

	lines := []string{"hello", "world"}
	if err := WriteFileLines(testFile, lines); err != nil {
		t.Fatalf("Failed to write lines: %v", err)
	}

	readLines, err := ReadFileLines(testFile)
	if err != nil {
		t.Fatalf("Failed to read lines: %v", err)
	}

	if len(readLines) != 2 || readLines[0] != "hello" || readLines[1] != "world" {
		t.Errorf("Read lines mismatch: %v", readLines)
	}
}

