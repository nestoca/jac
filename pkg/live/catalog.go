package live

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"github.com/nestoca/jac/pkg/filtering"
	"gopkg.in/godo.v2/glob"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Catalog struct {
	// All groups and people in the catalog
	All GroupsAndPeople `json:"all,omitempty"`

	// Groups and people with no parents
	Root GroupsAndPeople `json:"root,omitempty"`

	Scheme *runtime.Scheme `json:""`
}

type GroupsAndPeople struct {
	Groups []*Group  `json:"groups,omitempty"`
	People []*Person `json:"people,omitempty"`
}

func NewCatalog() *Catalog {
	sch := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(sch)

	return &Catalog{
		Scheme: sch,
	}
}

func LoadCatalog(dir, glob string) (*Catalog, error) {
	fullGlob := filepath.Join(dir, glob)

	c := NewCatalog()
	if err := c.Load(fullGlob); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Catalog) Load(globExpr string) error {
	// Find all files matching the glob expression
	fileAssets, _, err := glob.Glob([]string{globExpr})
	if err != nil {
		return fmt.Errorf("matching files with glob expression %s: %w", globExpr, err)
	}
	if len(fileAssets) == 0 {
		return fmt.Errorf("no files found matching glob expression %q", globExpr)
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
			err = group.LoadValues()
			if err != nil {
				return err
			}
			c.All.Groups = append(c.All.Groups, group)
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
			err = person.LoadValues()
			if err != nil {
				return err
			}
			c.All.People = append(c.All.People, person)
		default:
			return fmt.Errorf("unsupported CRD kind: %s", gvk.Kind)
		}
	}

	// Resolve group parents for all groups
	for _, group := range c.All.Groups {
		if group.Spec.Parent != "" {
			parent := c.GetGroup(group.Spec.Parent)
			if parent == nil {
				return fmt.Errorf("group %s's parent does not exist: %s", group.Name, group.Spec.Parent)
			}
			group.Parent = parent
			parent.Children = append(parent.Children, group)
		} else {
			c.Root.Groups = append(c.Root.Groups, group)
		}
	}

	// Sort root groups
	sort.Slice(c.Root.Groups, func(i, j int) bool {
		return c.Root.Groups[i].Name < c.Root.Groups[j].Name
	})

	// Resolve groups, parent and children for all people
	for _, person := range c.All.People {
		// Parent
		if person.Spec.Parent != "" {
			parent := c.GetPerson(person.Spec.Parent)
			if parent == nil {
				return fmt.Errorf("person %s's parent does not exist: %s", person.Name, person.Spec.Parent)
			}
			person.Parent = parent
			parent.Children = append(parent.Children, person)
		} else {
			c.Root.People = append(c.Root.People, person)
		}

		// Groups
		sort.Strings(person.Spec.Groups)
		for _, groupName := range person.Spec.Groups {
			group := c.GetGroup(groupName)
			if group == nil {
				return fmt.Errorf("person %s's group does not exist: %s", person.Name, groupName)
			}
			person.Groups = append(person.Groups, group)
			group.Members = append(group.Members, person)
			group.AllMembers = append(group.AllMembers, person)
		}
		c.resolveInheritedGroups(person)
		for _, group := range person.InheritedGroups {
			group.IndirectMembers = append(group.IndirectMembers, person)
		}
	}

	// Sort people children
	for _, person := range c.All.People {
		sort.Slice(person.Children, func(i, j int) bool {
			return person.Children[i].Name < person.Children[j].Name
		})
	}

	return nil
}

const recursionLimit = 50

func (c *Catalog) resolveInheritedGroups(person *Person) {
	var inheritedGroups []*Group
	for _, group := range person.Groups {
		inheritedGroups = append(inheritedGroups, c.resolveInheritedGroupsRecursively(person, group, 1)...)
	}
	sort.Slice(inheritedGroups, func(i, j int) bool {
		return inheritedGroups[i].Name < inheritedGroups[j].Name
	})
	person.InheritedGroups = inheritedGroups
	person.AllGroups = append(person.Groups, inheritedGroups...)
	for _, group := range person.AllGroups {
		person.AllGroupNames = append(person.AllGroupNames, group.Name)
	}

}

func (c *Catalog) resolveInheritedGroupsRecursively(person *Person, group *Group, depth int) []*Group {
	var inheritedGroups []*Group
	if depth > recursionLimit {
		panic(fmt.Sprintf("cyclic group parent references detected for person %s", person.Name))
	}
	if group.Parent != nil {
		inheritedGroups = append(inheritedGroups, group.Parent)
		inheritedGroups = append(inheritedGroups, c.resolveInheritedGroupsRecursively(person, group.Parent, depth+1)...)
	}
	return inheritedGroups
}

func (c *Catalog) GetGroup(name string) *Group {
	for _, group := range c.All.Groups {
		if group.Name == name {
			return group
		}
	}
	return nil
}

func (c *Catalog) GetPerson(name string) *Person {
	for _, person := range c.All.People {
		if person.Name == name {
			return person
		}
	}
	return nil
}

func (c *Catalog) GetPeople(groupsPattern *filtering.PatternFilter, nameFilter *filtering.PatternFilter, findFilter *filtering.FindFilter, immediateGroupsOnly bool) []*Person {
	var people []*Person
	for _, person := range c.All.People {
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
	for _, group := range c.All.Groups {
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
