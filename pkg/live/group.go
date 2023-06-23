package live

import "github.com/nestoca/jac/api/v1alpha1"

type Group struct {
	v1alpha1.Group

	Yaml     string `json:""`
	Parent   *Group
	Children []*Group
}

func (g *Group) GetYaml() string {
	return g.Yaml
}

func (g *Group) GetDisplayName(showNames, allowEmoji bool) string {
	if !showNames && g.Spec.FullName != "" {
		if allowEmoji && g.Spec.Emoji != "" {
			return g.Spec.Emoji + " " + g.Spec.FullName
		}
		return g.Spec.FullName
	}
	return g.Name
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

func (g *Group) IsContainedIn(groups []*Group) bool {
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
