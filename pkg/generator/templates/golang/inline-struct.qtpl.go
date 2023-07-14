// Code generated by qtc from "inline-struct.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/golang/inline-struct.qtpl:1
package golang

//line templates/golang/inline-struct.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/golang/inline-struct.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/golang/inline-struct.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/golang/inline-struct.qtpl:3
func StreamGenerateInlineStructCode(qw422016 *qt422016.Writer, resource *model.Resource, properties []*model.ResourceProperty) {
//line templates/golang/inline-struct.qtpl:3
	qw422016.N().S(`struct {
`)
//line templates/golang/inline-struct.qtpl:5
	for _, property := range properties {
//line templates/golang/inline-struct.qtpl:5
		qw422016.N().S(`	    `)
//line templates/golang/inline-struct.qtpl:6
		StreamGoName(qw422016, property.Name)
//line templates/golang/inline-struct.qtpl:6
		qw422016.N().S(` `)
//line templates/golang/inline-struct.qtpl:6
		StreamPropertyType(qw422016, resource, property)
//line templates/golang/inline-struct.qtpl:6
		qw422016.N().S(`
`)
//line templates/golang/inline-struct.qtpl:7
	}
//line templates/golang/inline-struct.qtpl:7
	qw422016.N().S(`}`)
//line templates/golang/inline-struct.qtpl:8
}

//line templates/golang/inline-struct.qtpl:8
func WriteGenerateInlineStructCode(qq422016 qtio422016.Writer, resource *model.Resource, properties []*model.ResourceProperty) {
//line templates/golang/inline-struct.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/golang/inline-struct.qtpl:8
	StreamGenerateInlineStructCode(qw422016, resource, properties)
//line templates/golang/inline-struct.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line templates/golang/inline-struct.qtpl:8
}

//line templates/golang/inline-struct.qtpl:8
func GenerateInlineStructCode(resource *model.Resource, properties []*model.ResourceProperty) string {
//line templates/golang/inline-struct.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/golang/inline-struct.qtpl:8
	WriteGenerateInlineStructCode(qb422016, resource, properties)
//line templates/golang/inline-struct.qtpl:8
	qs422016 := string(qb422016.B)
//line templates/golang/inline-struct.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/golang/inline-struct.qtpl:8
	return qs422016
//line templates/golang/inline-struct.qtpl:8
}