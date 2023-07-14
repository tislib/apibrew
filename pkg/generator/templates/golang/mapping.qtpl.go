// Code generated by qtc from "mapping.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/golang/mapping.qtpl:1
package golang

//line templates/golang/mapping.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/golang/mapping.qtpl:2
import "github.com/apibrew/apibrew/pkg/util"

//line templates/golang/mapping.qtpl:3
import "strings"

//line templates/golang/mapping.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/golang/mapping.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/golang/mapping.qtpl:5
func StreamPropertyTo(qw422016 *qt422016.Writer, property *model.ResourceProperty, isNullable bool, varName string) {
//line templates/golang/mapping.qtpl:6
	if isNullable {
//line templates/golang/mapping.qtpl:7
		varName = "*" + varName

//line templates/golang/mapping.qtpl:8
	}
//line templates/golang/mapping.qtpl:9
	if isPrimitive(property) {
//line templates/golang/mapping.qtpl:9
		qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:10
		StreamGoVarName(qw422016, property.Name)
//line templates/golang/mapping.qtpl:10
		qw422016.N().S(`, err := types.ByResourcePropertyType(model.ResourceProperty_`)
//line templates/golang/mapping.qtpl:10
		qw422016.E().S(property.Type.String())
//line templates/golang/mapping.qtpl:10
		qw422016.N().S(`).Pack(`)
//line templates/golang/mapping.qtpl:10
		qw422016.E().S(varName)
//line templates/golang/mapping.qtpl:10
		qw422016.N().S(`)
        if err != nil {
            panic(err)
        }
        properties["`)
//line templates/golang/mapping.qtpl:14
		qw422016.E().S(property.Name)
//line templates/golang/mapping.qtpl:14
		qw422016.N().S(`"] = `)
//line templates/golang/mapping.qtpl:14
		StreamGoVarName(qw422016, property.Name)
//line templates/golang/mapping.qtpl:14
		qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:15
	} else {
//line templates/golang/mapping.qtpl:16
	}
//line templates/golang/mapping.qtpl:17
}

//line templates/golang/mapping.qtpl:17
func WritePropertyTo(qq422016 qtio422016.Writer, property *model.ResourceProperty, isNullable bool, varName string) {
//line templates/golang/mapping.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/mapping.qtpl:17
	StreamPropertyTo(qw422016, property, isNullable, varName)
//line templates/golang/mapping.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/mapping.qtpl:17
}

//line templates/golang/mapping.qtpl:17
func PropertyTo(property *model.ResourceProperty, isNullable bool, varName string) string {
//line templates/golang/mapping.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/mapping.qtpl:17
	WritePropertyTo(qb422016, property, isNullable, varName)
//line templates/golang/mapping.qtpl:17
	qs422016 := string(qb422016.B)
//line templates/golang/mapping.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/mapping.qtpl:17
	return qs422016
//line templates/golang/mapping.qtpl:17
}

//line templates/golang/mapping.qtpl:22
func StreamGenerateResourceMappingCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:22
	qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:24
	StreamGenerateResourceMappingHeaderCode(qw422016, pkg, resource, resources)
//line templates/golang/mapping.qtpl:24
	qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:25
	StreamGenerateResourceMappingBodyCode(qw422016, pkg, resource, GoName(resource.Name), resource.Properties, resources)
//line templates/golang/mapping.qtpl:25
	qw422016.N().S(`

`)
//line templates/golang/mapping.qtpl:27
	for _, subType := range getAllSubTypes(resource) {
//line templates/golang/mapping.qtpl:28
		typeName := GoName(resource.Name + "_" + subType.Name)

//line templates/golang/mapping.qtpl:29
		StreamGenerateResourceMappingBodyCode(qw422016, pkg, resource, typeName, subType.Properties, resources)
//line templates/golang/mapping.qtpl:29
		qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:30
	}
//line templates/golang/mapping.qtpl:30
	qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:32
}

//line templates/golang/mapping.qtpl:32
func WriteGenerateResourceMappingCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/mapping.qtpl:32
	StreamGenerateResourceMappingCode(qw422016, pkg, resource, resources)
//line templates/golang/mapping.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/mapping.qtpl:32
}

//line templates/golang/mapping.qtpl:32
func GenerateResourceMappingCode(pkg string, resource *model.Resource, resources []*model.Resource) string {
//line templates/golang/mapping.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/mapping.qtpl:32
	WriteGenerateResourceMappingCode(qb422016, pkg, resource, resources)
//line templates/golang/mapping.qtpl:32
	qs422016 := string(qb422016.B)
//line templates/golang/mapping.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/mapping.qtpl:32
	return qs422016
//line templates/golang/mapping.qtpl:32
}

