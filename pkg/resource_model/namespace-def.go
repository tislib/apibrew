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

var NamespaceResource = &model.Resource{
	Name:      "Namespace",
	Namespace: "system",
	Types: []*model.ResourceSubType{
		{
			Name:        "AuditData",
			Title:       "Audit Data",
			Description: "Audit Data is a type that represents the audit data of a resource/record. ",
			Properties: []*model.ResourceProperty{
				{
					Name:         "createdBy",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("admin"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedBy",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("admin"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "createdOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					Immutable:    true,
					ExampleValue: structpb.NewStringValue("2024-04-25T00:16:31+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2024-04-25T00:16:31+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
	},
	Properties: []*model.ResourceProperty{
		{
			Name:         "id",
			Type:         model.ResourceProperty_UUID,
			Primary:      true,
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:         "version",
			Type:         model.ResourceProperty_INT32,
			Required:     true,
			DefaultValue: structpb.NewNumberValue(1),
			ExampleValue: structpb.NewNumberValue(1),

			Annotations: map[string]string{
				"AllowEmptyPrimitive": "true",
				"SpecialProperty":     "true",
			},
		},
		{
			Name:         "auditData",
			Type:         model.ResourceProperty_STRUCT,
			TypeRef:      util.Pointer("AuditData"),
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"updatedOn": structpb.NewStringValue("2024-04-25T00:16:31+04:00"), "createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2024-04-25T00:16:31+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:      "name",
			Type:      model.ResourceProperty_STRING,
			Length:    256,
			Required:  true,
			Unique:    true,
			Immutable: true,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name:   "description",
			Type:   model.ResourceProperty_STRING,
			Length: 256,
		},
		{
			Name: "details",
			Type: model.ResourceProperty_OBJECT,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "internal",
	},
}
