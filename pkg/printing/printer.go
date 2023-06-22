package printing

import (
	"fmt"
	"github.com/nestoca/jac/pkg/live"
	"github.com/olekukonko/tablewriter"
	"github.com/silphid/ppds/tree"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
)

type Printer struct {
	yaml        bool
	tree        bool
	isFiltering bool
}

func NewPrinter(yaml, tree, isFiltering bool) *Printer {
	return &Printer{yaml, tree, isFiltering}
}

type YamlResource interface {
	GetYaml() string
}

func (p *Printer) PrintGroups(rootGroups []*live.Group, filteredGroups []*live.Group) error {
	if p.yaml {
		return p.printGroupYaml(filteredGroups)
	} else if p.tree {
		p.printGroupTree(rootGroups, filteredGroups)
	} else {
		p.printGroupsTable(filteredGroups)
	}
	return nil
}

func (p *Printer) printGroupsTable(groups []*live.Group) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	table.SetHeader([]string{"NAME", "FULL NAME", "EMAIL", "TYPE", "PARENT"})

	for _, obj := range groups {
		parent := ""
		if obj.Spec.Parent != "" {
			parent = obj.Spec.Parent
		}
		table.Append([]string{obj.Name, obj.Spec.FullName, obj.Spec.Email, obj.Spec.Type, parent})
	}

	table.Render()
}

func (p *Printer) printGroupYaml(groups []*live.Group) error {
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

func (p *Printer) PrintPeople(people []*live.Person, showGroups bool) error {
	if p.yaml {
		return p.printPeopleYaml(people)
	} else {
		p.printPeopleTable(people, showGroups)
	}
	return nil
}

func (p *Printer) printPeopleTable(groups []*live.Person, showGroups bool) {
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

func (p *Printer) printPeopleYaml(people []*live.Person) error {
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

type Node struct {
	name     string
	children []*Node
}

func (n *Node) Data() interface{} {
	return n.name
}

func (n *Node) Children() (c []tree.Node) {
	for _, child := range n.children {
		c = append(c, tree.Node(child))
	}
	return
}

func (p *Printer) printGroupTree(rootGroups []*live.Group, filteredGroups []*live.Group) {
	tree.PrintHr(p.newTreeForGroups("", rootGroups, filteredGroups))
}

func (p *Printer) newTreeForGroup(group *live.Group, filteredGroups []*live.Group, highlight bool) *Node {
	name := group.Spec.FullName
	if p.isFiltering && highlight {
		name = "\033[33m\033[1m" + name + "\033[0m"
	}
	return p.newTreeForGroups(name, group.Children, filteredGroups)
}

func (p *Printer) newTreeForGroups(name string, groups []*live.Group, filteredGroups []*live.Group) *Node {
	node := Node{name: name}
	for _, child := range groups {
		isIn := child.IsIn(filteredGroups)
		isInOrHasAnyDescendant := isIn || child.HasAnyDescendant(filteredGroups)
		if isInOrHasAnyDescendant {
			node.children = append(node.children, p.newTreeForGroup(child, filteredGroups, isIn))
		}
	}
	return &node
}
