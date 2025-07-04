package tools

import (
	"fmt"
	"os"
)

func DeleteFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file at the path: %v is not found", path)
	}
	return os.Remove(path)
}
