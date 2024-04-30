package builtins

import (
	"fmt"
	"io"
	"os"
)

func PrintWorkingDirectory(w io.Writer) error {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, workingDirectory)
	return err
}
