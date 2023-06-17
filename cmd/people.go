package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/strings/slices"
)

func getPeople(objs []runtime.Object, groupsFilter []*v1alpha1.Group, nameFilter *PatternFilter) []*v1alpha1.Person {
	var people []*v1alpha1.Person
	for _, obj := range objs {
		person, ok := obj.(*v1alpha1.Person)
		if !ok {
			continue
		}

		// Filter by group
		if len(groupsFilter) > 0 {
			found := false
			for _, group := range groupsFilter {
				if slices.Contains(person.Spec.Groups, group.Name) {
					found = true
					break
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
	return people
}
