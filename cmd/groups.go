package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"sort"
)

func (c *Catalog) GetGroups(typeFilter *PatternFilter, nameFilter *PatternFilter) []*v1alpha1.Group {
	var groups []*v1alpha1.Group
	for _, group := range c.Groups {
		// Filter by type
		if typeFilter != nil && !typeFilter.Match([]string{group.Spec.Type}) {
			continue
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match([]string{group.Name}) {
			continue
		}

		groups = append(groups, group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})
	return groups
}
