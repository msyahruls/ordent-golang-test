package utils

import (
	"os"
)

func EnsureDataDirectory() error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		return os.Mkdir("data", 0755) // Create the data directory
	}
	return nil
}
