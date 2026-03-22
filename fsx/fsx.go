// Package fsx provides high-level file system utilities that simplify common tasks
// like checking for file existence, ensuring directories exist, and reading/writing
// files as slices of strings.
//
// Usage:
//
//	if fsx.Exists("config.json") {
//	    lines, _ := fsx.ReadFileLines("config.json")
//	}
//	_ = fsx.EnsureDir("logs/app")
package fsx

import (
	"bufio"
	"io"
	"os"
)

// Exists returns true if the specified file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// EnsureDir checks if a directory exists and creates it (including parents) if it doesn't.
func EnsureDir(path string) error {
	if Exists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}

// ReadFileLines reads a file and returns its content as a slice of strings,
// with each element representing a line.
func ReadFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteFileLines writes a slice of strings to a file, each element as a new line.
func WriteFileLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
