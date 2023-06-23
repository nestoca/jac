package printing

import (
	"fmt"
	"github.com/nestoca/jac/pkg/live"
	"github.com/olekukonko/tablewriter"
	"github.com/silphid/ppds/tree"
	"os"
)

func (p *Printer) PrintPeople(rootPeople []*live.Person, allPeople []*live.Person, filteredPeople []*live.Person, showGroups bool) error {
	if p.yaml {
		return p.printPeopleYaml(filteredPeople)
	} else if p.tree {
		p.printPeopleTree(rootPeople, filteredPeople)
	} else {
		p.printPeopleTable(allPeople, filteredPeople, showGroups)
	}
	return nil
}

func (p *Printer) printPeopleTable(allPeople []*live.Person, filteredPeople []*live.Person, showGroups bool) {
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

	// Show all?
	people := filteredPeople
	if p.showAll {
		people = allPeople
	}

	for _, person := range people {
		// Concatenate group names
		groupNames := ""
		for _, group := range person.Spec.Groups {
			groupNames += group + " "
		}

		// Concatenate inherited group names
		inheritedGroupNames := ""
		for _, group := range person.InheritedGroupsNames {
			inheritedGroupNames += group + " "
		}

		// Build row
		row := []string{person.Name, person.Spec.FirstName, person.Spec.LastName, person.Spec.Email}
		if showGroups {
			row = append(row, groupNames)
			row = append(row, inheritedGroupNames)
		}

		// Highlight
		if p.isFiltering && p.showAll {
			if person.IsContainedIn(filteredPeople) {
				row = highlightAll(row)
			}
		}

		// Add row to table
		table.Append(row)
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

func (p *Printer) printPeopleTree(rootPeople []*live.Person, filteredPeople []*live.Person) {
	tree.PrintHr(p.newTreeForPeople("", rootPeople, filteredPeople, 1))
}

func (p *Printer) newTreeForPerson(person *live.Person, filteredPeople []*live.Person, highlight bool, depth int) *Node {
	if depth > recursionLimit {
		panic(fmt.Sprintf("cyclic people parent references detected for person %s", person.Name))
	}
	name := person.GetDisplayName(p.showNames)
	if p.isFiltering && highlight {
		name = "\033[33m\033[1m" + name + "\033[0m"
	}
	return p.newTreeForPeople(name, person.Children, filteredPeople, depth)
}

func (p *Printer) newTreeForPeople(name string, people []*live.Person, filteredPeople []*live.Person, depth int) *Node {
	node := &Node{name: name}
	for _, person := range people {
		isMatching := person.IsContainedIn(filteredPeople)
		isIncluded := p.showAll || isMatching || person.HasAnyDescendant(filteredPeople)
		if isIncluded {
			node.children = append(node.children, p.newTreeForPerson(person, filteredPeople, isMatching, depth+1))
		}
	}
	return node
}
