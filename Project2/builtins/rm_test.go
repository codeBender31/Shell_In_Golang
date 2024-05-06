package builtins

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveFileAndOrDirectory(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "testrm")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "testfile.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0666); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name        string
		args        []string
		expectedErr bool
	}{
		{
			name:        "Remove file",
			args:        []string{testFile},
			expectedErr: false,
		},
		{
			name:        "Invalid path",
			args:        []string{filepath.Join(tempDir, "nonexistent")},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RemoveFileAndOrDirectory(tt.args...)
			if (err != nil) != tt.expectedErr {
				t.Errorf("%s: expected error %v, got error %v", tt.name, tt.expectedErr, err)
			}
		})
	}
}
