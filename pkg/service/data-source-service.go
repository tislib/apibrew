package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	log "github.com/sirupsen/logrus"
)

type dataSourceService struct {
	resourceService        abs.ResourceService
	recordService          abs.RecordService
	ServiceName            string
	backendProviderService abs.BackendProviderService
}

func (d *dataSourceService) ListEntities(ctx context.Context, id string) ([]*model.DataSourceCatalog, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("req", id).Debug("Begin data-source ListEntities")
	defer logger.Debug("End data-source ListEntities")

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return nil, err
	}

	return bck.ListEntities(ctx)
}

func (d *dataSourceService) GetStatus(ctx context.Context, id string) (connectionAlreadyInitiated bool, testConnection bool, err errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).Debug("Begin data-source GetStatus")
	defer logger.Debug("End data-source GetStatus")

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return
	}

	return bck.GetStatus(ctx)
}

func (d *dataSourceService) PrepareResourceFromEntity(ctx context.Context, id string, catalog, entity string) (*model.Resource, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("id", id).WithField("entity", entity).Debug("Begin data-source PrepareResourceFromEntity")
	defer logger.Debug("End data-source PrepareResourceFromEntity")

	dsRecord, err := d.recordService.Get(ctx, abs.RecordGetParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Id:        id,
	})

	if err != nil {
		return nil, err
	}

	bck, err := d.backendProviderService.GetBackendByDataSourceId(ctx, id)

	if err != nil {
		return nil, err
	}

	resource, err := bck.PrepareResourceFromEntity(ctx, catalog, entity)

	if err != nil {
		return nil, err
	}

	resource.SourceConfig = &model.ResourceSourceConfig{
		DataSource: dsRecord.Properties["name"].GetStringValue(),
		Catalog:    catalog,
		Entity:     entity,
	}

	return resource, nil
}

func (d *dataSourceService) Delete(ctx context.Context, ids []string) errors.ServiceError {
	logger := log.WithFields(logging.CtxFields(ctx))
	logger.WithField("ids", ids).Debug("Begin data-source Delete")
	defer logger.Debug("End data-source Delete")

	return d.recordService.Delete(ctx, abs.RecordDeleteParams{
		Namespace: resources.DataSourceResource.Namespace,
		Resource:  resources.DataSourceResource.Name,
		Ids:       ids,
	})
}

func (d *dataSourceService) Init(config *model.AppConfig) {

}

func NewDataSourceService(resourceService abs.ResourceService, recordService abs.RecordService, backendProviderService abs.BackendProviderService) abs.DataSourceService {
	return &dataSourceService{
		ServiceName:            "DataSourceService",
		resourceService:        resourceService,
		recordService:          recordService,
		backendProviderService: backendProviderService,
	}
}
