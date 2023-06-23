package cmd

import (
	"github.com/spf13/cobra"
)

var (
	catalogGlob string
	catalogDir  string
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jac",
		Short: "CLI tool for managing people and groups as Infrastructure as Code",
	}

	rootCmd.PersistentFlags().StringVar(&catalogGlob, "glob", "**/*.yaml", "Glob expression for matching CRD files")
	rootCmd.PersistentFlags().StringVarP(&catalogDir, "dir", "d", "", "Directory to search for CRD files (defaults to ~/.jac/repo)")

	rootCmd.AddCommand(newGroupsCmd())
	rootCmd.AddCommand(newPeopleCmd())
	rootCmd.AddCommand(newPullCmd())
	rootCmd.AddCommand(newVersionCmd(version))
	return rootCmd
}
