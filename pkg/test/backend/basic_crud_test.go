package backend

import (
	"fmt"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/test/setup"
	"github.com/tislib/data-handler/pkg/types"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

var dataSources = []*model.DataSource{
	setup.DhTest,
	//dhTestMysql,
}

var resources = make(map[*model.DataSource]*model.Resource)

func TestMain(t *testing.M) {
	setup.SetupDataSources(setup.Ctx, dataSources)

	var pendingResources []*model.Resource

	for _, dataSource := range dataSources {
		resources[dataSource] = setup.PrepareRichResource1()

		resources[dataSource].SourceConfig.DataSource = dataSource.Name
		resources[dataSource].Name = dataSource.Name + "-" + resources[dataSource].Name
		pendingResources = append(pendingResources, resources[dataSource])
	}

	setup.SetupResources(setup.Ctx, pendingResources)

	t.Run()

	setup.DestroyResources(setup.Ctx, pendingResources)
}

func TestCreateRecord(t *testing.T) {
	for _, dataSource := range dataSources {
		t.Run(fmt.Sprintf("%s[%s]", dataSource.Backend.String(), dataSource.Name), func(t *testing.T) {
			record1 := new(model.Record)
			st, err := structpb.NewStruct(map[string]interface{}{
				"bool":   true,
				"bytes":  "YXNk",
				"date":   "2006-01-02",
				"double": 12.3,
				"float":  31.200000762939453,
				"int32":  12,
				"int64":  34,
				"object": map[string]interface{}{ //@todo fixme double packing problem
					"test1": "test-123",
				},
				"string":    "asdasdksadjsakldksal",
				"text":      "test1233321",
				"time":      "17:04:05",
				"timestamp": "2006-01-02T15:04:05Z",
				"uuid":      "bdedf5b8-5179-11ed-bdc3-0242ac120002",
			})

			record1.Properties = st.GetFields()

			if err != nil {
				t.Error(err)
			}

			res, err := setup.GetTestDhClient().GetRecordClient().Create(setup.Ctx, &stub.CreateRecordRequest{
				Token:    "",
				Resource: setup.RichResource1.Name,
				Records:  []*model.Record{record1},
			})

			if err != nil {
				t.Error(err)
				return
			}

			getRes, err := setup.GetTestDhClient().GetRecordClient().Get(setup.Ctx, &stub.GetRecordRequest{
				Token:    "",
				Resource: setup.RichResource1.Name,
				Id:       res.Records[0].Id,
			})

			if err != nil {
				t.Error(err)
			}
			if len(record1.Properties) != len(getRes.Record.Properties) {
				t.Error("created and get records has different property count")
				return
			}

			for _, property := range setup.RichResource1.Properties {
				propertyType := types.ByResourcePropertyType(property.Type)
				val1, _ := propertyType.UnPack(record1.Properties[property.Name])
				val2, _ := propertyType.UnPack(getRes.Record.Properties[property.Name])

				if !propertyType.Equals(val1, val2) {
					t.Errorf("created and get records has different values: %v <=> %v", val1, val2)
					return
				}
			}
		})
	}
}
