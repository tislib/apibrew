package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourceActionResource = &model.Resource{
	Name:      "ResourceAction",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource_action",
	},
	Types: []*model.ResourceSubType{
		SubTypeType,
		ReferenceType,
		PropertyType,
		special.AuditDataSubType,
	},
	Properties: map[string]*model.ResourceProperty{
		"id":        special.IdProperty,
		"version":   special.VersionProperty,
		"auditData": special.AuditProperty,
		"resource": {
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				Resource:  ResourceResource.Name,
				Namespace: ResourceResource.Namespace,
				Cascade:   false,
			},
		},
		"name": {
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"title": {
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"description": {
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"internal": {
			Type:         model.ResourceProperty_BOOL,
			Required:     true,
			DefaultValue: structpb.NewBoolValue(false),
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		"types": {
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("SubType"),
			},
		},
		"input": {
			Type: model.ResourceProperty_MAP,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
		},
		"output": {
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("Property"),
		},
		"annotations": special.AnnotationsProperty,
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "resource",
				},
				{
					Name: "name",
				},
			},
			Unique: true,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:     annotations.Enabled,
		annotations.RestApiDisabled: annotations.Enabled,
		annotations.OpenApiGroup:    OpenApiMeta,
	},
}
