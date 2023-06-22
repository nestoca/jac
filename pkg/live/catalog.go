package live

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/nestoca/jac/pkg/filtering"
	"gopkg.in/godo.v2/glob"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"sort"
	"strings"
)

type Catalog struct {
	Groups []*Group
	People []*Person

	RootGroups []*Group

	Scheme *runtime.Scheme
}

func NewCatalog() *Catalog {
	sch := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(sch)

	return &Catalog{
		Scheme: sch,
	}
}

func (c *Catalog) Load(globExpr string) error {
	// Find all files matching the glob expression
	fileAssets, _, err := glob.Glob([]string{globExpr})
	if err != nil {
		return fmt.Errorf("matching files with glob expression %s: %w", globExpr, err)
	}

	// Load all matched files
	decoder := serializer.NewCodecFactory(c.Scheme).UniversalDeserializer()
	for _, fileAsset := range fileAssets {
		data, err := os.ReadFile(fileAsset.Path)
		if err != nil {
			return fmt.Errorf("reading file %s: %w", fileAsset.Path, err)
		}

		obj, gvk, err := decoder.Decode(data, nil, nil)
		if err != nil {
			return fmt.Errorf("decoding file %s: %w", fileAsset.Path, err)
		}

		switch gvk.Kind {
		case "Group":
			var crdObj v1alpha1.Group
			if err := c.Scheme.Convert(obj, &crdObj, nil); err != nil {
				return fmt.Errorf("converting object to Group: %w", err)
			}
			group := &Group{
				Group: crdObj,
			}
			group.Yaml = strings.TrimSpace(string(data))
			c.Groups = append(c.Groups, group)
		case "Person":
			var crdObj v1alpha1.Person
			if err := c.Scheme.Convert(obj, &crdObj, nil); err != nil {
				return fmt.Errorf("converting object to Person: %w", err)
			}
			sort.Strings(crdObj.Spec.Groups)
			person := &Person{
				Person: crdObj,
			}
			person.Yaml = strings.TrimSpace(string(data))
			c.People = append(c.People, person)
		default:
			return fmt.Errorf("unsupported CRD kind: %s", gvk.Kind)
		}
	}

	// Resolve group parents for all groups
	for _, group := range c.Groups {
		if group.Spec.Parent != "" {
			parent := c.GetGroup(group.Spec.Parent)
			if parent == nil {
				return fmt.Errorf("group %s's parent does not exist: %s", group.Name, group.Spec.Parent)
			}
			group.Parent = parent
			parent.Children = append(parent.Children, group)
		} else {
			c.RootGroups = append(c.RootGroups, group)
		}
	}

	// Resolve groups for all people
	for _, person := range c.People {
		sort.Strings(person.Spec.Groups)
		for _, groupName := range person.Spec.Groups {
			group := c.GetGroup(groupName)
			if group == nil {
				return fmt.Errorf("person %s's group does not exist: %s", person.Name, groupName)
			}
			person.Groups = append(person.Groups, group)
		}
		c.resolveInheritedGroups(person)
	}

	return nil
}

const recursionLimit = 50

func (c *Catalog) resolveInheritedGroups(person *Person) {
	var inheritedGroupNames []string
	for _, group := range person.Groups {
		inheritedGroupNames = append(inheritedGroupNames, c.resolveInheritedGroupsRecursively(person, group, 1)...)
	}
	sort.Strings(inheritedGroupNames)
	person.InheritedGroupsNames = inheritedGroupNames
	person.AllGroupNames = append(person.Spec.Groups, inheritedGroupNames...)
}

func (c *Catalog) resolveInheritedGroupsRecursively(person *Person, group *Group, depth int) []string {
	var inheritedGroupNames []string
	if depth > recursionLimit {
		panic(fmt.Sprintf("cyclic group parent references detected for person %s", person.Name))
	}
	if group.Parent != nil {
		inheritedGroupNames = append(inheritedGroupNames, group.Parent.Name)
		inheritedGroupNames = append(inheritedGroupNames, c.resolveInheritedGroupsRecursively(person, group.Parent, depth+1)...)
	}
	return inheritedGroupNames
}

func (c *Catalog) GetGroup(name string) *Group {
	for _, group := range c.Groups {
		if group.Name == name {
			return group
		}
	}
	return nil
}

func (c *Catalog) GetPeople(groupsPattern *filtering.PatternFilter, nameFilter *filtering.PatternFilter, findFilter *filtering.FindFilter, immediateGroupsOnly bool) []*Person {
	var people []*Person
	for _, person := range c.People {
		// Filter by group
		if groupsPattern != nil {
			if immediateGroupsOnly {
				if !groupsPattern.Match(person.Spec.Groups) {
					continue
				}
			} else {
				if !groupsPattern.Match(person.AllGroupNames) {
					continue
				}
			}
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match([]string{person.Name}) {
			continue
		}

		// Filter by find filter
		if findFilter != nil &&
			!findFilter.Match([]string{
				person.Name,
				person.Spec.FirstName,
				person.Spec.LastName,
				person.Spec.Email}) {
			continue
		}

		people = append(people, person)
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	return people
}

func (c *Catalog) GetGroups(typeFilter *filtering.PatternFilter, nameFilter *filtering.PatternFilter, findFilter *filtering.FindFilter) []*Group {
	var groups []*Group
	for _, group := range c.Groups {
		// Filter by type
		if typeFilter != nil && !typeFilter.Match([]string{group.Spec.Type}) {
			continue
		}

		// Filter by names
		if nameFilter != nil && !nameFilter.Match([]string{group.Name}) {
			continue
		}

		// Filter by find filter
		if findFilter != nil &&
			!findFilter.Match([]string{
				group.Name,
				group.Spec.FullName,
				group.Spec.Email}) {
			continue
		}

		groups = append(groups, group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})
	return groups
}
