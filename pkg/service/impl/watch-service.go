package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type watchService struct {
	backendEventHandler  backend_event_handler.BackendEventHandler
	authorizationService service.AuthorizationService
	resourceService      service.ResourceService
}

func (w watchService) WatchResource(ctx context.Context, params service.WatchParams) (<-chan *model.Event, errors.ServiceError) {
	if params.Selector == nil {
		return nil, errors.RecordValidationError.WithMessage("selector is required")
	}

	if len(params.Selector.Resources) != 1 {
		return nil, errors.RecordValidationError.WithMessage("resource is not provided or more than one resource is provided")
	}

	if len(params.Selector.Namespaces) != 1 {
		return nil, errors.RecordValidationError.WithMessage("namespace is not provided or more than one namespace is provided")
	}

	var resource, err = w.resourceService.GetResourceByName(ctx, params.Selector.Namespaces[0], params.Selector.Resources[0])

	if err != nil {
		return nil, err
	}

	if err := w.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resource,
		Operation: resource_model.PermissionOperation_READ,
	}); err != nil {
		return nil, err
	}

	ch, err := w.watch(ctx, params)

	if err != nil {
		return nil, err
	}

	// check each event if the user has access to it

	result := make(chan *model.Event, params.BufferSize)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-ch:
				if event.Id != "heartbeat-message" {
					if err := w.authorizationService.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
						Resource:  resource,
						Operation: resource_model.PermissionOperation_READ,
						Records:   &event.Records,
					}); err != nil {
						log.Warnf("User don't have permission to access this event: %v", event.Id)
						continue
					}
				}

				result <- event
			}
		}
	}()

	return result, nil
}

func (w watchService) Watch(ctx context.Context, params service.WatchParams) (<-chan *model.Event, errors.ServiceError) {
	if err := w.authorizationService.CheckIsExtensionController(ctx); err != nil {
		return nil, err
	}

	return w.watch(ctx, params)
}

func (w watchService) watch(ctx context.Context, p service.WatchParams) (<-chan *model.Event, errors.ServiceError) {
	if p.BufferSize < 0 || p.BufferSize > 1000 {
		p.BufferSize = 100
	}

	out := make(chan *model.Event, p.BufferSize)
	result := make(chan *model.Event, p.BufferSize)
	watchHandlerId := util.RandomHex(6)
	watchHandler := backend_event_handler.Handler{
		Id:   "watch-handler-" + watchHandlerId,
		Name: "watch-handler-" + watchHandlerId,
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
			out <- event

			return event, nil
		},
		Selector: p.Selector,
		Order:    101,
	}

	// heartbeat
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-out:
				result <- event
			case <-time.After(3 * time.Second):
				log.Tracef("Heartbeat message sent to watcher: %v", watchHandlerId)
				result <- &model.Event{
					Id:   "heartbeat-message",
					Time: timestamppb.Now(),
				}
			}
		}
	}()

	go func() {
		<-ctx.Done()

		w.backendEventHandler.UnRegisterHandler(watchHandler)
		close(out)
	}()

	w.backendEventHandler.RegisterHandler(watchHandler)

	return result, nil
}

func NewWatchService(backendEventHandler backend_event_handler.BackendEventHandler, authorizationService service.AuthorizationService, resourceService service.ResourceService) service.WatchService {
	return &watchService{backendEventHandler: backendEventHandler, authorizationService: authorizationService, resourceService: resourceService}
}
