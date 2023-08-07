package live

import (
	"fmt"
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
	Values          map[string]interface{}
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

func (p *Person) LoadValues() error {
	var err error
	p.Values, err = loadValues(p.Spec.Values.Raw)
	if err != nil {
		return fmt.Errorf("loading values for person %s: %w", p.Name, err)
	}
	return nil
}

func (p *Person) GetValueOrDefault(keyPath, defaultValue string) string {
	value, ok := p.GetValue(keyPath)
	if ok {
		return value
	}
	return defaultValue
}

func (p *Person) GetValue(keyPath string) (string, bool) {
	return getValue(p.Values, keyPath)
}

func (p *Person) IsMemberOfGroup(group *Group) bool {
	for _, g := range p.Groups {
		if g.Name == group.Name {
			return true
		}
	}
	return false
}

func (p *Person) HasAsDescendant(person *Person) bool {
	for _, child := range p.Children {
		if child.Name == person.Name {
			return true
		}
		if child.HasAsDescendant(person) {
			return true
		}
	}
	return false
}

func (p *Person) IsAmongst(people []*Person) bool {
	for _, person := range people {
		if p.Name == person.Name {
			return true
		}
	}
	return false
}

func (p *Person) HasAnyOfThoseAsDescendant(people []*Person) bool {
	for _, person := range people {
		if p.HasAsDescendant(person) {
			return true
		}
	}
	return false
}
