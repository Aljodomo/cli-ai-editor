package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileContent(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filepath, err)
	}

	return string(data), nil
}

func DeleteFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return fmt.Errorf("error deleting file %s: %w", filepath, err)
	}

	return nil
}

func WriteFileContent(filepath string, content string) error {
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filepath, err)
	}

	return nil
}

func GetFileNamesInDirectory(rootDirectoryPath string) ([]string, error) {
	var files []string

	// Walk through the directory structure
	err := filepath.Walk(rootDirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip invisible files and directories
		if filepath.Base(path)[0] == '.' {
			return nil
		}

		// Skip directories themselves, we only want files
		if !info.IsDir() {
			// Convert absolute path to relative path from rootDirectoryPath
			relPath, err := filepath.Rel(rootDirectoryPath, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
