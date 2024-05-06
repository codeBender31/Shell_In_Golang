package builtins

import (
	"bytes"
	"os"
	"testing"
)

func TestListDirectory(t *testing.T) {
	var outPut bytes.Buffer
	tempDirectory, err := os.MkdirTemp("", "testls")
	if err != nil {
		t.Fatalf("Creation of temp has failed", err)
	}
	defer os.RemoveAll(tempDirectory)

	testingFile := "testingFile.txt"
	if err := os.WriteFile(tempDirectory+"/"+testingFile, []byte("trial content"), 0666); err != nil {
		t.Fatalf("TestingFile creation failed. %v", err)
	}
	if err := ListDirectory(&outPut, tempDirectory); err != nil {
		t.Fatalf("ListDirectory() returned error: %v", err)
	}

}
