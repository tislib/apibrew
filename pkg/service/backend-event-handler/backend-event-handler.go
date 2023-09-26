package backend_event_handler

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
)

type BackendEventHandler interface {
	RegisterHandler(handler Handler)
	UnRegisterHandler(handler Handler)
	HandleInternalOperation(ctx context.Context, incoming *model.Event, actualHandler HandlerFunc) (*model.Event, errors.ServiceError)
	PrepareInternalEvent(ctx context.Context, event *model.Event) *model.Event
}

type backendEventHandler struct {
	handlers                      []Handler
	authorizationService          service.AuthorizationService
	extensionEventSelectorMatcher *helper.ExtensionEventSelectorMatcher
}

func (b *backendEventHandler) PrepareInternalEvent(ctx context.Context, event *model.Event) *model.Event {
	event.Id = "internal-event-" + util.RandomHex(6)
	event.Time = timestamppb.Now()

	return event
}

func (b *backendEventHandler) HandleInternalOperation(ctx context.Context, originalEvent *model.Event, actualHandler HandlerFunc) (*model.Event, errors.ServiceError) {
	nextEvent := originalEvent

	var handlers []Handler

	if !annotations.IsEnabled(annotations.FromCtx(ctx), annotations.BypassExtensions) {
		handlers = b.filterHandlersForEvent(nextEvent)
	} else {
		err := b.authorizationService.CheckIsExtensionController(ctx)

		if err != nil {
			return nil, err
		}
	}

	if !nextEvent.Resource.Virtual {
		handlers = append(handlers, Handler{
			Id:        "actualHandler",
			Name:      "actualHandler",
			Fn:        actualHandler,
			Order:     NaturalOrder,
			Finalizes: false,
			Sync:      true,
			Responds:  true,
		})
	}

	sort.Sort(ByOrder(handlers))

	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Starting handler chain: %d", len(handlers))
	for _, handler := range handlers {
		nextEvent.Resource = originalEvent.Resource

		logger.Debugf("Calling handler: %s", handler.Name)
		logger.Tracef("Processing event: %s", nextEvent)
		if !handler.Sync {
			nextEvent.Sync = false
			go func(localHandler Handler) {
				_, err := localHandler.Fn(context.TODO(), nextEvent)

				if err != nil {
					logger.Error("Error from async handler", err)
				}
			}(handler)
		} else {
			nextEvent.Sync = true
			result, err := handler.Fn(ctx, nextEvent)
			logger.Debugf("Handler responded: %s", handler.Name)
			logger.Tracef("Handler responded: %s", nextEvent)

			if err != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, err)
				return nil, err
			}

			if result != nil && result.Error != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, result.Error)
				return nil, errors.FromProtoError(result.Error)
			}

			if handler.Responds {
				logger.Debugf("Handler [%s] responded with result", handler.Name)
				nextEvent = result
			}

			if nextEvent == nil {
				break
			}

			if nextEvent.Error != nil {
				logger.Warnf("Handler [%s] responded with error: %v", handler.Name, nextEvent.Error)
				return nil, errors.FromProtoError(nextEvent.Error)
			}

			if handler.Finalizes {
				logger.Debugf("Handler [%s] finalizes", handler.Name)
				break
			}

		}
	}
	logger.Debugf("Finished handler chain")
	logger.Tracef("Processed event: %s", nextEvent)

	return nextEvent, nil
}

func (b *backendEventHandler) RegisterHandler(handler Handler) {
	log.Debugf("Registering handler: %s [%v]", handler.Id, handler)

	b.handlers = append(b.handlers, handler)
}

func (b *backendEventHandler) UnRegisterHandler(handler Handler) {
	log.Debugf("Unregister handler: %s [%v]", handler.Id, handler)

	for i, h := range b.handlers {
		if h.Id == handler.Id {
			b.handlers = append(b.handlers[:i], b.handlers[i+1:]...)
			return
		}
	}

	log.Debugf("Unregister handler[not found]: %s [%v]", handler.Id, handler)
}

func (b *backendEventHandler) filterHandlersForEvent(incoming *model.Event) []Handler {
	var result []Handler

	for _, handler := range b.handlers {
		if b.extensionEventSelectorMatcher.SelectorMatches(incoming, handler.Selector) {
			log.Tracef("Handler matches: %s [%v]", handler.Id, handler)
			result = append(result, handler)
		}
	}

	return result
}

func NewBackendEventHandler(authorizationService service.AuthorizationService) BackendEventHandler {
	return &backendEventHandler{authorizationService: authorizationService, extensionEventSelectorMatcher: &helper.ExtensionEventSelectorMatcher{}}
}
