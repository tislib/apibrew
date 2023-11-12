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

var PermissionResource = &model.Resource{
	Name:        "Permission",
	Namespace:   "system",
	Title:       util.Pointer("Permission"),
	Description: util.Pointer("Permission is a resource that defines the access control rules for resources for users."),
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
					ExampleValue: structpb.NewStringValue("2023-11-13T01:12:12+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2023-11-13T01:12:12+04:00"),

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
		{
			Name: "BooleanExpression",
			Properties: []*model.ResourceProperty{
				{
					Name: "and",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "or",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "not",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "equal",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "lessThan",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "greaterThan",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "lessThanOrEqual",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "greaterThanOrEqual",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "in",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "isNull",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "regexMatch",
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "PairExpression",
			Properties: []*model.ResourceProperty{
				{
					Name: "left",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "right",
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "RegexMatchExpression",
			Properties: []*model.ResourceProperty{
				{
					Name: "pattern",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "expression",
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "meta",
			},
		},
		{
			Name: "Expression",
			Properties: []*model.ResourceProperty{
				{
					Name: "property",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "value",
					Type: model.ResourceProperty_OBJECT,
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
			Required:     true,
			Immutable:    true,
			ExampleValue: structpb.NewStringValue("a39621a4-6d48-11ee-b962-0242ac120002"),

			Annotations: map[string]string{
				"SpecialProperty": "true",
				"PrimaryProperty": "true",
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
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2023-11-13T01:12:12+04:00"), "updatedOn": structpb.NewStringValue("2023-11-13T01:12:12+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:         "namespace",
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("default"),
		},
		{
			Name:         "resource",
			Type:         model.ResourceProperty_STRING,
			Length:       255,
			ExampleValue: structpb.NewStringValue("Book"),
		},
		{
			Name: "recordSelector",
			Type: model.ResourceProperty_STRUCT,
		},
		{
			Name:         "operation",
			Type:         model.ResourceProperty_ENUM,
			Length:       255,
			Required:     true,
			DefaultValue: structpb.NewStringValue("FULL"),
			ExampleValue: structpb.NewStringValue("READ"),
		},
		{
			Name: "before",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "after",
			Type: model.ResourceProperty_TIMESTAMP,
		},
		{
			Name: "user",
			Type: model.ResourceProperty_REFERENCE,
		},
		{
			Name: "role",
			Type: model.ResourceProperty_REFERENCE,
		},
		{
			Name:     "permit",
			Type:     model.ResourceProperty_ENUM,
			Length:   255,
			Required: true,
		},
		{
			Name: "localFlags",
			Type: model.ResourceProperty_OBJECT,
		},
	},

	Annotations: map[string]string{
		"EnableAudit":  "true",
		"OpenApiGroup": "meta",
	},
}
