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
	findFlag := ""
	treeFlag := false
	namesFlag := false
	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "List groups",
		Aliases: []string{"group"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if yamlFlag && treeFlag {
				return fmt.Errorf("cannot use both --yaml and --tree")
			}

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
			var findFilter *filtering.FindFilter
			if findFlag != "" {
				findFilter = filtering.NewFindFilter(findFlag)
			}

			// Print groups
			isFiltering := typeFlag != "" || len(args) > 0 || findFlag != ""
			printer := printing.NewPrinter(yamlFlag, treeFlag, namesFlag, isFiltering)
			return printer.PrintGroups(catalog.RootGroups, catalog.GetGroups(typeFilter, nameFilter, findFilter))
		},
	}

	cmd.Flags().StringVar(&typeFlag, "type", "", "Filter by group type")
	cmd.Flags().BoolVarP(&treeFlag, "tree", "t", false, "Print groups as a tree")
	cmd.Flags().StringVarP(&findFlag, "find", "f", "", "Find people via freeform text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&namesFlag, "show-names", "N", false, "Show identifier names instead of full names")
	return cmd
}
