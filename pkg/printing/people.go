package printing

import (
	"fmt"
	"github.com/nestoca/jac/pkg/live"
	"github.com/olekukonko/tablewriter"
	"github.com/silphid/ppds/tree"
	"os"
)

func (p *Printer) PrintPeople(matchingPeople []*live.Person) error {
	switch p.opts.Format {
	case FormatTable:
		p.printPeopleTable(matchingPeople)
	case FormatYAML:
		return p.printPeopleYaml(matchingPeople)
	case FormatTree:
		p.printPeopleTree(matchingPeople)
	default:
		return fmt.Errorf("unsupported format: %d", p.opts.Format)
	}
	return nil
}

func (p *Printer) printPeopleTable(matchingPeople []*live.Person) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	headers := []string{"NAME", "FIRST NAME", "LAST NAME", "EMAIL"}
	if !p.opts.HideGroupColumns {
		headers = append(headers, "GROUPS")
		headers = append(headers, "INHERITED GROUPS")
	}
	table.SetHeader(headers)

	// Show all?
	people := matchingPeople
	if p.opts.ShowAll {
		people = p.catalog.All.People
	}

	for _, person := range people {
		// Concatenate group names
		groupNames := ""
		for _, group := range person.Groups {
			groupNames += group.GetDisplayName(p.opts.ShowIdentifierNames, false) + " "
		}

		// Concatenate inherited group names
		inheritedGroupNames := ""
		for _, group := range person.InheritedGroups {
			inheritedGroupNames += group.GetDisplayName(p.opts.ShowIdentifierNames, false) + " "
		}

		// Build row
		row := []string{person.Name, person.Spec.FirstName, person.Spec.LastName, person.Spec.Email}
		if !p.opts.HideGroupColumns {
			row = append(row, groupNames)
			row = append(row, inheritedGroupNames)
		}

		// Highlight
		if p.opts.HighlightMatches {
			if person.IsAmongst(matchingPeople) {
				row = highlightAll(row)
			}
		}

		// Add row to table
		table.Append(row)
	}

	table.Render()
	p.printCount(len(matchingPeople))
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

func (p *Printer) printPeopleTree(matchingPeople []*live.Person) {
	tree.PrintHrn(p.newTreeForPeople("", p.catalog.Root.People, matchingPeople, 1))
	p.printCount(len(matchingPeople))
}

func (p *Printer) newTreeForPerson(person *live.Person, matchingPeople []*live.Person, isHighlighted bool, depth int) *Node {
	if depth > recursionLimit {
		panic(fmt.Sprintf("cyclic people parent references detected for person %s", person.Name))
	}
	name := person.GetDisplayName(p.opts.ShowIdentifierNames)
	if p.opts.HighlightMatches && isHighlighted {
		name = highlight(name)
	}
	return p.newTreeForPeople(name, person.Children, matchingPeople, depth)
}

func (p *Printer) newTreeForPeople(name string, people []*live.Person, matchingPeople []*live.Person, depth int) *Node {
	node := &Node{name: name}
	for _, person := range people {
		isMatching := person.IsAmongst(matchingPeople)
		isIndirectlyMatching := p.opts.ShowAll || isMatching || person.HasAnyOfThoseAsDescendant(matchingPeople)
		if isIndirectlyMatching {
			node.children = append(node.children, p.newTreeForPerson(person, matchingPeople, isMatching, depth+1))
		}
	}
	return node
}
