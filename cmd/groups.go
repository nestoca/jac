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
	yamlFlag := false
	treeFlag := false
	showAllFlag := false
	showNamesFlag := false
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
				return fmt.Errorf("loading catalog: %w\n", err)
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
			printer := printing.NewPrinter(yamlFlag, treeFlag, showNamesFlag, showAllFlag, isFiltering)
			return printer.PrintGroups(catalog.Root.Groups, catalog.All.Groups, catalog.GetGroups(typeFilter, nameFilter, findFilter))
		},
	}

	cmd.Flags().StringVarP(&typeFlag, "type", "T", "", "Filter by group type")
	cmd.Flags().BoolVarP(&yamlFlag, "yaml", "y", false, "Print groups as YAML")
	cmd.Flags().BoolVarP(&treeFlag, "tree", "t", false, "Print groups as a tree")
	cmd.Flags().StringVarP(&findFlag, "find", "f", "", "Find people with free-text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&showNamesFlag, "show-names", "N", false, "Show identifier names instead of full names")
	cmd.Flags().BoolVarP(&showAllFlag, "show-all", "A", false, "Show all groups in tree, regardless of filter, highlighting matches")
	return cmd
}
