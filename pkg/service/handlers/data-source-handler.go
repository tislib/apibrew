package handlers

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
)

type dataSourceHandler struct {
	backendProviderService service.BackendProviderService
}

func (h *dataSourceHandler) Register(eventHandler backend_event_handler.BackendEventHandler) {
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_UPDATE, h.AfterUpdate, resources.DataSourceResource))
	eventHandler.RegisterHandler(prepareStdHandler(101, model.Event_DELETE, h.AfterDelete, resources.DataSourceResource))
}

func (h *dataSourceHandler) AfterUpdate(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	for _, dataSource := range event.Records {
		err := h.backendProviderService.DestroyDataSource(ctx, dataSource.Id)

		if err != nil {
			return nil, err
		}
	}

	return event, nil
}

func (h *dataSourceHandler) AfterDelete(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
	for _, record := range event.Records {
		err := h.backendProviderService.DestroyDataSource(ctx, record.Id)

		if err != nil {
			return nil, err
		}
	}

	return event, nil
}