//line templates/golang/mapping.qtpl:35
func StreamGenerateResourceMappingHeaderCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:36
	pkgParts := strings.Split(pkg, "/")

//line templates/golang/mapping.qtpl:37
	pkgName := pkgParts[len(pkgParts)-1]

//line templates/golang/mapping.qtpl:37
	qw422016.N().S(`package `)
//line templates/golang/mapping.qtpl:38
	qw422016.E().S(pkgName)
//line templates/golang/mapping.qtpl:38
	qw422016.N().S(`

import (
    "github.com/apibrew/apibrew/pkg/model"
    "github.com/apibrew/apibrew/pkg/types"
    "google.golang.org/protobuf/types/known/structpb"
)

`)
//line templates/golang/mapping.qtpl:46
	for _, importLine := range getImports(resource.Properties) {
//line templates/golang/mapping.qtpl:46
		qw422016.N().S(`import "`)
//line templates/golang/mapping.qtpl:47
		qw422016.E().S(importLine)
//line templates/golang/mapping.qtpl:47
		qw422016.N().S(`"
`)
//line templates/golang/mapping.qtpl:48
	}
//line templates/golang/mapping.qtpl:49
}

//line templates/golang/mapping.qtpl:49
func WriteGenerateResourceMappingHeaderCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/mapping.qtpl:49
	StreamGenerateResourceMappingHeaderCode(qw422016, pkg, resource, resources)
//line templates/golang/mapping.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/mapping.qtpl:49
}

//line templates/golang/mapping.qtpl:49
func GenerateResourceMappingHeaderCode(pkg string, resource *model.Resource, resources []*model.Resource) string {
//line templates/golang/mapping.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/mapping.qtpl:49
	WriteGenerateResourceMappingHeaderCode(qb422016, pkg, resource, resources)
//line templates/golang/mapping.qtpl:49
	qs422016 := string(qb422016.B)
//line templates/golang/mapping.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/mapping.qtpl:49
	return qs422016
//line templates/golang/mapping.qtpl:49
}

//line templates/golang/mapping.qtpl:51
func StreamGenerateResourceMappingBodyCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource, typeName string, properties []*model.ResourceProperty, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:51
	qw422016.N().S(`type `)
//line templates/golang/mapping.qtpl:52
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:52
	qw422016.N().S(`Mapper struct {
}

func New`)
//line templates/golang/mapping.qtpl:55
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:55
	qw422016.N().S(`Mapper() *`)
//line templates/golang/mapping.qtpl:55
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:55
	qw422016.N().S(`Mapper {
    return &`)
//line templates/golang/mapping.qtpl:56
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:56
	qw422016.N().S(`Mapper{}
}

var `)
//line templates/golang/mapping.qtpl:59
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:59
	qw422016.N().S(`MapperInstance = New`)
//line templates/golang/mapping.qtpl:59
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:59
	qw422016.N().S(`Mapper()

func (m *`)
//line templates/golang/mapping.qtpl:61
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:61
	qw422016.N().S(`Mapper) New() *`)
//line templates/golang/mapping.qtpl:61
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:61
	qw422016.N().S(` {
    return &`)
//line templates/golang/mapping.qtpl:62
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:62
	qw422016.N().S(`{}
}

func (m *`)
//line templates/golang/mapping.qtpl:65
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:65
	qw422016.N().S(`Mapper) ToRecord(`)
//line templates/golang/mapping.qtpl:65
	StreamGoVarName(qw422016, typeName)
//line templates/golang/mapping.qtpl:65
	qw422016.N().S(` *`)
//line templates/golang/mapping.qtpl:65
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:65
	qw422016.N().S(`) *model.Record {
    var rec = &model.Record{}
    rec.Properties = m.ToProperties(`)
//line templates/golang/mapping.qtpl:67
	StreamGoVarName(qw422016, typeName)
//line templates/golang/mapping.qtpl:67
	qw422016.N().S(`)
    return rec
}

func (m *`)
//line templates/golang/mapping.qtpl:71
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:71
	qw422016.N().S(`Mapper) FromRecord(record *model.Record) *`)
//line templates/golang/mapping.qtpl:71
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:71
	qw422016.N().S(`  {
    return m.FromProperties(record.Properties)
}

func (m *`)
//line templates/golang/mapping.qtpl:75
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:75
	qw422016.N().S(`Mapper) ToProperties(`)
