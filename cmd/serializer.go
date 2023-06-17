package main

import (
	"github.com/nestoca/jac/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

type Serializer struct {
	Scheme     *runtime.Scheme
	Serializer *json.Serializer
	Decoder    runtime.Decoder
}

func NewSerializer() *Serializer {
	sch := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(sch)

	return &Serializer{
		Scheme:     sch,
		Serializer: json.NewSerializerWithOptions(json.DefaultMetaFactory, sch, sch, json.SerializerOptions{Yaml: true}),
		Decoder:    serializer.NewCodecFactory(sch).UniversalDeserializer(),
	}
}
