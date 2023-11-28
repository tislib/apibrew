// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourceActionResource = &model.Resource{
	Name:      "ResourceAction",
	Namespace: "system",
	Types: []*model.ResourceSubType{
		{
			Name:        "SubType",
			Title:       "Sub Type",
			Description: "Sub Type is a type that represents a sub type of a resource. It is mostly used by STRUCT type to define the properties of the struct. ",
			Properties: []*model.ResourceProperty{
				{
					Name:         "name",
					Type:         model.ResourceProperty_STRING,
					Required:     true,
					ExampleValue: structpb.NewStringValue("Book"),
				},
				{
					Name:         "title",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("Book"),
				},
				{
					Name:         "description",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("Book is a sub type of Resource. It represents a book in the system. "),
				},
				{
					Name:         "properties",
					Type:         model.ResourceProperty_MAP,
					Required:     true,
					ExampleValue: structpb.NewListValue(&structpb.ListValue{Values: []*structpb.Value{structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"name": structpb.NewStringValue("title"), "type": structpb.NewStringValue("STRING")}})}}),
				},
			},

			Annotations: map[string]string{
				"EnableAudit":     "true",
				"RestApiDisabled": "true",
				"OpenApiGroup":    "meta",
			},
		},
		{
			Name:        "Reference",
			Title:       "Reference",
			Description: "Reference is a type that represents a reference to another resource. It is used to define the reference to another resource. ",
			Properties: []*model.ResourceProperty{
				{
					Name:         "resource",
					Type:         model.ResourceProperty_REFERENCE,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"namespace": structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"name": structpb.NewStringValue("test-namespace")}}), "name": structpb.NewStringValue("Book")}}),
				},
				{
					Name: "cascade",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name:         "backReference",
					Type:         model.ResourceProperty_STRING,
					ExampleValue: structpb.NewStringValue("author"),
				},
			},

			Annotations: map[string]string{
				"EnableAudit":     "true",
				"RestApiDisabled": "true",
				"OpenApiGroup":    "meta",
			},
		},
		{
			Name:        "Property",
			Title:       "Property",
			Description: "Property is a type that represents a property of a resource. It is like an API properties or properties of class in a programming language",
			Properties: []*model.ResourceProperty{
				{
					Name:         "type",
					Type:         model.ResourceProperty_ENUM,
					Required:     true,
					ExampleValue: structpb.NewStringValue("STRING"),
				},
				{
					Name:         "typeRef",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("BookPublishingDetails"),
				},
				{
					Name:         "primary",
					Type:         model.ResourceProperty_BOOL,
					Required:     true,
					DefaultValue: nil,
				},
				{
					Name:         "required",
					Type:         model.ResourceProperty_BOOL,
					Required:     true,
					DefaultValue: nil,
				},
				{
					Name:         "unique",
					Type:         model.ResourceProperty_BOOL,
					Required:     true,
					DefaultValue: nil,
				},
				{
					Name:         "immutable",
					Type:         model.ResourceProperty_BOOL,
					Required:     true,
					DefaultValue: nil,
				},
				{
					Name:         "length",
					Type:         model.ResourceProperty_INT32,
					Required:     true,
					DefaultValue: structpb.NewNumberValue(256),
					ExampleValue: structpb.NewNumberValue(256),
				},
				{
					Name:         "item",
					Type:         model.ResourceProperty_STRUCT,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"type": structpb.NewStringValue("STRING")}}),
				},
				{
					Name:         "reference",
					Type:         model.ResourceProperty_STRUCT,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"resource": structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"namespace": structpb.NewStringValue("default"), "resource": structpb.NewStringValue("Book")}})}}),
				},
				{
					Name:         "defaultValue",
					Type:         model.ResourceProperty_OBJECT,
					ExampleValue: structpb.NewStringValue("Lord of the Rings"),
				},
				{
					Name:         "enumValues",
					Type:         model.ResourceProperty_LIST,
					ExampleValue: structpb.NewListValue(&structpb.ListValue{Values: []*structpb.Value{structpb.NewStringValue("UNKNOWN"), structpb.NewStringValue("ASC"), structpb.NewStringValue("DESC")}}),
				},
				{
					Name:         "exampleValue",
					Type:         model.ResourceProperty_OBJECT,
					ExampleValue: structpb.NewStringValue("no-book-name"),
				},
				{
					Name:         "title",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("Book Title"),
				},
				{
					Name:         "description",
					Type:         model.ResourceProperty_STRING,
					Length:       256,
					ExampleValue: structpb.NewStringValue("Book Title is a property of Book Resource. It represents the title of the book."),
				},
				{
					Name:         "annotations",
					Type:         model.ResourceProperty_MAP,
					ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"CommonType": structpb.NewStringValue("testType"), "CheckVersion": structpb.NewStringValue("true"), "IgnoreIfExists": structpb.NewStringValue("true")}}),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"RestApiDisabled": "true",
				"OpenApiGroup":    "meta",
				"EnableAudit":     "true",
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
					ExampleValue: structpb.NewStringValue("2023-11-24T23:20:38+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
				{
					Name:         "updatedOn",
					Type:         model.ResourceProperty_TIMESTAMP,
					ExampleValue: structpb.NewStringValue("2023-11-24T23:20:38+04:00"),

					Annotations: map[string]string{
						"SpecialProperty": "true",
					},
				},
			},

			Annotations: map[string]string{
				"EnableAudit":     "true",
				"RestApiDisabled": "true",
				"OpenApiGroup":    "meta",
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
			ExampleValue: structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{"createdBy": structpb.NewStringValue("admin"), "updatedBy": structpb.NewStringValue("admin"), "createdOn": structpb.NewStringValue("2023-11-24T23:20:38+04:00"), "updatedOn": structpb.NewStringValue("2023-11-24T23:20:38+04:00")}}),

			Annotations: map[string]string{
				"SpecialProperty": "true",
			},
		},
		{
			Name:     "resource",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
		},
		{
			Name:     "name",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name:   "title",
			Type:   model.ResourceProperty_STRING,
			Length: 256,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name:   "description",
			Type:   model.ResourceProperty_STRING,
			Length: 256,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name:         "internal",
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: nil,

			Annotations: map[string]string{
				"IsHclLabel": "true",
			},
		},
		{
			Name: "types",
			Type: model.ResourceProperty_LIST,
		},
		{
			Name: "input",
			Type: model.ResourceProperty_MAP,
		},
		{
			Name: "output",
			Type: model.ResourceProperty_STRUCT,
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
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name:  "resource",
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
		"RestApiDisabled": "true",
		"OpenApiGroup":    "meta",
		"EnableAudit":     "true",
	},
}
