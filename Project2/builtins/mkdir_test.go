package builtins

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestMakeDirectory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testmkdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	tests := []struct {
		name          string
		args          []string
		expectedError bool
	}{
		{
			name:          "No arguments",
			args:          []string{},
			expectedError: true,
		},
		{
			name:          "Create single directory",
			args:          []string{filepath.Join(tempDir, "newdir1")},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			err := MakeDirectory(&out, tt.args...)
			if (err != nil) != tt.expectedError {
				t.Errorf("%s: expected error = %v, got %v", tt.name, tt.expectedError, err != nil)
			}
		})
	}
}
