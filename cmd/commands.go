package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	globFlag string
	yamlFlag bool
)

func createRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jac",
		Short: "Just Another CLI",
	}

	rootCmd.PersistentFlags().StringVar(&globFlag, "glob", "**/*.yaml", "Glob expression for matching CRD files")
	rootCmd.PersistentFlags().BoolVarP(&yamlFlag, "yaml", "y", false, "Output in YAML format")

	getCmd := createGetCmd()
	getCmd.AddCommand(createGetGroupsCmd())
	getCmd.AddCommand(createGetPeopleCmd())
	rootCmd.AddCommand(getCmd)
	return rootCmd
}

func createGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get groups or people",
	}

	return cmd
}

func createGetGroupsCmd() *cobra.Command {
	typeFlag := ""
	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "Get groups",
		Aliases: []string{"group"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog := NewCatalog()
			err := catalog.Load(globFlag)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
			}
			typeFilter, err := NewPatternFilter(typeFlag)
			if err != nil {
				return fmt.Errorf("parsing type filter %q: %w", typeFlag, err)
			}
			nameFilter, err := NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("parsing name filter %q: %w", args, err)
			}
			printer := NewPrinter(catalog.Serializer, yamlFlag)
			return printer.PrintGroups(catalog.GetGroups(typeFilter, nameFilter))
		},
	}

	cmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Filter by group type")
	return cmd
}

func createGetPeopleCmd() *cobra.Command {
	groupFlag := ""
	inheritedFlag := false
	cmd := &cobra.Command{
		Use:     "people",
		Short:   "Get people",
		Aliases: []string{"person"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			catalog := NewCatalog()
			err := catalog.Load(globFlag)
			if err != nil {
				return fmt.Errorf("loading CRDs: %w\n", err)
			}
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
			printer := NewPrinter(catalog.Serializer, yamlFlag)
			return printer.PrintPeople(catalog.GetPeople(groupFilter, nameFilter, inheritedFlag))
		},
	}

	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Filter by group")
	cmd.Flags().BoolVarP(&inheritedFlag, "inherited", "i", false, "Include inherited groups in the filter")

	return cmd
}
