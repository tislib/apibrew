// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var RecordResource = &model.Resource{
	Name:        "Record",
	Namespace:   "system",
	Title:       util.Pointer("Generic Record"),
	Description: util.Pointer("A generic record resource. All Apis are extended from Generic Record resource"),
	Properties: []*model.ResourceProperty{
		{
			Name:         "id",
			Type:         model.ResourceProperty_UUID,
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"SpecialProperty": "true",
				"PrimaryProperty": "true",
			},
		},
		{
			Name:     "properties",
			Type:     model.ResourceProperty_OBJECT,
			Required: true,
		},
		{
			Name: "packedProperties",
			Type: model.ResourceProperty_LIST,

			Annotations: map[string]string{
				"OpenApiHide": "true",
			},
		},
	},
	Virtual: true,

	Annotations: map[string]string{
		"RestApiDisabled": "true",
	},
}