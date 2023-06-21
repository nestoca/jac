package main

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/olekukonko/tablewriter"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
)

type Printer struct {
	yaml bool
}

func NewPrinter(yaml bool) *Printer {
	return &Printer{yaml: yaml}
}

type YamlResource interface {
	GetYaml() string
}

func (p *Printer) PrintGroups(groups []*v1alpha1.Group) error {
	if p.yaml {
		return p.printGroupYaml(groups)
	} else {
		p.printGroupsTable(groups)
	}
	return nil
}

func (p *Printer) printGroupsTable(groups []*v1alpha1.Group) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	table.SetHeader([]string{"NAME", "FULL NAME", "EMAIL", "TYPE", "PARENTS"})

	for _, obj := range groups {
		parentNames := ""
		for _, group := range obj.Spec.Parents {
			parentNames += group + " "
		}
		table.Append([]string{obj.Name, obj.Spec.FullName, obj.Spec.Email, obj.Spec.Type, parentNames})
	}

	table.Render()
}

func (p *Printer) printGroupYaml(groups []*v1alpha1.Group) error {
	for i, group := range groups {
		if i > 0 {
			println("---")
		}
		err := p.printYaml(group)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Printer) PrintPeople(people []*v1alpha1.Person, showGroups bool) error {
	if p.yaml {
		return p.printPeopleYaml(people)
	} else {
		p.printPeopleTable(people, showGroups)
	}
	return nil
}

func (p *Printer) printPeopleTable(groups []*v1alpha1.Person, showGroups bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	headers := []string{"NAME", "FIRST NAME", "LAST NAME", "EMAIL"}
	if showGroups {
		headers = append(headers, "GROUPS")
		headers = append(headers, "INHERITED GROUPS")
	}
	table.SetHeader(headers)

	for _, obj := range groups {
		// Concatenate group names
		groupNames := ""
		for _, group := range obj.Spec.Groups {
			groupNames += group + " "
		}

		// Concatenate inherited group names
		inheritedGroupNames := ""
		for _, group := range obj.InheritedGroupsNames {
			inheritedGroupNames += group + " "
		}

		// Add row to table
		values := []string{obj.Name, obj.Spec.FirstName, obj.Spec.LastName, obj.Spec.Email}
		if showGroups {
			values = append(values, groupNames)
			values = append(values, inheritedGroupNames)
		}
		table.Append(values)
	}

	table.Render()
}

func (p *Printer) printPeopleYaml(people []*v1alpha1.Person) error {
	for i, person := range people {
		if i > 0 {
			println("---")
		}
		err := p.printYaml(person)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Printer) printYaml(obj runtime.Object) error {
	if yamlResource, ok := obj.(YamlResource); ok {
		fmt.Println(yamlResource.GetYaml())
		return nil
	} else {
		return fmt.Errorf("object is not a YamlResource")
	}
}
