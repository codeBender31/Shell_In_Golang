package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	var output bytes.Buffer
	err := PrintWorkingDirectory(&output)
	if err != nil {
		t.Fatalf("PrintWorkingDirectory() returned an error: %v", err)
	}

	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned an error: %v", err)
	}

	if got := output.String(); got != expected+"\n" {
		t.Errorf("PrintWorkingDirectory() = %q, want %q", got, expected+"\n")
	}
}
