package main

import (
	"testing"
)

// Test to check if isBlocked works correctly
func TestIsBlocked(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"cat .env", true},
		{"ls -la", false},
		{"ssh-keygen id_rsa", true},
		{"cat credentials.json", true},
	}

	for _, test := range tests {
		if got := isBlocked(test.input); got != test.expected {
			t.Errorf("isBlocked(%q) = %v, expected %v", test.input, got, test.expected)
		}
	}
}

//Test to check if isDangerous works correctly

func TestIsDangerous(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"rm -rf /system", true},
		{"ls -la", false},
		{"git reset --hard HEAD", true},
		{"git push --force origin main", true},
		{"git status", false},
	}

	for _, test := range tests {
		if got := isDangerous(test.input); got != test.expected {
			t.Errorf("isDangerous(%q) = %v, expected %v", test.input, got, test.expected)
		}
	}
}
