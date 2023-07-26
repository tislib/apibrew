package rest

import (
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type RecordApi interface {
	ConfigureRouter(r *mux.Router)
}

type recordApi struct {
	recordService   service.RecordService
	resourceService service.ResourceService
}

func (r *recordApi) ConfigureRouter(router *mux.Router) {
	subRoute := router.MatcherFunc(r.matchFunc).Subrouter()
	// collection level operations
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordList).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordCreate).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}", r.handleRecordApply).Methods("PATCH")

	// search
	subRoute.HandleFunc("/{resourceSlug}/_search", r.handleRecordSearch).Methods("POST")
	subRoute.HandleFunc("/{resourceSlug}/_resource", r.handleRecordResource).Methods("GET")

	// record level operations
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordGet).Methods("GET")
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordUpdate).Methods("PUT")
	subRoute.HandleFunc("/{resourceSlug}/{id}", r.handleRecordDelete).Methods("DELETE")
}

func (r *recordApi) matchFunc(request *http.Request, match *mux.RouteMatch) bool {
	pathParts := strings.Split(request.URL.Path, "/")

	if len(pathParts) < 2 {
		return false
	}

	slug := pathParts[1]

	return r.resourceService.GetSchema().ResourceBySlug[slug] != nil
}

func (r *recordApi) handleRecordList(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	// handle query parameters

	queryMap := make(map[string]interface{})

	for key := range request.URL.Query() {
		queryMap[key] = request.URL.Query().Get(key)
	}

	query, srvErr := r.recordService.PrepareQuery(resource, queryMap)

	if srvErr != nil {
		handleClientError(writer, srvErr)
		return
	}

	limit := 10
	offset := 0

	if request.URL.Query().Get("limit") != "" {
		var _err error
		limit, _err = strconv.Atoi(request.URL.Query().Get("limit"))

		if _err != nil {
			handleClientError(writer, _err)
			return
		}
	}

	if request.URL.Query().Get("offset") != "" {
		var _err error
		offset, _err = strconv.Atoi(request.URL.Query().Get("offset"))

		if _err != nil {
			handleClientError(writer, _err)
			return
		}
	}

	result, total, serviceErr := r.recordService.List(request.Context(), service.RecordListParams{
		Query:      query,
		Namespace:  resource.Namespace,
		Resource:   resource.Name,
		Limit:      uint32(limit),
		Offset:     uint64(offset),
		UseHistory: getRequestBoolFlag(request, "useHistory"),
	})

	ServiceResponder[*stub.ListRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&RecordList{
			Total:   uint64(total),
			Records: util.ArrayMap(result, NewRecordWrapper),
		}, serviceErr)
}

func (r *recordApi) handleRecordCreate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	record1 := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, record1)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.recordService.Create(request.Context(), service.RecordCreateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record1.toRecord()},
	})

	ServiceResponder[*stub.CreateRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&RecordList{
			Total:   uint64(len(res)),
			Records: util.ArrayMap(res, NewRecordWrapper),
		}, serviceErr)
}

func (r *recordApi) handleRecordApply(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	record1 := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, record1)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	if err != nil {
		handleClientError(writer, err)
		return
	}

	res, serviceErr := r.recordService.Apply(request.Context(), service.RecordUpdateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record1.toRecord()},
	})

	ServiceResponder[*stub.CreateRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&RecordList{
			Total:   uint64(len(res)),
			Records: util.ArrayMap(res, NewRecordWrapper),
		}, serviceErr)
}

func (r *recordApi) handleRecordGet(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	record, serviceErr := r.recordService.Get(request.Context(), service.RecordGetParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Id:        id,
	})

	ServiceResponder[*stub.GetRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(NewRecordWrapper(record), serviceErr)
}

func (r *recordApi) handleRecordUpdate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	recordWrap := NewEmptyRecordWrapper()

	err := parseRequestMessage(request, recordWrap)

	record := recordWrap.toRecord()

	if err != nil {
		handleClientError(writer, err)
		return
	}

	record.Id = id

	result, serviceErr := r.recordService.Update(request.Context(), service.RecordUpdateParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Records:   []*model.Record{record},
	})

	var updatedRecord *model.Record = nil

	if len(result) == 1 {
		updatedRecord = result[0]
	}

	ServiceResponder[*stub.UpdateRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(NewRecordWrapper(updatedRecord), serviceErr)
}

func (r *recordApi) handleRecordDelete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]
	id := vars["id"]

	serviceErr := r.recordService.Delete(request.Context(), service.RecordDeleteParams{
		Namespace: resource.Namespace,
		Resource:  resource.Name,
		Ids:       []string{id},
	})

	ServiceResponder[*stub.DeleteRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(nil, serviceErr)
}

func (r *recordApi) handleRecordSearch(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	listRecordRequest := new(SearchRecordRequest)

	err := parseRequestMessage(request, listRecordRequest)

	if err != nil {
		handleClientError(writer, err)
		return
	}

	result, total, serviceErr := r.recordService.List(request.Context(), service.RecordListParams{
		Query:      listRecordRequest.Query.expr,
		Namespace:  resource.Namespace,
		Resource:   resource.Name,
		Limit:      listRecordRequest.Limit,
		Offset:     listRecordRequest.Offset,
		UseHistory: listRecordRequest.UseHistory,
	})

	ServiceResponder[*stub.ListRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&RecordList{
			Total:   uint64(total),
			Records: util.ArrayMap(result, NewRecordWrapper),
		}, serviceErr)
}

func (r *recordApi) handleRecordResource(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resource := r.resourceService.GetSchema().ResourceBySlug[vars["resourceSlug"]]

	ServiceResponder[*stub.ListRecordRequest]().
		Writer(writer).
		Request(request).
		Respond(&ResourceWrapper{
			resource: resource,
		}, nil)
}

func NewRecordApi(container service.Container) RecordApi {
	return &recordApi{
		recordService:   container.GetRecordService(),
		resourceService: container.GetResourceService(),
	}
}