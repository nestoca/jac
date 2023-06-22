package live

import (
	"github.com/nestoca/jac/api/v1alpha1"
)

type Person struct {
	v1alpha1.Person

	Yaml                 string
	Groups               []*Group
	InheritedGroupsNames []string
	AllGroupNames        []string
	Parent               *Person
	Children             []*Person
}

func (p *Person) GetYaml() string {
	return p.Yaml
}

func (g *Person) GetDisplayName(showNames bool) string {
	if !showNames && g.Spec.FirstName != "" {
		return g.Spec.FirstName + " " + g.Spec.LastName
	}
	return g.Name
}
