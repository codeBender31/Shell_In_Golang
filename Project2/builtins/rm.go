package builtins

import (
	"fmt"
	"os"
)

func RemoveFileAndOrDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Error, missing target")
	}
	for _, path := range args {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}
