package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/git"
	"github.com/spf13/cobra"
)

// Create command to pull git repo at directory resolved from dirFlag
func newPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull git repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := resolveDirectory(dirFlag)
			if err != nil {
				return fmt.Errorf("resolving directory: %w", err)
			}

			return git.Pull(dir)
		},
	}

	return cmd
}
