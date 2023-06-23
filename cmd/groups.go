package cmd

import (
	"fmt"
	"github.com/nestoca/jac/pkg/filtering"
	"github.com/nestoca/jac/pkg/live"
	"github.com/nestoca/jac/pkg/printing"
	"github.com/spf13/cobra"
)

func newGroupsCmd() *cobra.Command {
	typePattern := ""
	findPattern := ""
	formatTree := false
	formatYaml := false
	showAll := false
	showGroupColumns := false
	showNameIdentifiers := false

	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "List groups",
		Aliases: []string{"group"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create opts
			highlightMatches := showAll && (len(args) > 0 || findPattern != "" || typePattern != "")
			opts := printing.NewPrintOpts(formatTree, formatYaml, showAll, showGroupColumns, showNameIdentifiers, highlightMatches)

			catalog, err := live.LoadCatalog(catalogDir, catalogGlob)
			if err != nil {
				return fmt.Errorf("loading catalog: %w\n", err)
			}

			// Create filters
			typeFilter, err := filtering.NewPatternFilter(typePattern)
			if err != nil {
				return fmt.Errorf("parsing type filter %q: %w", typePattern, err)
			}
			nameFilter, err := filtering.NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w", args, err)
			}
			var findFilter *filtering.FindFilter
			if findPattern != "" {
				findFilter = filtering.NewFindFilter(findPattern)
			}

			// Print groups
			printer := printing.NewPrinter(opts, catalog)
			matchingGroups := catalog.GetGroups(typeFilter, nameFilter, findFilter)
			return printer.PrintGroups(matchingGroups)
		},
	}

	cmd.Flags().StringVarP(&typePattern, "type", "T", "", "Filter by group type")
	cmd.Flags().BoolVarP(&formatYaml, "yaml", "y", false, "Print groups as YAML")
	cmd.Flags().BoolVarP(&formatTree, "tree", "t", false, "Print groups as a tree")
	cmd.Flags().StringVarP(&findPattern, "find", "f", "", "Find people with free-text search in their first or last name, email or name identifier")
	cmd.Flags().BoolVarP(&showNameIdentifiers, "show-names", "N", false, "Show identifier names instead of full names")
	cmd.Flags().BoolVarP(&showAll, "show-all", "A", false, "Show all groups in tree, regardless of filter, highlighting matches")
	return cmd
}
