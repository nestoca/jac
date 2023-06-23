package printing

import (
	"github.com/nestoca/jac/pkg/live"
	"github.com/olekukonko/tablewriter"
	"github.com/silphid/ppds/tree"
	"os"
)

func (p *Printer) PrintGroups(rootGroups []*live.Group, allGroups []*live.Group, filteredGroups []*live.Group) error {
	if p.yaml {
		return p.printGroupYaml(filteredGroups)
	} else if p.tree {
		p.printGroupTree(rootGroups, filteredGroups)
	} else {
		p.printGroupsTable(allGroups, filteredGroups)
	}
	return nil
}

func (p *Printer) printGroupsTable(allGroups []*live.Group, filteredGroups []*live.Group) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	table.SetHeader([]string{"NAME", "FULL NAME", "EMAIL", "TYPE", "PARENT"})

	// Show all?
	groups := filteredGroups
	if p.showAll {
		groups = allGroups
	}

	for _, group := range groups {
		parent := ""
		if group.Spec.Parent != "" {
			parent = group.Spec.Parent
		}

		// Build row
		row := []string{group.Name, group.GetDisplayName(false, false), group.Spec.Email, group.Spec.Type, parent}

		// Highlight
		if p.isFiltering && p.showAll {
			if group.IsContainedIn(filteredGroups) {
				row = highlightAll(row)
			}
		}

		// Add row
		table.Append(row)
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

func (p *Printer) printGroupTree(rootGroups []*live.Group, filteredGroups []*live.Group) {
	tree.PrintHr(p.newTreeForGroups("", rootGroups, filteredGroups))
}

func (p *Printer) newTreeForGroup(group *live.Group, filteredGroups []*live.Group, isHighlighted bool) *Node {
	name := group.GetDisplayName(p.showNames, false)
	if p.isFiltering && isHighlighted {
		name = highlight(name)
	}
	return p.newTreeForGroups(name, group.Children, filteredGroups)
}

func (p *Printer) newTreeForGroups(name string, groups []*live.Group, filteredGroups []*live.Group) *Node {
	node := Node{name: name}
	for _, child := range groups {
		isMatching := child.IsContainedIn(filteredGroups)
		isIncluded := p.showAll || isMatching || child.HasAnyDescendant(filteredGroups)
		if isIncluded {
			node.children = append(node.children, p.newTreeForGroup(child, filteredGroups, isMatching))
		}
	}
	return &node
}
