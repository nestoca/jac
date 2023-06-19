package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"sort"
)

func (c *Catalog) GetPeople(groupsPattern Pattern, nameFilter *PatternFilter, includeInheritedGroups bool) []*v1alpha1.Person {
	var people []*v1alpha1.Person
	for _, person := range c.People {
		// Filter by group
		if groupsPattern != nil &&
			!groupsPattern.Match(person.Spec.Groups) &&
			(!includeInheritedGroups || groupsPattern.Match(person.InheritedGroups)) {
			continue
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match([]string{person.Name}) {
			continue
		}

		people = append(people, person)
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	return people
}
