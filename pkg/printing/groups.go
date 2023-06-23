package printing

import (
	"fmt"
	"github.com/nestoca/jac/pkg/live"
	"github.com/olekukonko/tablewriter"
	"github.com/silphid/ppds/tree"
	"os"
)

func (p *Printer) PrintGroups(matchingGroups []*live.Group) error {
	switch p.opts.Format {
	case FormatTable:
		p.printGroupsTable(matchingGroups)
	case FormatYAML:
		return p.printGroupsYaml(matchingGroups)
	case FormatTree:
		p.printGroupsTree(matchingGroups)
	default:
		return fmt.Errorf("unsupported format: %d", p.opts.Format)
	}
	return nil
}

func (p *Printer) printGroupsTable(matchingGroups []*live.Group) {
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
	groups := matchingGroups
	if p.opts.ShowAll {
		groups = p.catalog.All.Groups
	}

	for _, group := range groups {
		parent := ""
		if group.Spec.Parent != "" {
			parent = group.Spec.Parent
		}

		// Build row
		row := []string{group.Name, group.GetDisplayName(false, false), group.Spec.Email, group.Spec.Type, parent}

		// Highlight
		if p.opts.HighlightMatches {
			if group.IsContainedIn(matchingGroups) {
				row = highlightAll(row)
			}
		}

		// Add row
		table.Append(row)
	}

	table.Render()
}

func (p *Printer) printGroupsYaml(groups []*live.Group) error {
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

func (p *Printer) printGroupsTree(matchingGroups []*live.Group) {
	tree.PrintHr(p.newTreeForGroups("", p.catalog.Root.Groups, matchingGroups))
}

func (p *Printer) newTreeForGroup(group *live.Group, matchingGroups []*live.Group, isMatching bool) *Node {
	name := group.GetDisplayName(p.opts.ShowIdentifierNames, false)
	if p.opts.HighlightMatches && isMatching {
		name = highlight(name)
	}
	return p.newTreeForGroups(name, group.Children, matchingGroups)
}

func (p *Printer) newTreeForGroups(name string, groups []*live.Group, matchingGroups []*live.Group) *Node {
	node := Node{name: name}
	for _, child := range groups {
		isMatching := child.IsContainedIn(matchingGroups)
		isIndirectlyMatching := p.opts.ShowAll || isMatching || child.HasAnyDescendant(matchingGroups)
		if isIndirectlyMatching {
			node.children = append(node.children, p.newTreeForGroup(child, matchingGroups, isMatching))
		}
	}
	return &node
}
