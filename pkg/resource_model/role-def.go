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

var RoleResource = &model.Resource{
	Name:        "Role",
	Namespace:   "system",
	Title:       util.Pointer("Role"),
	Description: util.Pointer("Role is a resource that defines the access control model. It is used to give permissions to users in a grouped way"),
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
					ExampleValue: structpb.NewStringValue("2024-03-22T17:52:25+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2024-03-22T17:52:25+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
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
				"SpecialProperty":     "true",
				"AllowEmptyPrimitive": "true",
			},
		},
		{
			Name:         "auditData",
			Type:         model.ResourceProperty_STRUCT,
			TypeRef:      util.Pointer("AuditData"),
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2024-03-22T17:52:25+04:00"), "updatedOn": structpb.NewStringValue("2024-03-22T17:52:25+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name: "permissions",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Name:      "",
				Type:      model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{Resource: "Permission", Namespace: "system"},
			},
		},
		{
			Name: "details",
			Type: model.ResourceProperty_OBJECT,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "meta",
	},
}
