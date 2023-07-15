// Code generated by qtc from "resource.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/golang/resource.qtpl:1
package golang

//line templates/golang/resource.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/golang/resource.qtpl:2
import "strings"

//line templates/golang/resource.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/golang/resource.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/golang/resource.qtpl:4
func StreamGenerateResourceCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource) {
//line templates/golang/resource.qtpl:5
	pkgParts := strings.Split(pkg, "/")

//line templates/golang/resource.qtpl:6
	pkgName := pkgParts[len(pkgParts)-1]

//line templates/golang/resource.qtpl:6
	qw422016.N().S(`package `)
//line templates/golang/resource.qtpl:7
	qw422016.E().S(pkgName)
//line templates/golang/resource.qtpl:7
	qw422016.N().S(`

`)
//line templates/golang/resource.qtpl:9
	for _, importLine := range getImports(resource) {
//line templates/golang/resource.qtpl:9
		qw422016.N().S(`import "`)
//line templates/golang/resource.qtpl:10
		qw422016.E().S(importLine)
//line templates/golang/resource.qtpl:10
		qw422016.N().S(`"
`)
//line templates/golang/resource.qtpl:11
	}
//line templates/golang/resource.qtpl:12
	for _, importLine := range getResourceSpecificImports(resource) {
//line templates/golang/resource.qtpl:12
		qw422016.N().S(`import "`)
//line templates/golang/resource.qtpl:13
		qw422016.E().S(importLine)
//line templates/golang/resource.qtpl:13
		qw422016.N().S(`"
`)
//line templates/golang/resource.qtpl:14
	}
//line templates/golang/resource.qtpl:14
	qw422016.N().S(`
`)
//line templates/golang/resource.qtpl:16
	StreamGenerateStructCode(qw422016, pkg, resource, resource.Name, resource.Properties, resource)
//line templates/golang/resource.qtpl:16
	qw422016.N().S(`
`)
//line templates/golang/resource.qtpl:17
	for _, subType := range getAllSubTypes(resource) {
//line templates/golang/resource.qtpl:18
		StreamGenerateStructCode(qw422016, pkg, resource, resource.Name+"_"+subType.Name, subType.Properties, subType)
//line templates/golang/resource.qtpl:18
		qw422016.N().S(`
`)
//line templates/golang/resource.qtpl:19
	}
//line templates/golang/resource.qtpl:20
	for _, enumProperty := range getAllEnums(resource) {
//line templates/golang/resource.qtpl:21
		StreamGenerateEnumCode(qw422016, enumProperty)
//line templates/golang/resource.qtpl:21
		qw422016.N().S(`
`)
//line templates/golang/resource.qtpl:22
	}
//line templates/golang/resource.qtpl:22
	qw422016.N().S(`
`)
//line templates/golang/resource.qtpl:24
}

//line templates/golang/resource.qtpl:24
func WriteGenerateResourceCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource) {
//line templates/golang/resource.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/resource.qtpl:24
	StreamGenerateResourceCode(qw422016, pkg, resource)
//line templates/golang/resource.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/resource.qtpl:24
}

//line templates/golang/resource.qtpl:24
func GenerateResourceCode(pkg string, resource *model.Resource) string {
//line templates/golang/resource.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/resource.qtpl:24
	WriteGenerateResourceCode(qb422016, pkg, resource)
//line templates/golang/resource.qtpl:24
	qs422016 := string(qb422016.B)
//line templates/golang/resource.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/resource.qtpl:24
	return qs422016
//line templates/golang/resource.qtpl:24
}