//line templates/golang/mapping.qtpl:75
	StreamGoVarName(qw422016, typeName)
//line templates/golang/mapping.qtpl:75
	qw422016.N().S(` *`)
//line templates/golang/mapping.qtpl:75
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:75
	qw422016.N().S(`) map[string]*structpb.Value {
    var properties = make(map[string]*structpb.Value)

`)
//line templates/golang/mapping.qtpl:78
	for _, property := range properties {
//line templates/golang/mapping.qtpl:79
		varName := GoVarName(typeName) + "." + GoName(property.Name)

//line templates/golang/mapping.qtpl:80
		if isNullable(property) {
//line templates/golang/mapping.qtpl:80
			qw422016.N().S(`    if `)
//line templates/golang/mapping.qtpl:81
			qw422016.E().S(varName)
//line templates/golang/mapping.qtpl:81
			qw422016.N().S(` != nil {
`)
//line templates/golang/mapping.qtpl:82
			StreamPropertyTo(qw422016, property, true, varName)
//line templates/golang/mapping.qtpl:82
			qw422016.N().S(`    }
`)
//line templates/golang/mapping.qtpl:84
		} else {
//line templates/golang/mapping.qtpl:85
			StreamPropertyTo(qw422016, property, false, varName)
//line templates/golang/mapping.qtpl:86
		}
//line templates/golang/mapping.qtpl:86
		qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:88
	}
//line templates/golang/mapping.qtpl:88
	qw422016.N().S(`    return properties
}

func (m *`)
//line templates/golang/mapping.qtpl:92
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:92
	qw422016.N().S(`Mapper) FromProperties(properties map[string]*structpb.Value) *`)
//line templates/golang/mapping.qtpl:92
	qw422016.E().S(typeName)
//line templates/golang/mapping.qtpl:92
	qw422016.N().S(`  {
    var s = m.New()
`)
//line templates/golang/mapping.qtpl:94
	for _, property := range properties {
//line templates/golang/mapping.qtpl:94
		qw422016.N().S(`    if properties["`)
//line templates/golang/mapping.qtpl:95
		qw422016.E().S(property.Name)
//line templates/golang/mapping.qtpl:95
		qw422016.N().S(`"] != nil {
        `)
//line templates/golang/mapping.qtpl:96
		valueVarName := "var_" + util.RandomHex(6)

//line templates/golang/mapping.qtpl:96
		qw422016.N().S(`
        `)
//line templates/golang/mapping.qtpl:97
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:97
		qw422016.N().S(` := properties["`)
//line templates/golang/mapping.qtpl:97
		qw422016.E().S(property.Name)
//line templates/golang/mapping.qtpl:97
		qw422016.N().S(`"]
        `)
//line templates/golang/mapping.qtpl:98
		StreamPreparePropertyMapping(qw422016, resource, property, valueVarName, false)
//line templates/golang/mapping.qtpl:98
		qw422016.N().S(`
        s.`)
//line templates/golang/mapping.qtpl:99
		StreamGoName(qw422016, property.Name)
//line templates/golang/mapping.qtpl:99
		qw422016.N().S(` = `)
//line templates/golang/mapping.qtpl:99
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:99
		qw422016.N().S(`_mapped
    }
`)
//line templates/golang/mapping.qtpl:101
	}
//line templates/golang/mapping.qtpl:101
	qw422016.N().S(`    return s
}

`)
//line templates/golang/mapping.qtpl:105
}

//line templates/golang/mapping.qtpl:105
func WriteGenerateResourceMappingBodyCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource, typeName string, properties []*model.ResourceProperty, resources []*model.Resource) {
//line templates/golang/mapping.qtpl:105
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/mapping.qtpl:105
	StreamGenerateResourceMappingBodyCode(qw422016, pkg, resource, typeName, properties, resources)
//line templates/golang/mapping.qtpl:105
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/mapping.qtpl:105
}

//line templates/golang/mapping.qtpl:105
func GenerateResourceMappingBodyCode(pkg string, resource *model.Resource, typeName string, properties []*model.ResourceProperty, resources []*model.Resource) string {
//line templates/golang/mapping.qtpl:105
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/mapping.qtpl:105
	WriteGenerateResourceMappingBodyCode(qb422016, pkg, resource, typeName, properties, resources)
//line templates/golang/mapping.qtpl:105
	qs422016 := string(qb422016.B)
//line templates/golang/mapping.qtpl:105
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/mapping.qtpl:105
	return qs422016
//line templates/golang/mapping.qtpl:105
}

