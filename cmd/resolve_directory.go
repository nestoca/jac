package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func resolveDirectory(dir string) (string, error) {
	if dir != "" {
		return dir, nil
	}

	// Resolve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home directory: %w", err)
	}

	// Check for directory in home
	dir = filepath.Join(home, ".jac/repo")
	if _, err := os.Stat(dir); err == nil {
		return dir, nil
	}

	// Assume current directory
	dir, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting current directory: %w", err)
	}
	return dir, nil
}
