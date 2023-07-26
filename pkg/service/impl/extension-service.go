package impl

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"time"
)

type extensionService struct {
	recordService          service.RecordService
	ServiceName            string
	backendProviderService service.BackendProviderService
	extensionVersionMap    map[string]uint32
	extensionHandlerMap    map[string]*backend_event_handler.Handler
	externalService        service.ExternalService
	backendEventHandler    backend_event_handler.BackendEventHandler
}

func (d *extensionService) Init(config *model.AppConfig) {
	d.runConfigureExtensions()

	go d.keepExtensionsRunning()
}

func (d *extensionService) keepExtensionsRunning() {
	for {
		time.Sleep(10 * time.Second)

		d.runConfigureExtensions()
	}
}

func (d *extensionService) runConfigureExtensions() {
	log.Trace("Start reconfiguring extension services")

	records, _, err := d.recordService.List(util.WithSystemContext(context.TODO()), service.RecordListParams{
		Namespace: resources.ExtensionResource.Namespace,
		Resource:  resources.ExtensionResource.Name,
	})

	if err != nil {
		panic(err)
	}

	var extensions = util.ArrayMap(records, resource_model.ExtensionMapperInstance.FromRecord)

	for _, ext := range extensions {
		log.Tracef("Configure extension: %v", ext)
		d.configureExtension(ext)
	}

	log.Trace("Finish reconfiguring extension services")
}

func (d *extensionService) RegisterExtension(extension *resource_model.Extension) {
	d.configureExtension(extension)
}

func (d *extensionService) UnRegisterExtension(extension *resource_model.Extension) {
	if d.extensionHandlerMap[extension.Id.String()] == nil {
		log.Warn("Trying to unregister extension that is not registered")
		return
	}

	d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id.String()])

	delete(d.extensionHandlerMap, extension.Id.String())
	delete(d.extensionVersionMap, extension.Id.String())
}

func (d *extensionService) configureExtension(extension *resource_model.Extension) {
	if d.extensionVersionMap[extension.Id.String()] == uint32(extension.Version) {
		return
	}

	d.extensionVersionMap[extension.Id.String()] = uint32(extension.Version)

	extensionHandler := d.prepareExtensionHandler(extension)
	if d.extensionHandlerMap[extension.Id.String()] != nil {
		d.backendEventHandler.UnRegisterHandler(*d.extensionHandlerMap[extension.Id.String()])
	}
	d.extensionHandlerMap[extension.Id.String()] = &extensionHandler

	d.backendEventHandler.RegisterHandler(*d.extensionHandlerMap[extension.Id.String()])
}

func (d *extensionService) prepareExtensionHandler(extension *resource_model.Extension) backend_event_handler.Handler {
	return backend_event_handler.Handler{
		Id:   "extension-" + extension.Id.String(),
		Name: "extension-" + extension.Name,
		//Selector:  extension.Selector,
		Order:     int(extension.Order),
		Finalizes: extension.Finalizes,
		Sync:      extension.Sync,
		Responds:  extension.Responds,
		Fn: func(ctx context.Context, event *model.Event) (*model.Event, errors.ServiceError) {
			return d.externalService.Call(ctx, extension.Call, event)
		},
	}
}

func NewExtensionService(recordService service.RecordService, backendProviderService service.BackendProviderService, backendEventHandler backend_event_handler.BackendEventHandler, externalService service.ExternalService) service.ExtensionService {
	return &extensionService{
		ServiceName:            "ExtensionService",
		extensionVersionMap:    make(map[string]uint32),
		extensionHandlerMap:    make(map[string]*backend_event_handler.Handler),
		recordService:          recordService,
		backendProviderService: backendProviderService,
		backendEventHandler:    backendEventHandler,
		externalService:        externalService,
	}
}