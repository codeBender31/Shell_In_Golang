package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Implementation of Echo command
func Echo(out io.Writer, args ...string) error {
	//Format the string and print it back out
	_, err := fmt.Fprintln(out, strings.Join(args, " "))
	return err
}
