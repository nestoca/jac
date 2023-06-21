package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func gitPull(dir string) error {
	// Ensure directory is a git repo
	dotGitDir := filepath.Join(dir, ".git")
	if _, err := os.Stat(dotGitDir); err != nil {
		return fmt.Errorf("directory %q does appear to be a git repo", dir)
	}

	// Exec git pull
	cmd := exec.Command("git", "pull")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
