// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package testing

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var TestExecutionResource = &model.Resource{
	Name:        "TestExecution",
	Namespace:   "testing",
	Title:       util.Pointer("Test Case"),
	Description: util.Pointer("Test Case is a test case"),
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
			Name:       "result",
			Type:       model.ResourceProperty_ENUM,
			EnumValues: []string{"SUCCESS", "FAILURE"},
		},
		{
			Name:   "logs",
			Type:   model.ResourceProperty_STRING,
			Length: 64000,
		},
		{
			Name:         "stored",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: nil,
		},
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Length:   255,
			Required: true,
			Unique:   true,
		},
		{
			Name:      "testCase",
			Type:      model.ResourceProperty_REFERENCE,
			Required:  true,
			Reference: &model.Reference{Resource: "TestCase", Namespace: "testing"},
		},
		{
			Name:         "version",
			Type:         model.ResourceProperty_INT32,
			Required:     true,
			DefaultValue: structpb.NewNumberValue(1),
			ExampleValue: structpb.NewNumberValue(1),

			Annotations: map[string]string{
				"SpecialProperty":     "true",
				"AllowEmptyPrimitive": "true",
			},
		},
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name:  "testCase",
					Order: model.Order_ORDER_UNKNOWN,
				},
				{
					Name:  "name",
					Order: model.Order_ORDER_UNKNOWN,
				},
			},
			IndexType: model.ResourceIndexType_BTREE,
			Unique:    true,
		},
	},

	Annotations: map[string]string{
		"NormalizedResource": "true",
	},
}
