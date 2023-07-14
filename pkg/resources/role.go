package resources

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources/special"
	"github.com/apibrew/apibrew/pkg/service/annotations"
)

var RoleResource = &model.Resource{
	Name:      "Role",
	Namespace: "system",
	SourceConfig: &model.ResourceSourceConfig{
		DataSource: "system",
		Entity:     "role",
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
			Unique:   true,
			Annotations: map[string]string{
				annotations.IsHclLabel: annotations.Enabled,
			},
		},
		special.SecurityConstraintsProperty,
		{
			Name: "details",

			Mapping: "details",
			Type:    model.ResourceProperty_OBJECT,
		},
	},
	Annotations: map[string]string{
		annotations.EnableAudit: "true",
	},
}