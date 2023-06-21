package main

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"gopkg.in/godo.v2/glob"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"sort"
	"strings"
)

type Catalog struct {
	Groups []*v1alpha1.Group
	People []*v1alpha1.Person

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
		return fmt.Errorf("matching files with glob expression %s: %w", globFlag, err)
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
			sort.Strings(crdObj.Spec.Parents)
			crdObj.Yaml = strings.TrimSpace(string(data))
			c.Groups = append(c.Groups, &crdObj)
		case "Person":
			var crdObj v1alpha1.Person
			if err := c.Scheme.Convert(obj, &crdObj, nil); err != nil {
				return fmt.Errorf("converting object to Person: %w", err)
			}
			sort.Strings(crdObj.Spec.Groups)
			crdObj.Yaml = strings.TrimSpace(string(data))
			c.People = append(c.People, &crdObj)
		default:
			return fmt.Errorf("unsupported CRD kind: %s", gvk.Kind)
		}
	}

	// Resolve group parents for all groups
	for _, group := range c.Groups {
		sort.Strings(group.Spec.Parents)
		for _, parentName := range group.Spec.Parents {
			parent := c.GetGroup(parentName)
			if parent == nil {
				return fmt.Errorf("group %s has parent %s which does not exist", group.Name, parentName)
			}
			group.Parents = append(group.Parents, parent)
		}
	}

	// Resolve groups for all people
	for _, person := range c.People {
		sort.Strings(person.Spec.Groups)
		for _, groupName := range person.Spec.Groups {
			group := c.GetGroup(groupName)
			if group == nil {
				return fmt.Errorf("person %s is in group %s which does not exist", person.Name, groupName)
			}
			person.Groups = append(person.Groups, group)
		}
		c.resolveInheritedGroups(person)
	}

	return nil
}

const recursionLimit = 50

func (c *Catalog) resolveInheritedGroups(person *v1alpha1.Person) {
	var inheritedGroupNames []string
	for _, group := range person.Groups {
		inheritedGroupNames = append(inheritedGroupNames, c.resolveInheritedGroupsRecursively(person, group, 1)...)
	}
	sort.Strings(inheritedGroupNames)
	person.InheritedGroupsNames = inheritedGroupNames
	person.AllGroupNames = append(person.Spec.Groups, inheritedGroupNames...)
}

func (c *Catalog) resolveInheritedGroupsRecursively(person *v1alpha1.Person, group *v1alpha1.Group, depth int) []string {
	var inheritedGroupNames []string
	if depth > recursionLimit {
		panic(fmt.Sprintf("cyclic group parent references detected for person %s", person.Name))
	}
	for _, group := range group.Parents {
		inheritedGroupNames = append(inheritedGroupNames, group.Name)
		inheritedGroupNames = append(inheritedGroupNames, c.resolveInheritedGroupsRecursively(person, group, depth+1)...)
	}
	return inheritedGroupNames
}

func (c *Catalog) GetGroup(name string) *v1alpha1.Group {
	for _, group := range c.Groups {
		if group.Name == name {
			return group
		}
	}
	return nil
}
