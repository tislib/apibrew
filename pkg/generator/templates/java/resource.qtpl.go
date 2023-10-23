// Code generated by qtc from "resource.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/java/resource.qtpl:1
package java

//line templates/java/resource.qtpl:1
import "github.com/apibrew/apibrew/pkg/model"

//line templates/java/resource.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/java/resource.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/java/resource.qtpl:3
func StreamGenerateClassCode(qw422016 *qt422016.Writer, pkg string, resource *model.Resource) {
//line templates/java/resource.qtpl:3
	qw422016.N().S(`package `)
//line templates/java/resource.qtpl:4
	qw422016.E().S(pkg)
//line templates/java/resource.qtpl:4
	qw422016.N().S(`;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class `)
//line templates/java/resource.qtpl:16
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:16
	qw422016.N().S(` extends Entity {
`)
//line templates/java/resource.qtpl:17
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:17
		qw422016.N().S(`    `)
//line templates/java/resource.qtpl:18
		qw422016.N().S(getJavaPropertyAnnotations(resource, property))
//line templates/java/resource.qtpl:18
		qw422016.N().S(`
    private `)
//line templates/java/resource.qtpl:19
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:19
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:19
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:19
		qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:20
	}
//line templates/java/resource.qtpl:20
	qw422016.N().S(`
    public static final String NAMESPACE = "`)
//line templates/java/resource.qtpl:22
	qw422016.E().S(resource.Namespace)
//line templates/java/resource.qtpl:22
	qw422016.N().S(`";
    public static final String RESOURCE = "`)
//line templates/java/resource.qtpl:23
	qw422016.E().S(resource.Name)
//line templates/java/resource.qtpl:23
	qw422016.N().S(`";

    @JsonIgnore
    public static final EntityInfo<`)
//line templates/java/resource.qtpl:26
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:26
	qw422016.N().S(`> entityInfo = new EntityInfo<>("`)
//line templates/java/resource.qtpl:26
	qw422016.E().S(resource.Namespace)
//line templates/java/resource.qtpl:26
	qw422016.N().S(`", "`)
//line templates/java/resource.qtpl:26
	qw422016.E().S(resource.Name)
//line templates/java/resource.qtpl:26
	qw422016.N().S(`", `)
//line templates/java/resource.qtpl:26
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:26
	qw422016.N().S(`.class, "`)
//line templates/java/resource.qtpl:26
	qw422016.E().S(getRestPath(resource))
//line templates/java/resource.qtpl:26
	qw422016.N().S(`");

`)
//line templates/java/resource.qtpl:28
	for _, subType := range getAllSubTypes(resource) {
//line templates/java/resource.qtpl:28
		qw422016.N().S(`    public static class `)
//line templates/java/resource.qtpl:29
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:29
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:30
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:30
			qw422016.N().S(`        `)
//line templates/java/resource.qtpl:31
			qw422016.N().S(getJavaPropertyAnnotations(resource, property))
//line templates/java/resource.qtpl:31
			qw422016.N().S(`
        private `)
//line templates/java/resource.qtpl:32
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:32
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:32
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:32
			qw422016.N().S(`;
`)
//line templates/java/resource.qtpl:33
		}
//line templates/java/resource.qtpl:33
		qw422016.N().S(`
`)
//line templates/java/resource.qtpl:35
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:35
			qw422016.N().S(`        public `)
//line templates/java/resource.qtpl:36
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:36
			qw422016.N().S(` get`)
//line templates/java/resource.qtpl:36
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:36
			qw422016.N().S(`() {
            return `)
//line templates/java/resource.qtpl:37
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:37
			qw422016.N().S(`;
        }

        public void set`)
//line templates/java/resource.qtpl:40
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:40
			qw422016.N().S(`(`)
//line templates/java/resource.qtpl:40
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:40
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:40
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:40
			qw422016.N().S(`) {
            this.`)
//line templates/java/resource.qtpl:41
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:41
			qw422016.N().S(` = `)
//line templates/java/resource.qtpl:41
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:41
			qw422016.N().S(`;
        }

        public `)
//line templates/java/resource.qtpl:44
			qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:44
			qw422016.N().S(` with`)
