package postgres

import (
	"data-handler/model"
	"data-handler/service/backend"
)

type postgresDataSourceBackend struct {
	Options      *model.PostgresqlOptions
	dataSourceId string
}

func (p postgresDataSourceBackend) GetDataSourceId() string {
	return p.dataSourceId
}

func NewPostgresDataSourceBackend(dataSourceId string, options *model.PostgresqlOptions) backend.DataSourceBackend {
	return &postgresDataSourceBackend{
		dataSourceId: dataSourceId,
		Options:      options,
	}
}