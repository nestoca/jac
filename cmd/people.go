package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"sort"
)

func (c *Catalog) GetPeople(groupsPattern *PatternFilter, nameFilter *PatternFilter, findFilter *FindFilter, immediateGroupsOnly bool) []*v1alpha1.Person {
	var people []*v1alpha1.Person
	for _, person := range c.People {
		// Filter by group
		if groupsPattern != nil {
			if immediateGroupsOnly {
				if !groupsPattern.Match(person.Spec.Groups) {
					continue
				}
			} else {
				if !groupsPattern.Match(person.AllGroupNames) {
					continue
				}
			}
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match([]string{person.Name}) {
			continue
		}

		// Filter by find filter
		if findFilter != nil &&
			!findFilter.Match([]string{
				person.Name,
				person.Spec.FirstName,
				person.Spec.LastName,
				person.Spec.Email}) {
			continue
		}

		people = append(people, person)
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	return people
}