//line templates/java/resource.qtpl:44
			qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:44
			qw422016.N().S(`(`)
//line templates/java/resource.qtpl:44
			qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:44
			qw422016.N().S(` `)
//line templates/java/resource.qtpl:44
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:44
			qw422016.N().S(`) {
            this.`)
//line templates/java/resource.qtpl:45
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:45
			qw422016.N().S(` = `)
//line templates/java/resource.qtpl:45
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:45
			qw422016.N().S(`;

            return this;
        }
`)
//line templates/java/resource.qtpl:49
		}
//line templates/java/resource.qtpl:49
		qw422016.N().S(`
        @Override
        public boolean equals(Object o) {
            if (!(o instanceof `)
//line templates/java/resource.qtpl:53
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:53
		qw422016.N().S(`)) {
                return false;
            }

            `)
//line templates/java/resource.qtpl:57
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:57
		qw422016.N().S(` obj = (`)
//line templates/java/resource.qtpl:57
		qw422016.E().S(javaClassName(subType.Name))
//line templates/java/resource.qtpl:57
		qw422016.N().S(`) o;

`)
//line templates/java/resource.qtpl:59
		for _, property := range subType.Properties {
//line templates/java/resource.qtpl:59
			qw422016.N().S(`            if (!Objects.equals(this.`)
//line templates/java/resource.qtpl:60
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:60
			qw422016.N().S(`, obj.`)
//line templates/java/resource.qtpl:60
			qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:60
			qw422016.N().S(`)) {
                return false;
            }
`)
//line templates/java/resource.qtpl:63
		}
//line templates/java/resource.qtpl:63
		qw422016.N().S(`
            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(`)
//line templates/java/resource.qtpl:70
		qw422016.N().S(hashcodePropertyNames(subType.Properties))
//line templates/java/resource.qtpl:70
		qw422016.N().S(`);
        }
    }
`)
//line templates/java/resource.qtpl:73
	}
//line templates/java/resource.qtpl:73
	qw422016.N().S(`
`)
//line templates/java/resource.qtpl:75
	for _, enum := range getAllEnums(resource) {
//line templates/java/resource.qtpl:75
		qw422016.N().S(`    public static enum `)
//line templates/java/resource.qtpl:76
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:76
		qw422016.N().S(` {
`)
//line templates/java/resource.qtpl:77
		for index, enumValue := range enum.EnumValues {
//line templates/java/resource.qtpl:77
			qw422016.N().S(`        `)
//line templates/java/resource.qtpl:78
			qw422016.N().S(enumName(enumValue))
//line templates/java/resource.qtpl:78
			qw422016.N().S(`("`)
//line templates/java/resource.qtpl:78
			qw422016.E().S(enumValue)
//line templates/java/resource.qtpl:78
			qw422016.N().S(`")`)
//line templates/java/resource.qtpl:78
			if index < len(enum.EnumValues)-1 {
//line templates/java/resource.qtpl:78
				qw422016.N().S(`,`)
//line templates/java/resource.qtpl:78
			} else {
//line templates/java/resource.qtpl:78
				qw422016.N().S(`;`)
//line templates/java/resource.qtpl:78
			}
//line templates/java/resource.qtpl:78
			qw422016.N().S(`
`)
//line templates/java/resource.qtpl:79
		}
//line templates/java/resource.qtpl:79
		qw422016.N().S(`
        private final String value;

        `)
//line templates/java/resource.qtpl:83
		qw422016.E().S(javaClassName(enum.Name))
//line templates/java/resource.qtpl:83
		qw422016.N().S(`(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
`)
//line templates/java/resource.qtpl:92
	}
//line templates/java/resource.qtpl:92
	qw422016.N().S(`
    `)
