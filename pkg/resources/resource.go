package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

var ResourcePropertyProperties = []*model.ResourceProperty{
	{
		Name:     "name",
		Mapping:  "name",
		Primary:  false,
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: true,
	},
	{
		Name:     "type",
		Mapping:  "type",
		Type:     model.ResourceProperty_INT32,
		Required: true,
	},
	{
		Name:     "typeRef",
		Mapping:  "type_ref",
		Primary:  false,
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	{
		Name:     "mapping",
		Mapping:  "mapping",
		Type:     model.ResourceProperty_STRING,
		Length:   64,
		Required: true,
	},
	{
		Name:     "primary",
		Mapping:  "primary",
		Type:     model.ResourceProperty_BOOL,
		Required: true,
	},
	{
		Name:     "required",
		Mapping:  "required",
		Type:     model.ResourceProperty_BOOL,
		Required: true,
	},
	{
		Name:     "unique",
		Mapping:  "unique",
		Type:     model.ResourceProperty_BOOL,
		Required: true,
	},
	{
		Name:     "immutable",
		Mapping:  "immutable",
		Type:     model.ResourceProperty_BOOL,
		Required: true,
	},
	{
		Name:     "length",
		Mapping:  "length",
		Type:     model.ResourceProperty_INT32,
		Required: true,
	},
	{
		Name:     "item",
		Mapping:  "item",
		Type:     model.ResourceProperty_STRUCT,
		Required: false,
		TypeRef:  util.Pointer("Property"),
	},
	{
		Name:     "properties",
		Mapping:  "properties",
		Type:     model.ResourceProperty_LIST,
		Required: true,
		Item: &model.ResourceProperty{
			Type:    model.ResourceProperty_STRUCT,
			TypeRef: util.Pointer("Property"),
		},
	},
	{
		Name:    "reference",
		Type:    model.ResourceProperty_STRUCT,
		TypeRef: util.Pointer("Reference"),
	},
	{
		Name:     "defaultValue",
		Mapping:  "default_value",
		Type:     model.ResourceProperty_OBJECT,
		Required: false,
	},
	{
		Name:    "enumValues",
		Mapping: "enum_values",
		Type:    model.ResourceProperty_LIST,
		Item: &model.ResourceProperty{
			Type: model.ResourceProperty_STRING,
		},
		Required: false,
	},
	{
		Name:     "exampleValue",
		Mapping:  "example_value",
		Type:     model.ResourceProperty_OBJECT,
		Required: false,
	},
	{
		Name:     "title",
		Mapping:  "title",
		Primary:  false,
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	{
		Name:     "description",
		Mapping:  "description",
		Primary:  false,
		Type:     model.ResourceProperty_STRING,
		Length:   256,
		Required: false,
	},
	special.AnnotationsProperty,
}

var ResourceResource = &model.Resource{
	Name:      "Resource",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "resource",
	},
	Types: []*model.ResourceSubType{
		{
			Name:       "Property",
			Properties: ResourcePropertyProperties,
		},
		{
			Name: "SubType",
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "title",
					Mapping:  "title",
					Primary:  false,
					Type:     model.ResourceProperty_STRING,
					Length:   256,
					Required: false,
				},
				{
					Name:     "description",
					Mapping:  "description",
					Primary:  false,
					Type:     model.ResourceProperty_STRING,
					Length:   256,
					Required: false,
				},
				{
					Name:     "properties",
					Type:     model.ResourceProperty_LIST,
					Required: true,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("Property"),
					},
				},
			},
		},
		{
			Name: "IndexProperty",
			Properties: []*model.ResourceProperty{
				{
					Name:     "name",
					Type:     model.ResourceProperty_STRING,
					Required: true,
				},
				{
					Name:     "order",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"UNKNOWN", "ASC", "DESC",
					},
					DefaultValue: structpb.NewStringValue("ASC"),
				},
			},
		},
		{
			Name: "Index",
			Properties: []*model.ResourceProperty{
				{
					Name: "properties",
					Type: model.ResourceProperty_LIST,
					Item: &model.ResourceProperty{
						Type:    model.ResourceProperty_STRUCT,
						TypeRef: util.Pointer("IndexProperty"),
					},
				},
				{
					Name:     "indexType",
					Type:     model.ResourceProperty_ENUM,
					Required: false,
					EnumValues: []string{
						"BTREE", "HASH",
					},
					DefaultValue: structpb.NewStringValue("BTREE"),
				},
				{
					Name:     "unique",
					Type:     model.ResourceProperty_BOOL,
					Required: false,
				},
				special.AnnotationsProperty,
			},
		},
		{
			Name: "Reference",
			Properties: []*model.ResourceProperty{
				{
					Name: "resource",
					Type: model.ResourceProperty_REFERENCE,
					Reference: &model.Reference{
						Namespace: "system",
						Resource:  "Resource",
					},
				},
				{
					Name: "cascade",
					Type: model.ResourceProperty_BOOL,
				},
				{
					Name: "backReference",
					Type: model.ResourceProperty_STRING,
				},
			},
		},
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "name",
			Mapping:  "name",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   false,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "namespace",
			Mapping:  "namespace",
			Type:     model.ResourceProperty_REFERENCE,
			Required: true,
			Reference: &model.Reference{
				Resource:  NamespaceResource.Name,
				Namespace: NamespaceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:     "virtual",
			Mapping:  "virtual",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "properties",
			Mapping:  "properties",
			Type:     model.ResourceProperty_LIST,
			Required: true,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Property"),
			},
		},
		{
			Name:     "indexes",
			Mapping:  "indexes",
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("Index"),
			},
		},
		{
			Name:     "types",
			Mapping:  "types",
			Primary:  false,
			Type:     model.ResourceProperty_LIST,
			Required: false,
			Item: &model.ResourceProperty{
				Type:    model.ResourceProperty_STRUCT,
				TypeRef: util.Pointer("SubType"),
			},
		},
		{
			Name:     "immutable",
			Mapping:  "immutable",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "abstract",
			Mapping:  "abstract",
			Primary:  false,
			Type:     model.ResourceProperty_BOOL,
			Required: true,
		},
		{
			Name:     "dataSource",
			Mapping:  "source_data_source",
			Type:     model.ResourceProperty_REFERENCE,
			Required: false,
			Reference: &model.Reference{
				Resource:  DataSourceResource.Name,
				Namespace: DataSourceResource.Namespace,
				Cascade:   false,
			},
		},
		{
			Name:     "entity",
			Mapping:  "mapping",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "catalog",
			Mapping:  "source_catalog",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "title",
			Mapping:  "title",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name:     "description",
			Mapping:  "description",
			Primary:  false,
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		special.AnnotationsProperty,
	},
	Indexes: []*model.ResourceIndex{
		{
			Properties: []*model.ResourceIndexProperty{
				{
					Name: "namespace",
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
	},
}
