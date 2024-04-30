package builtins

import (
	"fmt"
	"io"
	"os"
)

func MakeDirectory(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Error, wrong input")
	}
	for _, dir := range args {
		if err := os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
