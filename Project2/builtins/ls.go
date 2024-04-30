package builtins

import (
	"fmt"
	"io"
	"io/ioutil"
)

func ListDirectory(w io.Writer, args ...string) error {
	directory := "."
	if len(args) > 0 {
		directory = args[0]
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
	return nil
}
