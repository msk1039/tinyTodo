package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

// Storage struct for handling JSON file
type Storage[T any] struct {
	FilePath string
}

// getStoragePath returns the appropriate path based on the OS
func getStoragePath() string {
	var dir string

	if runtime.GOOS == "windows" {
		dir = filepath.Join(os.Getenv("APPDATA"), "todo")
	} else if runtime.GOOS == "darwin" { // macOS
		dir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "todo")
	} else { // Linux
		dir = filepath.Join(os.Getenv("HOME"), ".config", "todo")
	}

	// Ensure directory exists
	os.MkdirAll(dir, 0755)

	return filepath.Join(dir, "todos.json")
}

// NewStorage initializes Storage with the correct file path
func NewStorage[T any]() *Storage[T] {
	return &Storage[T]{FilePath: getStoragePath()}
}

// Save writes the data to the JSON file
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, fileData, 0644)
}

// Load reads the JSON file into the given data structure
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FilePath)
	if os.IsNotExist(err) {
		return nil // No existing file is fine, return without error
	} else if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}