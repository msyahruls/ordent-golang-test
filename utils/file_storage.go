package utils

import (
	"encoding/json"
	"os"
)

// ReadFromFile reads data from a file and unmarshals it into the provided destination.
func ReadFromFile(filePath string, dest interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist yet; treat as empty.
		}
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(dest)
}

// WriteToFile marshals the provided data and writes it to a file.
func WriteToFile(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	return encoder.Encode(data)
}
