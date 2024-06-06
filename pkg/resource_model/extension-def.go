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
					Item: &model.ResourceProperty{
						Name:    "",
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("BooleanExpression"),
					},
				},
				{
					Name: "or",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name:    "",
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("BooleanExpression"),
					},
				},
				{
					Name:    "not",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("BooleanExpression"),
				},
				{
					Name:    "equal",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "lessThan",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "greaterThan",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "lessThanOrEqual",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "greaterThanOrEqual",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "in",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "like",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "ilike",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "regex",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("PairExpression"),
				},
				{
					Name:    "isNull",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
				},
				{
					Name: "filters",
					Type: model.ResourceProperty_MAP,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_OBJECT,
					},
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
					Name:    "left",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
				},
				{
					Name:    "right",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Expression"),
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
					ExampleValue: structpb.NewStringValue("2024-06-06T20:13:37+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2024-06-06T20:13:37+04:00"),

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
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
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
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
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
					Name:    "functionCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("FunctionCall"),
				},
				{
					Name:    "httpCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("HttpCall"),
				},
				{
					Name:    "channelCall",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("ChannelCall"),
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
					Item: &model.ResourceProperty{
						Name:       "action",
						Type:       model.ResourceProperty_ENUM,
						EnumValues: []string{"CREATE", "UPDATE", "DELETE", "GET", "LIST", "OPERATE"},

						Annotations: map[string]string{
							"TypeName": "EventAction",
						},
					},
				},
				{
					Name:    "recordSelector",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("BooleanExpression"),
				},
				{
					Name: "namespaces",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
				},
				{
					Name: "resources",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
				},
				{
					Name: "ids",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
				},
				{
					Name: "annotations",
					Type: model.ResourceProperty_MAP,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
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
					Name:    "query",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("BooleanExpression"),
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
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
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
					Name:       "action",
					Type:       model.ResourceProperty_ENUM,
					Required:   true,
					EnumValues: []string{"CREATE", "UPDATE", "DELETE", "GET", "LIST", "OPERATE"},
				},
				{
					Name:    "recordSearchParams",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("RecordSearchParams"),
				},
				{
					Name:      "resource",
					Type:      model.ResourceProperty_REFERENCE,
					Reference: &model.Reference{Resource: "Resource", Namespace: ""},
				},
				{
					Name: "records",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name:      "",
						Type:      model.ResourceProperty_REFERENCE,
						Reference: &model.Reference{Resource: "Record", Namespace: ""},
					},
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
					Name: "annotations",
					Type: model.ResourceProperty_MAP,
					Item: &model.ResourceProperty{
						Name: "",
						Type: model.ResourceProperty_STRING,
					},
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true"), "CommonType": structpb.NewStringValue("testType")}}),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:    "error",
					Type:    model.ResourceProperty_STRUCT,
					TypeRef: util.Pointer("Error"),
				},
			},

			Annotations: map[string]string{
				"EnableAudit":  "true",
				"OpenApiGroup": "internal",
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
					Name:       "code",
					Type:       model.ResourceProperty_ENUM,
					EnumValues: []string{"UNKNOWN_ERROR", "RECORD_NOT_FOUND", "UNABLE_TO_LOCATE_PRIMARY_KEY", "INTERNAL_ERROR", "PROPERTY_NOT_FOUND", "RECORD_VALIDATION_ERROR", "RESOURCE_VALIDATION_ERROR", "AUTHENTICATION_FAILED", "ALREADY_EXISTS", "ACCESS_DENIED", "BACKEND_ERROR", "UNIQUE_VIOLATION", "REFERENCE_VIOLATION", "RESOURCE_NOT_FOUND", "UNSUPPORTED_OPERATION", "EXTERNAL_BACKEND_COMMUNICATION_ERROR", "EXTERNAL_BACKEND_ERROR", "RATE_LIMIT_ERROR"},
				},
				{
					Name: "message",
					Type: model.ResourceProperty_STRING,
				},
				{
					Name: "fields",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Name:    "",
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("ErrorField"),
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
				"SpecialProperty":     "true",
				"AllowEmptyPrimitive": "true",
			},
		},
		{
			Name:         "auditData",
			Type:         model.ResourceProperty_STRUCT,
			TypeRef:      util.Pointer("AuditData"),
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdOn": structpb.NewStringValue("2024-06-06T20:13:37+04:00"), "updatedOn": structpb.NewStringValue("2024-06-06T20:13:37+04:00"), "createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin")}}),

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
			Name:    "selector",
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("EventSelector"),
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
			TypeRef:  util.Pointer("ExternalCall"),
		},
		{
			Name: "annotations",
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Name: "",
				Type: model.ResourceProperty_STRING,
			},
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
}
