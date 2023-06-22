package live

import (
	"github.com/nestoca/jac/api/v1alpha1"
)

type Person struct {
	v1alpha1.Person

	Yaml                 string   `json:""`
	Groups               []*Group `json:""`
	InheritedGroupsNames []string `json:""`
	AllGroupNames        []string `json:""`
}

func (p *Person) GetYaml() string {
	return p.Yaml
}
