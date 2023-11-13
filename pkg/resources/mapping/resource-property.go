package mapping

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

func ResourcePropertyToRecord(property *model.ResourceProperty, resource *model.Resource) *model.Record {
	properties := make(map[string]*structpb.Value)

	properties["name"] = structpb.NewStringValue(property.Name)
	if property.Title != nil {
		properties["title"] = structpb.NewStringValue(*property.Title)
	}
	if property.Description != nil {
		properties["description"] = structpb.NewStringValue(*property.Description)
	}
	properties["type"] = structpb.NewStringValue(property.Type.String())

	if property.Type == model.ResourceProperty_LIST || property.Type == model.ResourceProperty_MAP {
		properties["item"] = structpb.NewStructValue(&structpb.Struct{Fields: ResourcePropertyToRecord(property.Item, resource).Properties})
	}

	if property.Type == model.ResourceProperty_STRUCT {
		properties["typeRef"] = structpb.NewStringValue(*property.TypeRef)
	}

	if property.Type == model.ResourceProperty_ENUM {
		properties["enumValues"] = structpb.NewListValue(&structpb.ListValue{Values: util.ArrayMap(property.EnumValues, func(v string) *structpb.Value {
			return structpb.NewStringValue(v)
		})})
	}

	properties["required"] = structpb.NewBoolValue(property.Required)
	properties["length"] = structpb.NewNumberValue(float64(property.Length))
	properties["unique"] = structpb.NewBoolValue(property.Unique)
	properties["immutable"] = structpb.NewBoolValue(property.Immutable)

	if property.Reference != nil {
		referenceNamespace := property.Reference.Namespace
		if referenceNamespace == "" {
			referenceNamespace = resource.Namespace
		}

		var fieldsMap = map[string]*structpb.Value{
			"resource": util.StructKv2("name", property.Reference.Resource, "namespace", map[string]interface{}{
				"name": referenceNamespace,
			}),
			"cascade": structpb.NewBoolValue(property.Reference.Cascade),
		}

		if property.BackReference != nil {
			fieldsMap["backReference"] = structpb.NewStringValue(property.BackReference.Property)
		}

		properties["reference"] = structpb.NewStructValue(&structpb.Struct{
			Fields: fieldsMap,
		})
	}

	properties["defaultValue"] = property.DefaultValue
	properties["exampleValue"] = property.ExampleValue

	properties["annotations"], _ = structpb.NewValue(convertMap(property.Annotations, func(v string) interface{} {
		return v
	}))

	MapSpecialColumnsToRecord(property, &properties)

	return &model.Record{
		Properties: properties,
	}
}

func ResourcePropertyFromRecord(record *model.Record) *model.ResourceProperty {
	if record == nil {
		return nil
	}

	var reference *model.Reference
	var backReference *model.BackReference

	if record.Properties["reference"] != nil {
		reference = &model.Reference{}
		var referenceProperties = record.Properties["reference"].GetStructValue().GetFields()
		reference.Resource = referenceProperties["resource"].GetStructValue().GetFields()["name"].GetStringValue()
		if referenceProperties["resource"].GetStructValue().GetFields()["namespace"] != nil && referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue() != nil {
			reference.Namespace = referenceProperties["resource"].GetStructValue().GetFields()["namespace"].GetStructValue().GetFields()["name"].GetStringValue()
		}

		if referenceProperties["cascade"] != nil {
			reference.Cascade = referenceProperties["cascade"].GetBoolValue()
		}

		if record.Properties["backReference"] != nil && record.Properties["backReference"].GetStringValue() != "" {
			backReference = &model.BackReference{
				Property: record.Properties["backReference"].GetStringValue(),
			}
		}
	}

	var resourceProperty = &model.ResourceProperty{
		Name:          record.Properties["name"].GetStringValue(),
		Type:          model.ResourceProperty_Type(model.ResourceProperty_Type_value[record.Properties["type"].GetStringValue()]),
		Required:      record.Properties["required"].GetBoolValue(),
		Length:        uint32(record.Properties["length"].GetNumberValue()),
		Unique:        record.Properties["unique"].GetBoolValue(),
		Immutable:     record.Properties["immutable"].GetBoolValue(),
		DefaultValue:  record.Properties["defaultValue"],
		ExampleValue:  record.Properties["exampleValue"],
		Reference:     reference,
		BackReference: backReference,
		Annotations: convertMap(record.Properties["annotations"].GetStructValue().AsMap(), func(v interface{}) string {
			return v.(string)
		}),
	}

	if record.Properties["title"] != nil {
		resourceProperty.Title = new(string)
		*resourceProperty.Title = record.Properties["title"].GetStringValue()
	}

	if record.Properties["description"] != nil {
		resourceProperty.Description = new(string)
		*resourceProperty.Description = record.Properties["description"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_LIST || resourceProperty.Type == model.ResourceProperty_MAP {
		resourceProperty.Item = ResourcePropertyFromRecord(&model.Record{
			Properties: record.Properties["item"].GetStructValue().GetFields(),
		})
	}

	if resourceProperty.Type == model.ResourceProperty_STRUCT {
		resourceProperty.TypeRef = new(string)
		*resourceProperty.TypeRef = record.Properties["typeRef"].GetStringValue()
	}

	if resourceProperty.Type == model.ResourceProperty_ENUM {
		resourceProperty.EnumValues = util.ArrayMap(record.Properties["enumValues"].GetListValue().GetValues(), func(v *structpb.Value) string {
			return v.GetStringValue()
		})
	}

	MapSpecialColumnsFromRecord(resourceProperty, &record.Properties)

	return resourceProperty
}
