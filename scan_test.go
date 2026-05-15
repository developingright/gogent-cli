package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test for ScanDirectory to Ignore Default Names
func TestScanDirectoryIgnoresDefaultNames(t *testing.T) {
	tempDir := t.TempDir()

	mainFilePath := filepath.Join(tempDir, "main.go")
	goCachePath := filepath.Join(tempDir, ".gocache")

	err := os.WriteFile(mainFilePath, []byte("package main"), 0644)
	if err != nil {
		t.Fatalf("failed to write main.go: %v", err)
	}

	err = os.Mkdir(goCachePath, 0755)

	if err != nil {
		t.Fatalf("failed to create .gocache directory: %v", err)
	}

	lines := scanDirectory(tempDir, 2)
	output := strings.Join(lines, "\n")

	if !strings.Contains(output, "main.go") {
		t.Errorf("ScanDirectory test failed. ScanDirectory skipped non-ignored file/folder.")
	}
	if strings.Contains(output, ".gocache") {
		t.Errorf("ScanDirectory test failed. ScanDirectory didn't skip ignored file/folder.")
	}
}
