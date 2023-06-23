package printing

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
)

type YamlResource interface {
	GetYaml() string
}

func (p *Printer) printYaml(obj runtime.Object) error {
	if yamlResource, ok := obj.(YamlResource); ok {
		fmt.Println(yamlResource.GetYaml())
		return nil
	} else {
		return fmt.Errorf("object is not a YamlResource")
	}
}
