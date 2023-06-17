package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/olekukonko/tablewriter"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
)

func printGroups(serializer *Serializer, groups []*v1alpha1.Group, yaml bool) error {
	if yaml {
		return printGroupYaml(serializer, groups)
	} else {
		printGroupsTable(groups)
	}
	return nil
}

func printGroupsTable(groups []*v1alpha1.Group) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	table.SetHeader([]string{"NAME", "FULL NAME", "TYPE"})

	for _, obj := range groups {
		table.Append([]string{obj.Name, obj.Spec.FullName, obj.Spec.Type})
	}

	table.Render()
}

func printGroupYaml(serializer *Serializer, groups []*v1alpha1.Group) error {
	for i, group := range groups {
		if i > 0 {
			println("---")
		}
		err := printYaml(serializer, group)
		if err != nil {
			return err
		}
	}
	return nil
}

func printPeople(serializer *Serializer, people []*v1alpha1.Person, yaml bool) error {
	if yaml {
		return printPeopleYaml(serializer, people)
	} else {
		printPeopleTable(people)
	}
	return nil
}

func printPeopleTable(groups []*v1alpha1.Person) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	headers := []string{"NAME", "FIRST NAME", "LAST NAME", "GROUPS"}
	table.SetHeader(headers)

	for _, obj := range groups {
		groupNames := ""
		for _, group := range obj.Spec.Groups {
			groupNames += group + " "
		}
		table.Append([]string{obj.Name, obj.Spec.FirstName, obj.Spec.LastName, groupNames})
	}

	table.Render()
}

func printPeopleYaml(serializer *Serializer, people []*v1alpha1.Person) error {
	for i, person := range people {
		if i > 0 {
			println("---")
		}
		err := printYaml(serializer, person)
		if err != nil {
			return err
		}
	}
	return nil
}

func printYaml(serializer *Serializer, obj runtime.Object) error {
	return serializer.Serializer.Encode(obj, os.Stdout)
}
