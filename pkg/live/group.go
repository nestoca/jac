package live

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"strings"
)

type Group struct {
	v1alpha1.Group

	Yaml            string `json:""`
	Parent          *Group
	Children        []*Group
	Members         []*Person
	IndirectMembers []*Person
	AllMembers      []*Person
	values          map[string]interface{}
}

func (g *Group) GetYaml() string {
	return g.Yaml
}

func (g *Group) GetDisplayName(showNames, allowEmoji bool) string {
	nonBreakingSpace := "\u00a0"
	if !showNames && g.Spec.FullName != "" {
		nonBreakingFullName := strings.ReplaceAll(g.Spec.FullName, " ", nonBreakingSpace)
		if allowEmoji && g.Spec.Emoji != "" {
			return g.Spec.Emoji + nonBreakingSpace + nonBreakingFullName
		}
		return nonBreakingFullName
	}
	return g.Name
}

func (g *Group) GetValue(keyPath string) (string, bool) {
	if g.values == nil {
		var err error
		g.values, err = loadValues(g.Spec.Values.Raw)
		if err != nil {
			panic(fmt.Errorf("loading values for group %s: %w", g.Name, err))
		}
	}
	return getValue(g.values, keyPath)
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
