package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test for ScanDirectory to Ignore Default Names
func TestScanDirectoryIgnoresDefaultNames(t *testing.T) {

	//creating a temp directory
	tempDir := t.TempDir()

	//defining & creating paths of files/folders which are to be created in tempDir
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

	var ignored = []string{
		".gocache",
	}
	//calling scanDirectory function to test if it ignores only Default Names or not
	lines := scanDirectory(tempDir, 2, ignored)
	output := strings.Join(lines, "\n")

	if !strings.Contains(output, "main.go") {
		t.Errorf("ScanDirectory test failed. ScanDirectory skipped non-ignored file/folder.")
	}
	if strings.Contains(output, ".gocache") {
		t.Errorf("ScanDirectory test failed. ScanDirectory didn't skip ignored file/folder.")
	}
}
