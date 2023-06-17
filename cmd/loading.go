package main

import (
	"fmt"
	"github.com/nestoca/jac/api/v1alpha1"
	"gopkg.in/godo.v2/glob"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
)

func loadObjects(serializer *Serializer, globExpr string) ([]runtime.Object, error) {
	fileAssets, _, err := glob.Glob([]string{globExpr})
	if err != nil {
		return nil, fmt.Errorf("failed to find files with glob expression %s: %v", globFlag, err)
	}

	var objs []runtime.Object

	for _, fileAsset := range fileAssets {
		data, err := os.ReadFile(fileAsset.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", fileAsset, err)
		}

		obj, gvk, err := serializer.Decoder.Decode(data, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to decode file %s: %v", fileAsset, err)
		}

		switch gvk.Kind {
		case "Group":
			var crdObj v1alpha1.Group
			if err := serializer.Scheme.Convert(obj, &crdObj, nil); err != nil {
				return nil, fmt.Errorf("failed to convert object to Group: %v", err)
			}
			objs = append(objs, &crdObj)
		case "Person":
			var crdObj v1alpha1.Person
			if err := serializer.Scheme.Convert(obj, &crdObj, nil); err != nil {
				return nil, fmt.Errorf("failed to convert object to Person: %v", err)
			}
			objs = append(objs, &crdObj)
		default:
			return nil, fmt.Errorf("unsupported CRD kind: %s", gvk.Kind)
		}
	}

	return objs, nil
}
