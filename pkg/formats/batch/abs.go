package batch

import (
	"github.com/tislib/apibrew/pkg/model"
)

type Writer interface {
	StartBatch(batch *model.BatchHeader) error
	EndBatch() error
	WriteResource(resource ...*model.Resource) error
	WriteRecord(namespace string, resourceName string, record ...*model.Record) error
}

type Reader interface {
}
