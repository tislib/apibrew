package jsonformat

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/formats"
	"github.com/apibrew/apibrew/pkg/formats/unstructured/ops"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
	"io"
)

type writer struct {
	output             io.Writer
	hasMessageWritten  bool
	annotations        map[string]string
	unstructuredWriter ops.Writer
}

func (w *writer) WriteRecord(namespace string, resourceName string, records ...*model.Record) error {
	for _, record := range records {
		data, err := w.unstructuredWriter.WriteRecord(namespace, resourceName, record)

		if err != nil {
			return err
		}

		body, err := json.Marshal(data)

		if err != nil {
			return err
		}

		w.writePrefix()
		_, err = w.output.Write(body)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) WriteRecords(resource *model.Resource, total uint32, records []*model.Record) error {
	for _, record := range records {
		err := w.WriteRecord(resource.Namespace, resource.Name, record)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) IsBinary() bool {
	return false
}

func (w *writer) WriteResource(resources ...*model.Resource) error {
	for _, resource := range resources {
		data, err := w.unstructuredWriter.WriteResource(resource)

		if err != nil {
			return err
		}

		out, err := json.Marshal(data)

		if err != nil {
			return err
		}

		w.writePrefix()

		_, err = w.output.Write(out)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *writer) writePrefix() {
	if w.hasMessageWritten {
		if _, err := w.output.Write([]byte("---\n")); err != nil {
			log.Fatal(err)
		}
	}

	w.hasMessageWritten = true
}

func NewWriter(output io.Writer, annotations map[string]string) formats.Writer {
	return &writer{output: output, annotations: annotations, unstructuredWriter: ops.Writer{Annotations: annotations}}
}