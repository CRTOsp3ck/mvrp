package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

type IFSUtil interface {
	ReadFileFromPath(path string) ([]byte, error)
	WriteToFile(jsonStr, relativePath string, createDirs bool) error
	FindProjectRoot(knownFileOrDir string) (string, error)
}

type FSUtil struct{}

func (f *FSUtil) ReadFileFromPath(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FSUtil) WriteToFile(jsonStr, relativePath string, createDirs bool) error {
	// Dynamically determine the base path of the project
	basePath, err := f.FindProjectRoot("go.mod")
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}

	// Append the /schema directory to the base path
	path := filepath.Join(basePath, relativePath)

	if createDirs {
		// Extract the directory part of the path
		dir := filepath.Dir(path)
		// Create all directories in the path, if they don't exist
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directories: %w", err)
		}
	}

	// Open or create the file for writing
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write the JSON string to the file
	_, err = file.WriteString(jsonStr)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	// fmt.Printf("Wrote JSON to %s\n", path)
	return nil
}

func (f *FSUtil) FindProjectRoot(knownFileOrDir string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}
	for {
		// Check if the known file or directory exists in the current directory
		path := filepath.Join(currentDir, knownFileOrDir)
		if _, err := os.Stat(path); err == nil {
			// Found the known file or directory, return the current directory as the root
			return currentDir, nil
		}

		// Move up one directory
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// Reached the filesystem root without finding the known file or directory
			return "", fmt.Errorf("project root not found from start directory: %s", currentDir)
		}
		currentDir = parentDir
	}
}
