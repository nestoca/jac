package live

import "github.com/nestoca/jac/api/v1alpha1"

type Group struct {
	v1alpha1.Group

	Yaml     string
	Parent   *Group
	Children []*Group
}

func (g *Group) GetYaml() string {
	return g.Yaml
}

func (g *Group) HasDescendant(group *Group) bool {
	for _, child := range g.Children {
		if child.Name == group.Name {
			return true
		}
		if child.HasDescendant(group) {
			return true
		}
	}
	return false
}

func (g *Group) IsIn(groups []*Group) bool {
	for _, group := range groups {
		if g.Name == group.Name {
			return true
		}
	}
	return false
}

func (g *Group) HasAnyDescendant(groups []*Group) bool {
	for _, group := range groups {
		if g.HasDescendant(group) {
			return true
		}
	}
	return false
}