//line templates/java/resource.qtpl:94
	if len(resourceActions) > 0 {
//line templates/java/resource.qtpl:94
		qw422016.N().S(`
        public static class Service {

        private final Client client;

        public Service(Client client) {
            this.client = client;
        }

        `)
//line templates/java/resource.qtpl:103
		for _, resourceAction := range resourceActions {
//line templates/java/resource.qtpl:103
			qw422016.N().S(`
        `)
//line templates/java/resource.qtpl:104
			if hasInput(resourceAction) {
//line templates/java/resource.qtpl:104
				qw422016.N().S(`
        public `)
//line templates/java/resource.qtpl:105
				qw422016.N().S(outputType(resourceAction))
//line templates/java/resource.qtpl:105
				qw422016.N().S(` `)
//line templates/java/resource.qtpl:105
				qw422016.N().S(javaVarName(resourceAction.Name))
//line templates/java/resource.qtpl:105
				qw422016.N().S(` (`)
//line templates/java/resource.qtpl:105
				qw422016.N().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:105
				qw422016.N().S(` `)
//line templates/java/resource.qtpl:105
				qw422016.N().S(javaVarName(resource.Name))
//line templates/java/resource.qtpl:105
				qw422016.N().S(`, `)
//line templates/java/resource.qtpl:105
				qw422016.N().S(javaClassName(resourceAction.Name))
//line templates/java/resource.qtpl:105
				qw422016.N().S(`Input input) {
            `)
//line templates/java/resource.qtpl:106
				if len(resourceAction.Properties) > 0 {
//line templates/java/resource.qtpl:106
					qw422016.N().S(` return `)
//line templates/java/resource.qtpl:106
				}
//line templates/java/resource.qtpl:106
				qw422016.N().S(` client.executeRecordAction(`)
//line templates/java/resource.qtpl:106
				qw422016.N().S(outputType(resourceAction))
//line templates/java/resource.qtpl:106
				qw422016.N().S(`.class, `)
//line templates/java/resource.qtpl:106
				qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:106
				qw422016.N().S(`.NAMESPACE, `)
//line templates/java/resource.qtpl:106
				qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:106
				qw422016.N().S(`.RESOURCE, instance.getId().toString(), "`)
//line templates/java/resource.qtpl:106
				qw422016.E().S(resourceAction.Name)
//line templates/java/resource.qtpl:106
				qw422016.N().S(`", input);
        }
        `)
//line templates/java/resource.qtpl:108
			} else {
//line templates/java/resource.qtpl:108
				qw422016.N().S(`
        public `)
//line templates/java/resource.qtpl:109
				qw422016.N().S(outputType(resourceAction))
//line templates/java/resource.qtpl:109
				qw422016.N().S(` `)
//line templates/java/resource.qtpl:109
				qw422016.N().S(javaVarName(resourceAction.Name))
//line templates/java/resource.qtpl:109
				qw422016.N().S(` (`)
//line templates/java/resource.qtpl:109
				qw422016.N().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:109
				qw422016.N().S(` `)
//line templates/java/resource.qtpl:109
				qw422016.N().S(javaVarName(resource.Name))
//line templates/java/resource.qtpl:109
				qw422016.N().S(`) {
            `)
//line templates/java/resource.qtpl:110
				if len(resourceAction.Properties) > 0 {
//line templates/java/resource.qtpl:110
					qw422016.N().S(` return `)
//line templates/java/resource.qtpl:110
				}
//line templates/java/resource.qtpl:110
				qw422016.N().S(`  client.executeRecordAction(`)
//line templates/java/resource.qtpl:110
				qw422016.N().S(outputType(resourceAction))
//line templates/java/resource.qtpl:110
				qw422016.N().S(`.class, `)
//line templates/java/resource.qtpl:110
				qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:110
				qw422016.N().S(`.NAMESPACE, `)
//line templates/java/resource.qtpl:110
				qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:110
				qw422016.N().S(`.RESOURCE, instance.getId().toString(), "`)
//line templates/java/resource.qtpl:110
				qw422016.E().S(resourceAction.Name)
//line templates/java/resource.qtpl:110
				qw422016.N().S(`", null);
        }
        `)
//line templates/java/resource.qtpl:112
			}
//line templates/java/resource.qtpl:112
			qw422016.N().S(`
        `)
//line templates/java/resource.qtpl:113
		}
//line templates/java/resource.qtpl:113
		qw422016.N().S(`
        }
    `)
//line templates/java/resource.qtpl:115
	}
