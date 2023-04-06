package yamlformat

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tislib/data-handler/pkg/client"
	"github.com/tislib/data-handler/pkg/formats"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type executor struct {
	params          ExecutorParams
	resources       []*model.Resource
	resourceNameMap map[string]*model.Resource
}

func (e *executor) Restore(ctx context.Context, file *os.File) error {

	var jsonUMo = protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: false,
		Resolver:       nil,
	}

	decoder := yaml.NewDecoder(file)

	for {
		var body map[string]interface{}
		var err = decoder.Decode(&body)

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		body = convert(body).(map[string]interface{})

		switch body["type"].(string) {
		case "resource":
			delete(body, "type")

			jsonData, err := json.Marshal(body)

			if err != nil {
				return err
			}

			var resource = new(model.Resource)
			err = jsonUMo.Unmarshal(jsonData, resource)

			if err != nil {
				return err
			}

			err = e.params.DhClient.ApplyResource(ctx, resource, e.params.DoMigration, e.params.ForceMigration)

			if err != nil {
				return err
			}
		case "record":
			var namespace string
			var resourceName string
			var ok bool

			if namespace, ok = body["namespace"].(string); !ok {
				namespace = "default"
			}

			if resourceName, ok = body["resource"].(string); !ok {
				return errors.New("resource field is required on record yaml definition")
			}

			var resource = e.resourceNameMap[namespace+"/"+resourceName]

			if resource == nil {
				return errors.New("Resource not found: " + namespace + "/" + resourceName)
			}

			delete(body, "type")
			delete(body, "resource")
			delete(body, "namespace")

			// locating resource

			jsonData, err := json.Marshal(body)

			if err != nil {
				return err
			}

			var record = new(model.Record)
			err = jsonUMo.Unmarshal(jsonData, record)

			if err != nil {
				return err
			}

			err = e.params.DhClient.ApplyRecord(ctx, resource, record)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *executor) init() error {
	resp, err := e.params.DhClient.GetResourceClient().List(context.TODO(), &stub.ListResourceRequest{})

	if err != nil {
		return err
	}

	e.resources = resp.Resources
	e.resourceNameMap = make(map[string]*model.Resource)

	for _, item := range e.resources {
		e.resourceNameMap[item.Namespace+"/"+item.Name] = item
	}

	return nil
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case map[string]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	Input          io.Reader
	DhClient       client.DhClient
	OverrideConfig OverrideConfig
	Token          string
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
}

func NewExecutor(params ExecutorParams) (formats.Executor, error) {
	exec := &executor{
		params: params,
	}

	return exec, exec.init()
}