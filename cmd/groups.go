package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/filtering"
	"github.com/nestoca/jac/pkg/live"
	"github.com/nestoca/jac/pkg/printing"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newGroupsCmd() *cobra.Command {
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
			catalog := live.NewCatalog()
			err = catalog.Load(glob)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
			}

			// Create filters
			typeFilter, err := filtering.NewPatternFilter(typeFlag)
			if err != nil {
				return fmt.Errorf("parsing type filter %q: %w", typeFlag, err)
			}
			nameFilter, err := filtering.NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w", args, err)
			}

			// Print groups
			printer := printing.NewPrinter(yamlFlag)
			return printer.PrintGroups(catalog.GetGroups(typeFilter, nameFilter))
		},
	}

	cmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Filter by group type")
	return cmd
}
