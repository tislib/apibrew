package test

import (
	"data-handler/stub/model"
	"testing"
)

func TestCreateResource(t *testing.T) {
	withDataSource(t, container, dataSource1, func(createdDataSource *model.DataSource) {
		// testing is done

	})
}
