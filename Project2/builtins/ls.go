package builtins

import (
	"fmt"
	"io"
	"os"
)

func ListDirectory(w io.Writer, args ...string) error {
	directory := "."
	if len(args) > 0 {
		directory = args[0]
	}
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
	return nil
}