//line templates/java/resource.qtpl:115
	qw422016.N().S(`

    public `)
//line templates/java/resource.qtpl:117
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:117
	qw422016.N().S(`() {
    }

`)
//line templates/java/resource.qtpl:120
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:120
		qw422016.N().S(`    public `)
//line templates/java/resource.qtpl:121
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:121
		qw422016.N().S(` get`)
//line templates/java/resource.qtpl:121
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:121
		qw422016.N().S(`() {
        return `)
//line templates/java/resource.qtpl:122
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:122
		qw422016.N().S(`;
    }

    public void set`)
//line templates/java/resource.qtpl:125
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:125
		qw422016.N().S(`(`)
//line templates/java/resource.qtpl:125
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:125
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:125
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:125
		qw422016.N().S(`) {
        this.`)
//line templates/java/resource.qtpl:126
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:126
		qw422016.N().S(` = `)
//line templates/java/resource.qtpl:126
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:126
		qw422016.N().S(`;
    }

    public `)
//line templates/java/resource.qtpl:129
		qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:129
		qw422016.N().S(` with`)
//line templates/java/resource.qtpl:129
		qw422016.N().S(javaClassName(property.Name))
//line templates/java/resource.qtpl:129
		qw422016.N().S(`(`)
//line templates/java/resource.qtpl:129
		qw422016.N().S(getJavaType(resource, property, false))
//line templates/java/resource.qtpl:129
		qw422016.N().S(` `)
//line templates/java/resource.qtpl:129
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:129
		qw422016.N().S(`) {
        this.`)
//line templates/java/resource.qtpl:130
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:130
		qw422016.N().S(` = `)
//line templates/java/resource.qtpl:130
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:130
		qw422016.N().S(`;

        return this;
    }
`)
//line templates/java/resource.qtpl:134
	}
//line templates/java/resource.qtpl:134
	qw422016.N().S(`
    @Override
    public boolean equals(Object o) {
        if (!(o instanceof `)
//line templates/java/resource.qtpl:138
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:138
	qw422016.N().S(`)) {
            return false;
        }

        `)
//line templates/java/resource.qtpl:142
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:142
	qw422016.N().S(` obj = (`)
//line templates/java/resource.qtpl:142
	qw422016.E().S(javaClassName(resource.Name))
//line templates/java/resource.qtpl:142
	qw422016.N().S(`) o;

`)
//line templates/java/resource.qtpl:144
	for _, property := range resource.Properties {
//line templates/java/resource.qtpl:144
		qw422016.N().S(`        if (!Objects.equals(this.`)
//line templates/java/resource.qtpl:145
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:145
		qw422016.N().S(`, obj.`)
//line templates/java/resource.qtpl:145
		qw422016.N().S(propertyName(property))
//line templates/java/resource.qtpl:145
		qw422016.N().S(`)) {
            return false;
        }
`)
//line templates/java/resource.qtpl:148
	}
//line templates/java/resource.qtpl:148
	qw422016.N().S(`
        return true;
    }

    @Override
    public int hashCode() {
        if (id == null) {
            return super.hashCode();
        }

        return id.hashCode();
    }
}


`)
//line templates/java/resource.qtpl:164
}

//line templates/java/resource.qtpl:164
func WriteGenerateClassCode(qq422016 qtio422016.Writer, pkg string, resource *model.Resource) {
//line templates/java/resource.qtpl:164
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/java/resource.qtpl:164
	StreamGenerateClassCode(qw422016, pkg, resource)
//line templates/java/resource.qtpl:164
	qt422016.ReleaseWriter(qw422016)
//line templates/java/resource.qtpl:164
}

//line templates/java/resource.qtpl:164
func GenerateClassCode(pkg string, resource *model.Resource) string {
//line templates/java/resource.qtpl:164
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/java/resource.qtpl:164
	WriteGenerateClassCode(qb422016, pkg, resource)
//line templates/java/resource.qtpl:164
	qs422016 := string(qb422016.B)
//line templates/java/resource.qtpl:164
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/java/resource.qtpl:164
	return qs422016
//line templates/java/resource.qtpl:164
}
