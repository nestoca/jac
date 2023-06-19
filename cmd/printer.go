package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/olekukonko/tablewriter"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"os"
)

type Printer struct {
	serializer *json.Serializer
	yaml       bool
}

func NewPrinter(serializer *json.Serializer, yaml bool) *Printer {
	return &Printer{serializer: serializer, yaml: yaml}
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

func (p *Printer) PrintPeople(people []*v1alpha1.Person) error {
	if p.yaml {
		return p.printPeopleYaml(people)
	} else {
		p.printPeopleTable(people)
	}
	return nil
}

func (p *Printer) printPeopleTable(groups []*v1alpha1.Person) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	headers := []string{"NAME", "FIRST NAME", "LAST NAME", "EMAIL", "GROUPS", "INHERITED GROUPS"}
	table.SetHeader(headers)

	for _, obj := range groups {
		// Concatenate group names
		groupNames := ""
		for _, group := range obj.Spec.Groups {
			groupNames += group + " "
		}

		// Concatenate inherited group names
		inheritedGroupNames := ""
		for _, group := range obj.InheritedGroups {
			inheritedGroupNames += group + " "
		}

		table.Append([]string{obj.Name, obj.Spec.FirstName, obj.Spec.LastName, obj.Spec.Email, groupNames, inheritedGroupNames})
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
	return p.serializer.Encode(obj, os.Stdout)
}
