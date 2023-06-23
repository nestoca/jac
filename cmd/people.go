package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/filtering"
	"github.com/nestoca/jac/pkg/live"
	"github.com/nestoca/jac/pkg/printing"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newPeopleCmd() *cobra.Command {
	groupFlag := ""
	findFlag := ""
	yamlFlag := false
	treeFlag := false
	showAllFlag := false
	immediateFlag := false
	showGroupsFlag := false
	showNamesFlag := false
	cmd := &cobra.Command{
		Use:     "people",
		Short:   "List people",
		Aliases: []string{"person"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate inputs
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
			nameFilter, err := filtering.NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w\n", args, err)
			}
			var groupFilter *filtering.PatternFilter
			if groupFlag != "" {
				groupFilter, err = filtering.NewPatternFilter(groupFlag)
				if err != nil {
					return fmt.Errorf("parsing group filter %q: %w\n", args, err)
				}
			}
			var findFilter *filtering.FindFilter
			if findFlag != "" {
				findFilter = filtering.NewFindFilter(findFlag)
			}

			// Print people
			isFiltering := len(args) > 0 || findFlag != "" || groupFlag != ""
			printer := printing.NewPrinter(yamlFlag, treeFlag, showNamesFlag, showAllFlag, isFiltering)
			return printer.PrintPeople(catalog.Root.People, catalog.All.People, catalog.GetPeople(groupFilter, nameFilter, findFilter, immediateFlag), showGroupsFlag)
		},
	}

	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Filter people by group")
	cmd.Flags().StringVarP(&findFlag, "find", "f", "", "Find people with free-text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&yamlFlag, "yaml", "y", false, "Print people as YAML")
	cmd.Flags().BoolVarP(&treeFlag, "tree", "t", false, "Print people as a tree")
	cmd.Flags().BoolVarP(&showGroupsFlag, "show-groups", "G", false, "Show groups for people matching filter")
	cmd.Flags().BoolVarP(&showNamesFlag, "show-names", "N", false, "Show identifier names instead of full names")
	cmd.Flags().BoolVarP(&showAllFlag, "show-all", "A", false, "Show all people in tree, regardless of filter, highlighting matches")
	cmd.Flags().BoolVarP(&immediateFlag, "immediate", "i", false, "Consider only immediate groups in filter, not inherited ones")

	return cmd
}
