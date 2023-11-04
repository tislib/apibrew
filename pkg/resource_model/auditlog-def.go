// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

var AuditLogResource = &model.Resource{
	Name:      "AuditLog",
	Namespace: "system",
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
			Name:     "namespace",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name:     "resource",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
		},
		{
			Name:     "recordId",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,

			Annotations: map[string]string{
				"SourceDef": "record_id",
			},
		},
		{
			Name:     "time",
			Type:     model.ResourceProperty_TIMESTAMP,
			Required: true,
		},
		{
			Name:     "username",
			Type:     model.ResourceProperty_STRING,
			Required: true,
		},
		{
			Name:     "operation",
			Type:     model.ResourceProperty_ENUM,
			Required: true,
		},
		{
			Name: "properties",
			Type: model.ResourceProperty_OBJECT,
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
	Immutable: true,

	Annotations: map[string]string{
		"OpenApiGroup":     "internal",
		"BypassExtensions": "true",
	},
}
