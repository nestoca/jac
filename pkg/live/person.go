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
}

func (p *Person) GetYaml() string {
	return p.Yaml
}
