package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/config"
	"github.com/nestoca/jac/pkg/git"
	"github.com/spf13/cobra"
)

// Create command to pull git repo at directory resolved from catalogDir
func newPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull git repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig(catalogDir)
			if err != nil {
				return fmt.Errorf("loading config: %w\n", err)
			}
			return git.Pull(cfg.Dir)
		},
	}

	return cmd
}
