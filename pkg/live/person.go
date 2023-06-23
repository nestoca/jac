package live

import (
	"github.com/nestoca/jac/api/v1alpha1"
)

type Person struct {
	v1alpha1.Person

	Yaml            string
	Groups          []*Group
	InheritedGroups []*Group
	AllGroups       []*Group
	AllGroupNames   []string
	Parent          *Person
	Children        []*Person
}

func (p *Person) GetYaml() string {
	return p.Yaml
}

func (p *Person) GetDisplayName(showNames bool) string {
	if !showNames && p.Spec.FirstName != "" {
		return p.Spec.FirstName + " " + p.Spec.LastName
	}
	return p.Name
}

func (p *Person) HasDescendant(person *Person) bool {
	for _, child := range p.Children {
		if child.Name == person.Name {
			return true
		}
		if child.HasDescendant(person) {
			return true
		}
	}
	return false
}

func (p *Person) IsContainedIn(people []*Person) bool {
	for _, person := range people {
		if p.Name == person.Name {
			return true
		}
	}
	return false
}

func (p *Person) HasAnyDescendant(people []*Person) bool {
	for _, person := range people {
		if p.HasDescendant(person) {
			return true
		}
	}
	return false
}
