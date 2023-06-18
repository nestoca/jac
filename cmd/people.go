package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"k8s.io/utils/strings/slices"
	"sort"
)

func (c *Catalog) GetPeople(groupsFilter []*v1alpha1.Group, nameFilter *PatternFilter, includeInheritedGroups bool) []*v1alpha1.Person {
	var people []*v1alpha1.Person
	for _, person := range c.People {
		// Filter by group
		if len(groupsFilter) > 0 {
			found := false
			for _, group := range groupsFilter {
				if slices.Contains(person.Spec.Groups, group.Name) {
					found = true
					break
				}
			}
			if !found && includeInheritedGroups {
				for _, group := range groupsFilter {
					for _, inheritedGroup := range person.InheritedGroups {
						if group.Name == inheritedGroup.Name {
							found = true
							break
						}
					}
				}
			}
			if !found {
				continue
			}
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match(person.Name) {
			continue
		}

		people = append(people, person)
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	return people
}
