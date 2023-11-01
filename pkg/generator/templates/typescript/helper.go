package typescript

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/util"
	"strings"
)

func propertyName(property *model.ResourceProperty) string {
	return property.Name
}

func getRestPath(resource *model.Resource) string {
	return util.ResourceRestPath(resource)
}

func getTypescriptType(resource *model.Resource, property *model.ResourceProperty, nonPrimitive bool) string {
	switch property.Type {
	case model.ResourceProperty_STRING:
		return "string"
	case model.ResourceProperty_FLOAT64, model.ResourceProperty_FLOAT32, model.ResourceProperty_INT64, model.ResourceProperty_INT32:
		return "number"
	case model.ResourceProperty_BOOL:
		return "boolean"
	case model.ResourceProperty_REFERENCE:
		return typescriptClassName(property.Reference.Resource)
	case model.ResourceProperty_ENUM:
		return typescriptClassName(property.Name)
	case model.ResourceProperty_OBJECT:
		return "object"
	case model.ResourceProperty_LIST:
		return getTypescriptType(resource, property.Item, true) + "[]"
	case model.ResourceProperty_MAP:
		return "{ [key: string]: " + getTypescriptType(resource, property.Item, true) + " }"
	case model.ResourceProperty_TIME:
		return "string | Date"
	case model.ResourceProperty_DATE:
		return "string | Date"
	case model.ResourceProperty_TIMESTAMP:
		return "string | Date"
	case model.ResourceProperty_UUID:
		return "string"
	case model.ResourceProperty_BYTES:
		return "string | byte[]"
	case model.ResourceProperty_STRUCT:
		return typescriptClassName(*property.TypeRef)
	}

	panic("Unknown type: " + property.Type.String())
}

func typescriptClassName(resourceName string) string {
	return util.Capitalize(util.SnakeCaseToCamelCase(resourceName))
}

func typescriptVarName(resourceName string) string {
	return util.DeCapitalize(util.SnakeCaseToCamelCase(resourceName))
}

func hasInput(resource *model.Resource) bool {
	return len(resource.Types) > 0
}

func outputType(resource *model.Resource) string {
	if len(resource.Properties) > 0 {
		return getTypescriptType(resource, resource.Properties[0], false)
	} else {
		return "void"
	}
}

func getAllSubTypes(resource *model.Resource) []*model.ResourceSubType {
	var allTypes = resource.Types

	for _, resourceAction := range resourceActions {
		allTypes = append(allTypes, resourceAction.Types...)
	}

	for _, subType := range allTypes {
		subType.Name = typescriptClassName(subType.Name)
	}

	return allTypes
}

func getAllEnums(resource *model.Resource) []*model.ResourceProperty {
	var enums []*model.ResourceProperty
	var addedEnum = make(map[string]bool)
	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		var name = prop.Name
		if prop.Type == model.ResourceProperty_ENUM {
			var enumName = name

			if addedEnum[enumName] {
				return
			}

			enums = append(enums, &model.ResourceProperty{
				Name:       enumName,
				EnumValues: prop.EnumValues,
			})
			addedEnum[enumName] = true
		}
	})
	return enums
}

func getAllReferencedResources(resource *model.Resource) []*model.Reference {
	var references []*model.Reference
	var addedReferencedResources = make(map[string]bool)

	addedReferencedResources[resource.Namespace+"."+resource.Name] = true

	util.ResourceWalkProperties(resource, func(path string, prop *model.ResourceProperty) {
		if prop.Type == model.ResourceProperty_REFERENCE {

			if addedReferencedResources[prop.Reference.Namespace+"."+prop.Reference.Resource] {
				return
			}

			references = append(references, prop.Reference)
			addedReferencedResources[prop.Reference.Namespace+"."+prop.Reference.Resource] = true
		}
	})

	return references
}

func fileName(name string) string {
	return util.ToDashCase(name)
}

func enumName(enumValue string) string {
	return strings.ToUpper(util.ToSnakeCase(enumValue))
}
