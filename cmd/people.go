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
	immediateFlag := false
	showGroupsFlag := false
	cmd := &cobra.Command{
		Use:     "people",
		Short:   "List people",
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
			catalog := live.NewCatalog()
			err = catalog.Load(glob)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
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
			isFiltering := len(args) > 0 || findFlag != ""
			printer := printing.NewPrinter(yamlFlag, false, isFiltering)
			return printer.PrintPeople(catalog.GetPeople(groupFilter, nameFilter, findFilter, immediateFlag), showGroupsFlag)
		},
	}

	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Filter people by group")
	cmd.Flags().StringVarP(&findFlag, "find", "f", "", "Find people via freeform text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&showGroupsFlag, "show-groups", "G", false, "Show groups for people matching filter")
	cmd.Flags().BoolVarP(&immediateFlag, "immediate", "i", false, "Consider only immediate groups in filter, not inherited ones")

	return cmd
}
