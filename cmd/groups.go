package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func getGroups(objs []runtime.Object, typeFilter *PatternFilter, nameFilter *PatternFilter) []*v1alpha1.Group {
	var groups []*v1alpha1.Group
	for _, obj := range objs {
		group, ok := obj.(*v1alpha1.Group)
		if !ok {
			continue
		}

		// Filter by type
		if typeFilter != nil && !typeFilter.Match(group.Spec.Type) {
			continue
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match(group.Name) {
			continue
		}

		groups = append(groups, group)
	}
	return groups
}
