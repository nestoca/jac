package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	globFlag string
	yamlFlag bool
	dirFlag  string
)

func createRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jac",
		Short: "Tool to manage people and groups as Infrastructure as Code",
	}

	rootCmd.PersistentFlags().StringVar(&globFlag, "glob", "**/*.yaml", "Glob expression for matching CRD files")
	rootCmd.PersistentFlags().BoolVarP(&yamlFlag, "yaml", "y", false, "Output in YAML format")
	rootCmd.PersistentFlags().StringVarP(&dirFlag, "dir", "d", "", "Directory to search for CRD files (defaults to ~/.jac/repo)")

	rootCmd.AddCommand(createGetGroupsCmd())
	rootCmd.AddCommand(createGetPeopleCmd())
	rootCmd.AddCommand(createPullCmd())
	return rootCmd
}

func createGetGroupsCmd() *cobra.Command {
	typeFlag := ""
	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "Get groups",
		Aliases: []string{"group"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Resolve directory
			dir, err := resolveDirectory(dirFlag)
			if err != nil {
				return fmt.Errorf("resolving directory: %w", err)
			}

			// Determine glob
			glob := filepath.Join(dir, globFlag)

			// Load catalog
			catalog := NewCatalog()
			err = catalog.Load(glob)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
			}

			// Create filters
			typeFilter, err := NewPatternFilter(typeFlag)
			if err != nil {
				return fmt.Errorf("parsing type filter %q: %w", typeFlag, err)
			}
			nameFilter, err := NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w", args, err)
			}

			// Print groups
			printer := NewPrinter(catalog.Serializer, yamlFlag)
			return printer.PrintGroups(catalog.GetGroups(typeFilter, nameFilter))
		},
	}

	cmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Filter by group type")
	return cmd
}

func createGetPeopleCmd() *cobra.Command {
	groupFlag := ""
	immediateFlag := false
	cmd := &cobra.Command{
		Use:     "people",
		Short:   "Get people",
		Aliases: []string{"person"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Resolve directory
			dir, err := resolveDirectory(dirFlag)
			if err != nil {
				return fmt.Errorf("resolving directory: %w", err)
			}

			// Determine glob
			glob := filepath.Join(dir, globFlag)

			// Load catalog
			catalog := NewCatalog()
			err = catalog.Load(glob)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
			}

			// Create filters
			nameFilter, err := NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w\n", args, err)
			}
			var groupFilter *PatternFilter
			if groupFlag != "" {
				groupFilter, err = NewPatternFilter(groupFlag)
				if err != nil {
					return fmt.Errorf("parsing group filter %q: %w\n", args, err)
				}
			}

			// Print people
			printer := NewPrinter(catalog.Serializer, yamlFlag)
			return printer.PrintPeople(catalog.GetPeople(groupFilter, nameFilter, immediateFlag))
		},
	}

	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Filter by group")
	cmd.Flags().BoolVarP(&immediateFlag, "immediate", "i", false, "Consider only immediate groups in filter, not inherited ones")

	return cmd
}

// Create command to pull git repo at directory resolved from dirFlag
func createPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull git repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := resolveDirectory(dirFlag)
			if err != nil {
				return fmt.Errorf("resolving directory: %w", err)
			}

			return gitPull(dir)
		},
	}

	return cmd
}
