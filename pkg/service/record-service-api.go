package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/backend"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	annotations2 "github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/service/params"
	"github.com/tislib/data-handler/pkg/system"
	"github.com/tislib/data-handler/pkg/util"
)

func (r *recordService) List(ctx context.Context, params params.RecordListParams) ([]*model.Record, uint32, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if err != nil {
		return nil, 0, err
	}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, 0, err
	}

	if err = r.genericHandler.BeforeList(ctx, resource, params); err != nil {
		return nil, 0, err
	}

	if handled, records, total, err := r.genericHandler.List(ctx, resource, params); handled {
		return records, total, err
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceId(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, 0, err
	}

	records, total, err := bck.ListRecords(ctx, backend.ListRecordParams{
		Resource:          resource,
		Query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
	})

	if err != nil {
		return nil, 0, err
	}

	if err := checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &records,
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, 0, err
	}

	if err = r.genericHandler.AfterList(ctx, resource, params, records, total); err != nil {
		return nil, 0, err
	}

	return records, total, err
}

func (r *recordService) Create(ctx context.Context, params params.RecordCreateParams) ([]*model.Record, []bool, errors.ServiceError) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range params.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record

	for _, item := range params.Records {
		item.DataType = model.DataType_USER
	}

	var insertedArray []bool
	var err errors.ServiceError

	var success = true
	var postOperationHandlers []func()

	defer func() {
		for _, handler := range postOperationHandlers {
			handler()
		}
	}()

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		resource, err = r.resourceService.GetResourceByName(ctx, params.Namespace, resourceName)

		if err != nil {
			success = false
			return nil, nil, err
		}

		if err = checkAccess(ctx, checkAccessParams{
			Resource:  resource,
			Records:   &list,
			Operation: model.OperationType_OPERATION_TYPE_CREATE,
		}); err != nil {
			return nil, nil, err
		}

		if isResourceRelatedResource(resource) {
			return nil, nil, errors.LogicalError.WithDetails("Resource and related resources cannot be modified from records API")
		}

		if err = r.validateRecords(resource, list, false); err != nil {
			return nil, nil, err
		}

		if err = r.genericHandler.BeforeCreate(ctx, resource, params); err != nil {
			return nil, nil, err
		}

		var records []*model.Record
		var inserted bool

		if handled, records, inserted, err := r.genericHandler.Create(ctx, resource, params); handled {
			return records, inserted, err
		}

		bck, err := r.backendServiceProvider.GetBackendByDataSourceId(ctx, resource.GetSourceConfig().DataSource)

		if err != nil {
			success = false
			return nil, []bool{}, err
		}

		tx, err := bck.BeginTransaction(ctx, false)

		if err != nil {
			success = false
			return nil, []bool{}, err
		}

		txCtx := context.WithValue(ctx, "transactionKey", tx)

		postOperationHandlers = append(postOperationHandlers, func() {
			if success {
				err = bck.CommitTransaction(txCtx)

				if err != nil {
					log.Print(err)
					success = false
				}
			} else {
				err = bck.RollbackTransaction(txCtx)

				if err != nil {
					log.Print(err)
				}
			}
		})

		records, inserted, err = bck.AddRecords(txCtx, backend.BulkRecordsParams{
			Resource:       resource,
			Records:        list,
			IgnoreIfExists: params.IgnoreIfExists,
		})

		insertedArray = append(insertedArray, inserted)

		if err != nil {
			success = false
			return nil, nil, err
		}

		if err = r.genericHandler.AfterCreate(ctx, resource, params, records); err != nil {
			success = false
			return nil, nil, err
		}

		result = append(result, records...)
	}

	return result, insertedArray, nil
}

func isResourceRelatedResource(resource *model.Resource) bool {
	return resource.Namespace == system.ResourceResource.Namespace && (resource.Name == system.ResourceResource.Name || resource.Name == system.ResourcePropertyResource.Name || resource.Name == system.ResourceReferenceResource.Name)
}

