package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var UserResource = &model.Resource{
	Name:      "User",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "user",
	},
	Properties: []*model.ResourceProperty{
		special.IdProperty,
		special.VersionProperty,
		special.AuditProperties[0],
		special.AuditProperties[1],
		special.AuditProperties[2],
		special.AuditProperties[3],
		{
			Name:     "username",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: true,
			Unique:   true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		{
			Name:     "password",
			Type:     model.ResourceProperty_STRING,
			Length:   256,
			Required: false,
		},
		{
			Name: "roles",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{
					Namespace: RoleResource.Namespace,
					Resource:  RoleResource.Name,
				},
			},
		},
		{
			Name: "securityConstraints",
			Type: model.ResourceProperty_LIST,
			Item: &model.ResourceProperty{
				Type: model.ResourceProperty_REFERENCE,
				Reference: &model.Reference{
					Namespace: "system",
					Resource:  "Permission",
				},
				BackReference: &model.BackReference{
					Property: "user",
				},
			},
			Required: false,
		},
		{
			Name: "details",
			Type: model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit:  annotations.Enabled,
		annotations.OpenApiGroup: OpenApiMeta,
	},
}
