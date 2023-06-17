package main

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/spf13/cobra"
	"os"
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
			serializer := NewSerializer()
			objs, err := loadObjects(serializer, globFlag)
			if err != nil {
				return fmt.Errorf("failed to load CRDs: %v", err)
			}
			typeFilter, err := NewPatternFilter(typeFlag)
			if err != nil {
				return fmt.Errorf("failed to parse type filter %q: %v", typeFlag, err)
			}
			nameFilter, err := NewPatternsFilter(args)
			if err != nil {
				return fmt.Errorf("failed to parse name filter %q: %v", args, err)
			}
			return printGroups(serializer, getGroups(objs, typeFilter, nameFilter), yamlFlag)
		},
	}

	cmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Filter by group type")
	return cmd
}

func createGetPeopleCmd() *cobra.Command {
	groupFlag := ""
	cmd := &cobra.Command{
		Use:     "people",
		Short:   "Get people",
		Aliases: []string{"person"},
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			serializer := NewSerializer()
			objs, err := loadObjects(serializer, globFlag)
			if err != nil {
				fmt.Printf("Failed to load CRDs: %v\n", err)
				os.Exit(1)
			}
			nameFilter, err := NewPatternsFilter(args)
			if err != nil {
				fmt.Printf("Failed to parse name filter %q: %v\n", args, err)
				os.Exit(1)
			}
			var groups []*v1alpha1.Group
			if groupFlag != "" {
				groupFilter, err := NewPatternFilter(groupFlag)
				if err != nil {
					fmt.Printf("Failed to parse name filter %q: %v\n", args, err)
					os.Exit(1)
				}
				groups = getGroups(objs, nil, groupFilter)
				if len(groups) == 0 {
					fmt.Printf("No groups found matching %q\n", groupFlag)
					os.Exit(1)
				}
			}
			return printPeople(serializer, getPeople(objs, groups, nameFilter), yamlFlag)
		},
	}

	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Filter by group")
	return cmd
}
