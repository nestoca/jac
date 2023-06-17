package main

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"gopkg.in/godo.v2/glob"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	globFlag string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "jac",
		Short: "Just Another CLI",
	}

	rootCmd.PersistentFlags().StringVarP(&globFlag, "glob", "g", "**/*.yaml", "Glob expression for matching CRD files")

	getCmd := createGetCmd()
	getCmd.AddCommand(createGetGroupsCmd())
	rootCmd.AddCommand(getCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get groups or people",
	}

	return cmd
}

func createGetGroupsCmd() *cobra.Command {
	typeFilter := ""
	cmd := &cobra.Command{
		Use:     "groups",
		Short:   "Get groups",
		Aliases: []string{"group"},
		Run: func(cmd *cobra.Command, args []string) {
			objs, err := loadObjects(globFlag)
			if err != nil {
				fmt.Printf("Failed to load CRDs: %v\n", err)
				os.Exit(1)
			}
			printGroups(getGroups(objs, typeFilter), typeFilter == "")
		},
	}

	cmd.Flags().StringVarP(&typeFilter, "type", "t", "", "Filter by group type")
	return cmd
}

func loadObjects(globExpr string) ([]runtime.Object, error) {
	fileAssets, _, err := glob.Glob([]string{globExpr})
	if err != nil {
		return nil, fmt.Errorf("failed to find files with glob expression %s: %v", globFlag, err)
	}

	var objs []runtime.Object

	sch := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(sch)

	deserializer := serializer.NewCodecFactory(sch).UniversalDeserializer()

	for _, fileAsset := range fileAssets {
		data, err := os.ReadFile(fileAsset.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", fileAsset, err)
		}

		obj, gvk, err := deserializer.Decode(data, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to decode file %s: %v", fileAsset, err)
		}

		switch gvk.Kind {
		case "Group":
			var crdObj v1alpha1.Group
			if err := sch.Convert(obj, &crdObj, nil); err != nil {
				return nil, fmt.Errorf("failed to convert object to Group: %v", err)
			}
			objs = append(objs, &crdObj)
		case "Person":
			var crdObj v1alpha1.Person
			if err := sch.Convert(obj, &crdObj, nil); err != nil {
				return nil, fmt.Errorf("failed to convert object to Person: %v", err)
			}
			objs = append(objs, &crdObj)
		default:
			return nil, fmt.Errorf("unsupported CRD kind: %s", gvk.Kind)
		}
	}

	return objs, nil
}

func getGroups(objs []runtime.Object, typeFilter string) []*v1alpha1.Group {
	var groups []*v1alpha1.Group
	for _, obj := range objs {
		group, ok := obj.(*v1alpha1.Group)
		if !ok {
			continue
		}
		if typeFilter != "" && group.Spec.Type != typeFilter {
			continue
		}
		groups = append(groups, group)
	}
	return groups
}

func printGroups(objs []*v1alpha1.Group, showType bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	headers := []string{"NAME", "FULL NAME"}
	if showType {
		headers = append(headers, "TYPE")
	}
	table.SetHeader(headers)

	for _, obj := range objs {
		row := []string{obj.Name, obj.Spec.FullName}
		if showType {
			row = append(row, obj.Spec.Type)
		}
		table.Append(row)
	}

	table.Render()
}
