package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/filtering"
	"github.com/nestoca/jac/pkg/live"
	"github.com/nestoca/jac/pkg/printing"
	"github.com/spf13/cobra"
)

func newPeopleCmd() *cobra.Command {
	groupPattern := ""
	findPattern := ""
	immediateFlag := false
	formatTree := false
	formatYaml := false
	showAll := false
	showGroupColumns := false
	showNameIdentifiers := false

	cmd := &cobra.Command{
		Use:     "people",
		Short:   "List people",
		Aliases: []string{"person"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create opts
			highlightMatches := showAll && (len(args) > 0 || findPattern != "" || groupPattern != "")
			opts := printing.NewPrintOpts(formatTree, formatYaml, showAll, showGroupColumns, showNameIdentifiers, highlightMatches)

			catalog, err := live.LoadCatalog(catalogDir, catalogGlob)
			if err != nil {
				return fmt.Errorf("loading catalog: %w\n", err)
			}

			// Create filters
			nameFilter, err := filtering.NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w\n", args, err)
			}
			var groupFilter *filtering.PatternFilter
			if groupPattern != "" {
				groupFilter, err = filtering.NewPatternFilter(groupPattern)
				if err != nil {
					return fmt.Errorf("parsing group filter %q: %w\n", args, err)
				}
			}
			var findFilter *filtering.FindFilter
			if findPattern != "" {
				findFilter = filtering.NewFindFilter(findPattern)
			}

			// Print people
			printer := printing.NewPrinter(opts, catalog)
			matchingPeople := catalog.GetPeople(groupFilter, nameFilter, findFilter, immediateFlag)
			return printer.PrintPeople(matchingPeople)
		},
	}

	cmd.Flags().StringVarP(&groupPattern, "group", "g", "", "Filter people by group")
	cmd.Flags().StringVarP(&findPattern, "find", "f", "", "Find people with free-text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&formatYaml, "yaml", "y", false, "Print people as YAML")
	cmd.Flags().BoolVarP(&formatTree, "tree", "t", false, "Print people as a tree")
	cmd.Flags().BoolVarP(&showGroupColumns, "show-groups", "G", false, "Show groups for people matching filter")
	cmd.Flags().BoolVarP(&showNameIdentifiers, "show-names", "N", false, "Show identifier names instead of full names")
	cmd.Flags().BoolVarP(&showAll, "show-all", "A", false, "Show all people in tree, regardless of filter, highlighting matches")
	cmd.Flags().BoolVarP(&immediateFlag, "immediate", "i", false, "Consider only immediate groups in filter, not inherited ones")

	return cmd
}
