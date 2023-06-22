package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Pull(dir string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("executing git pull on %q directory: %w", dir, err)
	}
	return nil
}