//line templates/golang/mapping.qtpl:107
func StreamPreparePropertyMapping(qw422016 *qt422016.Writer, resource *model.Resource, property *model.ResourceProperty, valueVarName string, skipOptionalFlag bool) {
//line templates/golang/mapping.qtpl:108
	if property.Type == model.ResourceProperty_REFERENCE {
//line templates/golang/mapping.qtpl:108
		qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:109
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:109
		qw422016.N().S(`_mapped := `)
//line templates/golang/mapping.qtpl:109
		StreamGoName(qw422016, property.Reference.Resource)
//line templates/golang/mapping.qtpl:109
		qw422016.N().S(`MapperInstance.FromProperties(`)
//line templates/golang/mapping.qtpl:109
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:109
		qw422016.N().S(`.GetStructValue().Fields)
`)
//line templates/golang/mapping.qtpl:110
	} else if property.Type == model.ResourceProperty_MAP {
//line templates/golang/mapping.qtpl:110
		qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:111
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:111
		qw422016.N().S(`_mapped := make(`)
//line templates/golang/mapping.qtpl:111
		StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:111
		qw422016.N().S(`)
        for k, v := range `)
//line templates/golang/mapping.qtpl:112
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:112
		qw422016.N().S(`.GetStructValue().Fields {
            `)
//line templates/golang/mapping.qtpl:113
		subValueVarName := "var_" + util.RandomHex(6)

//line templates/golang/mapping.qtpl:113
		qw422016.N().S(`
            `)
//line templates/golang/mapping.qtpl:114
		qw422016.E().S(subValueVarName)
//line templates/golang/mapping.qtpl:114
		qw422016.N().S(` := v
            `)
//line templates/golang/mapping.qtpl:115
		StreamPreparePropertyMapping(qw422016, resource, property.Item, subValueVarName, true)
//line templates/golang/mapping.qtpl:115
		qw422016.N().S(`
            `)
//line templates/golang/mapping.qtpl:116
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:116
		qw422016.N().S(`_mapped[k] = `)
//line templates/golang/mapping.qtpl:116
		qw422016.E().S(subValueVarName)
//line templates/golang/mapping.qtpl:116
		qw422016.N().S(`_mapped
        }
`)
//line templates/golang/mapping.qtpl:118
	} else if property.Type == model.ResourceProperty_LIST {
//line templates/golang/mapping.qtpl:118
		qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:119
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:119
		qw422016.N().S(`_mapped := []`)
//line templates/golang/mapping.qtpl:119
		StreamPropPureGoType(qw422016, resource, property.Item, property.Name)
//line templates/golang/mapping.qtpl:119
		qw422016.N().S(`{}
        for _, v := range `)
//line templates/golang/mapping.qtpl:120
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:120
		qw422016.N().S(`.GetListValue().Values {
            `)
//line templates/golang/mapping.qtpl:121
		subValueVarName := "var_" + util.RandomHex(6)

//line templates/golang/mapping.qtpl:121
		qw422016.N().S(`
            `)
//line templates/golang/mapping.qtpl:122
		qw422016.E().S(subValueVarName)
//line templates/golang/mapping.qtpl:122
		qw422016.N().S(` := v
            `)
//line templates/golang/mapping.qtpl:123
		StreamPreparePropertyMapping(qw422016, resource, property.Item, subValueVarName, true)
//line templates/golang/mapping.qtpl:123
		qw422016.N().S(`
            `)
//line templates/golang/mapping.qtpl:124
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:124
		qw422016.N().S(`_mapped = append(`)
//line templates/golang/mapping.qtpl:124
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:124
		qw422016.N().S(`_mapped, `)
//line templates/golang/mapping.qtpl:124
		qw422016.E().S(subValueVarName)
//line templates/golang/mapping.qtpl:124
		qw422016.N().S(`_mapped)
        }
`)
//line templates/golang/mapping.qtpl:126
	} else if property.Type == model.ResourceProperty_ENUM {
//line templates/golang/mapping.qtpl:127
		if isNullable(property) && !skipOptionalFlag {
//line templates/golang/mapping.qtpl:127
			qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:128
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:128
			qw422016.N().S(`_mapped := new(`)
//line templates/golang/mapping.qtpl:128
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:128
			qw422016.N().S(`)
        *`)
//line templates/golang/mapping.qtpl:129
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:129
			qw422016.N().S(`_mapped = (`)
//line templates/golang/mapping.qtpl:129
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:129
			qw422016.N().S(`)(`)
//line templates/golang/mapping.qtpl:129
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:129
			qw422016.N().S(`.GetStringValue())
`)
//line templates/golang/mapping.qtpl:130
		} else {
//line templates/golang/mapping.qtpl:130
			qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:131
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:131
			qw422016.N().S(`_mapped := (`)
//line templates/golang/mapping.qtpl:131
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:131
			qw422016.N().S(`)(`)
//line templates/golang/mapping.qtpl:131
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:131
			qw422016.N().S(`.GetStringValue())
`)
//line templates/golang/mapping.qtpl:132
		}
//line templates/golang/mapping.qtpl:133
	} else if property.Type == model.ResourceProperty_STRUCT {
//line templates/golang/mapping.qtpl:133
		qw422016.N().S(`        var mappedValue = `)
//line templates/golang/mapping.qtpl:134
		StreamGoName(qw422016, resource.Name+"_"+*property.TypeRef)
//line templates/golang/mapping.qtpl:134
		qw422016.N().S(`MapperInstance.FromProperties(`)
//line templates/golang/mapping.qtpl:134
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:134
		qw422016.N().S(`.GetStructValue().Fields)
        `)
//line templates/golang/mapping.qtpl:135
		if property.Required {
//line templates/golang/mapping.qtpl:135
			qw422016.N().S(`
        `)
//line templates/golang/mapping.qtpl:136
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:136
			qw422016.N().S(`_mapped := *mappedValue
        `)
//line templates/golang/mapping.qtpl:137
		} else {
//line templates/golang/mapping.qtpl:137
			qw422016.N().S(`
        `)
//line templates/golang/mapping.qtpl:138
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:138
			qw422016.N().S(`_mapped := mappedValue
        `)
//line templates/golang/mapping.qtpl:139
		}
//line templates/golang/mapping.qtpl:139
		qw422016.N().S(`
`)
//line templates/golang/mapping.qtpl:140
	} else {
//line templates/golang/mapping.qtpl:140
		qw422016.N().S(`        val, err := types.ByResourcePropertyType(model.ResourceProperty_`)
//line templates/golang/mapping.qtpl:141
		qw422016.E().S(property.Type.String())
//line templates/golang/mapping.qtpl:141
		qw422016.N().S(`).UnPack(`)
//line templates/golang/mapping.qtpl:141
		qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:141
		qw422016.N().S(`)

        if err != nil {
            panic(err)
        }

`)
//line templates/golang/mapping.qtpl:147
		if isNullable(property) && !skipOptionalFlag {
//line templates/golang/mapping.qtpl:147
			qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:148
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:148
			qw422016.N().S(`_mapped := new(`)
//line templates/golang/mapping.qtpl:148
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:148
			qw422016.N().S(`)
        *`)
//line templates/golang/mapping.qtpl:149
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:149
			qw422016.N().S(`_mapped = val.(`)
//line templates/golang/mapping.qtpl:149
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:149
			qw422016.N().S(`)
`)
//line templates/golang/mapping.qtpl:150
		} else {
//line templates/golang/mapping.qtpl:150
			qw422016.N().S(`        `)
//line templates/golang/mapping.qtpl:151
			qw422016.E().S(valueVarName)
//line templates/golang/mapping.qtpl:151
			qw422016.N().S(`_mapped := val.(`)
//line templates/golang/mapping.qtpl:151
			StreamPropPureGoType(qw422016, resource, property, property.Name)
//line templates/golang/mapping.qtpl:151
			qw422016.N().S(`)
`)
//line templates/golang/mapping.qtpl:152
		}
//line templates/golang/mapping.qtpl:153
	}
//line templates/golang/mapping.qtpl:154
}

//line templates/golang/mapping.qtpl:154
func WritePreparePropertyMapping(qq422016 qtio422016.Writer, resource *model.Resource, property *model.ResourceProperty, valueVarName string, skipOptionalFlag bool) {
//line templates/golang/mapping.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/mapping.qtpl:154
	StreamPreparePropertyMapping(qw422016, resource, property, valueVarName, skipOptionalFlag)
//line templates/golang/mapping.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/mapping.qtpl:154
}

//line templates/golang/mapping.qtpl:154
func PreparePropertyMapping(resource *model.Resource, property *model.ResourceProperty, valueVarName string, skipOptionalFlag bool) string {
//line templates/golang/mapping.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/mapping.qtpl:154
	WritePreparePropertyMapping(qb422016, resource, property, valueVarName, skipOptionalFlag)
//line templates/golang/mapping.qtpl:154
	qs422016 := string(qb422016.B)
//line templates/golang/mapping.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/mapping.qtpl:154
	return qs422016
//line templates/golang/mapping.qtpl:154
}