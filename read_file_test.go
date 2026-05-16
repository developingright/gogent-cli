package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// Test to check if read_file does it's job perfectly in normal circumstances
func TestHappyPath(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "test.txt")

	err := os.WriteFile(testFilePath, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("failed to write test.txt: %v", err)
	}

	content, err := readFile(testFilePath)
	if err != nil {
		t.Errorf("couldn't read test file: %v", err)
	}

	if content != "1: hello world" {
		t.Errorf("content of test file doesn't match")
	}
}

// Test to check if read_file returns appropriate error for file not found case
func TestFileDoesntExist(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "test.txt")

	_, err := readFile(testFilePath)
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("read_file doesn't return appropriate error for file not found")
	}
}

// Test to check if read_file returns appropriate error for file too large case
func TestFileTooLarge(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "test.txt")

	err := os.WriteFile(testFilePath, bytes.Repeat([]byte("a"), 100*1024+1), 0644)
	if err != nil {
		t.Fatalf("failed to write test.txt: %v", err)
	}

	_, err = readFile(testFilePath)
	if err == nil || err.Error() != "file too large" {
		t.Errorf("read_file doesn't return appropriate error for file too large")
	}
}
