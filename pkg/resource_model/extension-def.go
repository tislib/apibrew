// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

var ExtensionResource = &model.Resource{
	Name:      "Extension",
	Namespace: "system",
	Types: []*model.ResourceSubType{
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
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
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
				"OpenApiGroup": "internal",
			},
		},
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
					ExampleValue: structpb.NewStringValue("2023-11-04T03:41:26+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2023-11-04T03:41:26+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"OpenApiGroup": "internal",
				"EnableAudit":  "true",
			},
		},
		{
			Name: "FunctionCall",
			Properties: []*model.ResourceProperty{
				{
					Name:     "host",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "functionName",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"OpenApiGroup": "internal",
				"EnableAudit":  "true",
			},
		},
		{
			Name: "HttpCall",
			Properties: []*model.ResourceProperty{
				{
					Name:     "uri",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "method",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"OpenApiGroup": "internal",
				"EnableAudit":  "true",
			},
		},
		{
			Name: "ChannelCall",
			Properties: []*model.ResourceProperty{
				{
					Name:     "channelKey",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "ExternalCall",
			Properties: []*model.ResourceProperty{
				{
					Name: "functionCall",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "httpCall",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "channelCall",
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "EventSelector",
			Properties: []*model.ResourceProperty{
				{
					Name: "actions",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "recordSelector",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "namespaces",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "resources",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "ids",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name:         "annotations",
					Type:         model.ResourceProperty_MAP,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType"), "CheckVersion": structpb.NewStringValue("true")}}),

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
		{
			Name: "RecordSearchParams",
			Properties: []*model.ResourceProperty{
				{
					Name: "query",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "limit",
					Type: model.ResourceProperty_INT32,
				},
				{
					Name: "offset",
					Type: model.ResourceProperty_INT32,
				},
				{
					Name: "resolveReferences",
					Type: model.ResourceProperty_LIST,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "Event",
			Properties: []*model.ResourceProperty{
				{
					Name:      "id",
					Type:      model.ResourceProperty_STRING,
					Required:  true,
					Immutable: true,
				},
				{
					Name:     "action",
					Type:     model.ResourceProperty_ENUM,
					Required: true,
				},
				{
					Name: "recordSearchParams",
					Type: model.ResourceProperty_STRUCT,
				},
				{
					Name: "actionSummary",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "actionDescription",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "resource",
					Type: model.ResourceProperty_REFERENCE,
				},
				{
					Name: "records",
					Type: model.ResourceProperty_LIST,
				},
				{
					Name: "finalizes",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name: "sync",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name: "time",
					Type: model.ResourceProperty_TIMESTAMP,
				},
				{
					Name: "total",
					Type: model.ResourceProperty_INT64,
				},
				{
					Name: "actionName",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "input",
					Type: model.ResourceProperty_OBJECT,
				},
				{
					Name: "output",
					Type: model.ResourceProperty_OBJECT,
				},
				{
					Name:         "annotations",
					Type:         model.ResourceProperty_MAP,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType")}}),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name: "error",
					Type: model.ResourceProperty_STRUCT,
				},
			},

			Annotations: map[string]string{
				"OpenApiGroup": "internal",
				"EnableAudit":  "true",
			},
		},
		{
			Name: "ErrorField",
			Properties: []*model.ResourceProperty{
				{
					Name: "recordId",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "property",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "message",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "value",
					Type: model.ResourceProperty_OBJECT,
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
			},
		},
		{
			Name: "Error",
			Properties: []*model.ResourceProperty{
				{
					Name: "code",
					Type: model.ResourceProperty_ENUM,
				},
				{
					Name: "message",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "fields",
					Type: model.ResourceProperty_LIST,
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
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"updatedOn": structpb.NewStringValue("2023-11-04T03:41:26+04:00"), "createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2023-11-04T03:41:26+04:00")}}),

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
			Name:   "description",
			Type:   model.ResourceProperty_STRING,
			Length: 1024,
		},
		{
			Name: "selector",
			Type: model.ResourceProperty_STRUCT,
		},
		{
			Name:     "order",
			Type:     model.ResourceProperty_INT32,
			Required: true,
		},
		{
			Name:     "finalizes",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "sync",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "responds",
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "call",
			Type:     model.ResourceProperty_STRUCT,
			Required: true,
		},
		{
			Name:         "annotations",
			Type:         model.ResourceProperty_MAP,
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
	},

	Annotations: map[string]string{
		"OpenApiGroup": "internal",
		"EnableAudit":  "true",
	},
}