func (r *recordService) Update(ctx context.Context, params params.RecordUpdateParams) ([]*model.Record, errors.ServiceError) {
	var entityRecordMap = make(map[string][]*model.Record)

	for _, record := range params.Records {
		entityRecordMap[record.Resource] = append(entityRecordMap[record.Resource], record)
	}

	var result []*model.Record
	var err errors.ServiceError

	var success = true
	var postOperationHandlers []func()

	defer func() {
		for _, handler := range postOperationHandlers {
			handler()
		}
	}()

	for resourceName, list := range entityRecordMap {
		var resource *model.Resource
		if resource, err = r.resourceService.GetResourceByName(ctx, params.Namespace, resourceName); err != nil {
			return nil, err
		}

		if isResourceRelatedResource(resource) {
			return nil, errors.LogicalError.WithDetails("Resource and related resources cannot be modified from records API")
		}

		if err = checkAccess(ctx, checkAccessParams{
			Resource:  resource,
			Records:   &list,
			Operation: model.OperationType_OPERATION_TYPE_UPDATE,
		}); err != nil {
			return nil, err
		}

		if annotations2.IsEnabled(resource, annotations2.KeepHistory) && !params.CheckVersion {
			success = false
			return nil, errors.RecordValidationError.WithMessage("checkVersion must be enabled if resource has keepHistory enabled")
		}

		err = r.validateRecords(resource, list, true)

		if err != nil {
			success = false
			return nil, err
		}

		if err = r.genericHandler.BeforeUpdate(ctx, resource, params); err != nil {
			success = false
			return nil, err
		}

		if handled, records, err := r.genericHandler.Update(ctx, resource, params); handled {
			success = false
			return records, err
		}

		var records []*model.Record

		bck, err := r.backendServiceProvider.GetBackendByDataSourceId(ctx, resource.GetSourceConfig().DataSource)

		if err != nil {
			success = false
			return nil, err
		}

		tx, err := bck.BeginTransaction(ctx, false)

		if err != nil {
			success = false
			return nil, err
		}

		txCtx := context.WithValue(ctx, "transactionKey", tx)

		postOperationHandlers = append(postOperationHandlers, func() {
			if success {
				err = bck.CommitTransaction(txCtx)

				if err != nil {
					log.Print(err)
					success = false
				}
			} else {
				err = bck.RollbackTransaction(txCtx)

				if err != nil {
					log.Print(err)
				}
			}
		})

		records, err = bck.UpdateRecords(txCtx, backend.BulkRecordsParams{
			Resource:     resource,
			Records:      list,
			CheckVersion: params.CheckVersion,
		})

		if err != nil {
			success = false
			return nil, err
		}

		if err = r.genericHandler.AfterUpdate(ctx, resource, params, records); err != nil {
			return nil, err
		}

		result = append(result, records...)
	}

	return result, nil
}

func (r *recordService) GetRecord(ctx context.Context, namespace, resourceName, id string) (*model.Record, errors.ServiceError) {
	resource, err := r.resourceService.GetResourceByName(ctx, namespace, resourceName)

	if err != nil {
		return nil, err
	}

	if isResourceRelatedResource(resource) {
		return nil, errors.LogicalError.WithDetails("Resource and related resources cannot be modified from records API")
	}

	if err = checkAccess(ctx, checkAccessParams{
		Resource: resource,
		Records: &[]*model.Record{
			{
				Id: id,
			},
		},
		Operation: model.OperationType_OPERATION_TYPE_READ,
	}); err != nil {
		return nil, err
	}

	if err = r.genericHandler.BeforeGet(ctx, resource, id); err != nil {
		return nil, err
	}

	if handled, res, err := r.genericHandler.Get(ctx, resource, id); handled {
		return res, err
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceId(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return nil, err
	}

	res, err := bck.GetRecord(ctx, resource, id)

	if err != nil {
		return nil, err
	}

	if err = r.genericHandler.AfterGet(ctx, resource, id, res); err != nil {
		return nil, err
	}

	return res, err
}

func (r *recordService) FindBy(ctx context.Context, namespace, resourceName, propertyName string, value interface{}) (*model.Record, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin record-service FindBy")
	defer logger.Debug("Finish record-service FindBy")

	resource, err := r.resourceService.GetResourceByName(ctx, namespace, resourceName)

	if err != nil {
		return nil, err
	}

	queryMap := make(map[string]interface{})

	queryMap[propertyName] = value

	logger.Debug("Call PrepareQuery: ", queryMap)
	query, err := PrepareQuery(resource, queryMap)
	logger.Debug("Result record-service: ", query)

	if err != nil {
		return nil, err
	}

	res, total, err := r.List(ctx, params.RecordListParams{
		Query:             query,
		Namespace:         namespace,
		Resource:          resourceName,
		Limit:             2,
		Offset:            0,
		UseHistory:        false,
		ResolveReferences: false,
	})

	if total == 0 {
		return nil, errors.RecordNotFoundError
	}

	if total > 1 {
		return nil, errors.LogicalError.WithDetails("We have more than 1 record")
	}

	return res[0], nil
}

func (r *recordService) Get(ctx context.Context, params params.RecordGetParams) (*model.Record, errors.ServiceError) {
	return r.GetRecord(ctx, params.Namespace, params.Resource, params.Id)
}

func (r *recordService) Delete(ctx context.Context, params params.RecordDeleteParams) errors.ServiceError {
	resource, err := r.resourceService.GetResourceByName(ctx, params.Namespace, params.Resource)

	if err != nil {
		return err
	}

	var recordForCheck = util.ArrayMap(params.Ids, func(t string) *model.Record {
		return &model.Record{
			Id: t,
		}
	})

	if err = checkAccess(ctx, checkAccessParams{
		Resource:  resource,
		Records:   &recordForCheck,
		Operation: model.OperationType_OPERATION_TYPE_DELETE,
	}); err != nil {
		return err
	}

	if isResourceRelatedResource(resource) {
		return errors.LogicalError.WithDetails("Resource and related resources cannot be modified from records API")
	}

	if err = r.genericHandler.BeforeDelete(ctx, resource, params); err != nil {
		return err
	}

	if handled, err := r.genericHandler.Delete(ctx, resource, params); handled {
		return err
	}

	bck, err := r.backendServiceProvider.GetBackendByDataSourceId(ctx, resource.GetSourceConfig().DataSource)

	if err != nil {
		return err
	}

	if err = bck.DeleteRecords(ctx, resource, params.Ids); err != nil {
		return err
	}

	if err = r.genericHandler.AfterDelete(ctx, resource, params); err != nil {
		return err
	}

	return nil
}