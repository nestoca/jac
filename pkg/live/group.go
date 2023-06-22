package live

import "github.com/nestoca/jac/api/v1alpha1"

type Group struct {
	v1alpha1.Group

	Yaml    string   `json:""`
	Parents []*Group `json:""`
}

func (g *Group) GetYaml() string {
	return g.Yaml
}
