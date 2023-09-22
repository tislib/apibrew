// AUTOGENERATED FILE

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type Extension struct {
	Id          *uuid.UUID              `json:"id,omitempty"`
	Version     int32                   `json:"version,omitempty"`
	CreatedBy   *string                 `json:"createdBy,omitempty"`
	UpdatedBy   *string                 `json:"updatedBy,omitempty"`
	CreatedOn   *time.Time              `json:"createdOn,omitempty"`
	UpdatedOn   *time.Time              `json:"updatedOn,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Selector    *ExtensionEventSelector `json:"selector,omitempty"`
	Order       int32                   `json:"order,omitempty"`
	Finalizes   bool                    `json:"finalizes,omitempty"`
	Sync        bool                    `json:"sync,omitempty"`
	Responds    bool                    `json:"responds,omitempty"`
	Call        ExtensionExternalCall   `json:"call,omitempty"`
	Annotations map[string]string       `json:"annotations,omitempty"`
}

func (s *Extension) GetId() *uuid.UUID {
	return s.Id
}
func (s *Extension) GetVersion() int32 {
	return s.Version
}
func (s *Extension) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *Extension) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *Extension) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *Extension) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}
func (s *Extension) GetName() string {
	return s.Name
}
func (s *Extension) GetDescription() *string {
	return s.Description
}
func (s *Extension) GetSelector() *ExtensionEventSelector {
	return s.Selector
}
func (s *Extension) GetOrder() int32 {
	return s.Order
}
func (s *Extension) GetFinalizes() bool {
	return s.Finalizes
}
func (s *Extension) GetSync() bool {
	return s.Sync
}
func (s *Extension) GetResponds() bool {
	return s.Responds
}
func (s *Extension) GetCall() ExtensionExternalCall {
	return s.Call
}
func (s *Extension) GetAnnotations() map[string]string {
	return s.Annotations
}

type ExtensionBooleanExpression struct {
	And                []ExtensionBooleanExpression   `json:"and,omitempty"`
	Or                 []ExtensionBooleanExpression   `json:"or,omitempty"`
	Not                *ExtensionBooleanExpression    `json:"not,omitempty"`
	Equal              *ExtensionPairExpression       `json:"equal,omitempty"`
	LessThan           *ExtensionPairExpression       `json:"lessThan,omitempty"`
	GreaterThan        *ExtensionPairExpression       `json:"greaterThan,omitempty"`
	LessThanOrEqual    *ExtensionPairExpression       `json:"lessThanOrEqual,omitempty"`
	GreaterThanOrEqual *ExtensionPairExpression       `json:"greaterThanOrEqual,omitempty"`
	In                 *ExtensionPairExpression       `json:"in,omitempty"`
	IsNull             *ExtensionExpression           `json:"isNull,omitempty"`
	RegexMatch         *ExtensionRegexMatchExpression `json:"regexMatch,omitempty"`
}

func (s *ExtensionBooleanExpression) GetAnd() []ExtensionBooleanExpression {
	return s.And
}
func (s *ExtensionBooleanExpression) GetOr() []ExtensionBooleanExpression {
	return s.Or
}
func (s *ExtensionBooleanExpression) GetNot() *ExtensionBooleanExpression {
	return s.Not
}
func (s *ExtensionBooleanExpression) GetEqual() *ExtensionPairExpression {
	return s.Equal
}
func (s *ExtensionBooleanExpression) GetLessThan() *ExtensionPairExpression {
	return s.LessThan
}
func (s *ExtensionBooleanExpression) GetGreaterThan() *ExtensionPairExpression {
	return s.GreaterThan
}
func (s *ExtensionBooleanExpression) GetLessThanOrEqual() *ExtensionPairExpression {
	return s.LessThanOrEqual
}
func (s *ExtensionBooleanExpression) GetGreaterThanOrEqual() *ExtensionPairExpression {
	return s.GreaterThanOrEqual
}
func (s *ExtensionBooleanExpression) GetIn() *ExtensionPairExpression {
	return s.In
}
func (s *ExtensionBooleanExpression) GetIsNull() *ExtensionExpression {
	return s.IsNull
}
func (s *ExtensionBooleanExpression) GetRegexMatch() *ExtensionRegexMatchExpression {
	return s.RegexMatch
}

type ExtensionPairExpression struct {
	Left  *ExtensionExpression `json:"left,omitempty"`
	Right *ExtensionExpression `json:"right,omitempty"`
}

func (s *ExtensionPairExpression) GetLeft() *ExtensionExpression {
	return s.Left
}
func (s *ExtensionPairExpression) GetRight() *ExtensionExpression {
	return s.Right
}

type ExtensionRefValue struct {
	Namespace  *string                `json:"namespace,omitempty"`
	Resource   *string                `json:"resource,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

func (s *ExtensionRefValue) GetNamespace() *string {
	return s.Namespace
}
func (s *ExtensionRefValue) GetResource() *string {
	return s.Resource
}
func (s *ExtensionRefValue) GetProperties() map[string]interface{} {
	return s.Properties
}

type ExtensionRegexMatchExpression struct {
	Pattern    *string              `json:"pattern,omitempty"`
	Expression *ExtensionExpression `json:"expression,omitempty"`
}

func (s *ExtensionRegexMatchExpression) GetPattern() *string {
	return s.Pattern
}
func (s *ExtensionRegexMatchExpression) GetExpression() *ExtensionExpression {
	return s.Expression
}

type ExtensionExpression struct {
	Property *string            `json:"property,omitempty"`
	Value    interface{}        `json:"value,omitempty"`
	RefValue *ExtensionRefValue `json:"refValue,omitempty"`
}

func (s *ExtensionExpression) GetProperty() *string {
	return s.Property
}
func (s *ExtensionExpression) GetValue() interface{} {
	return s.Value
}
func (s *ExtensionExpression) GetRefValue() *ExtensionRefValue {
	return s.RefValue
}

type ExtensionFunctionCall struct {
	Host         string `json:"host,omitempty"`
	FunctionName string `json:"functionName,omitempty"`
}

func (s *ExtensionFunctionCall) GetHost() string {
	return s.Host
}
func (s *ExtensionFunctionCall) GetFunctionName() string {
	return s.FunctionName
}

type ExtensionHttpCall struct {
	Uri    string `json:"uri,omitempty"`
	Method string `json:"method,omitempty"`
}

func (s *ExtensionHttpCall) GetUri() string {
	return s.Uri
}
func (s *ExtensionHttpCall) GetMethod() string {
	return s.Method
}

type ExtensionChannelCall struct {
	ChannelKey string `json:"channelKey,omitempty"`
}

func (s *ExtensionChannelCall) GetChannelKey() string {
	return s.ChannelKey
}

type ExtensionExternalCall struct {
	FunctionCall *ExtensionFunctionCall `json:"functionCall,omitempty"`
	HttpCall     *ExtensionHttpCall     `json:"httpCall,omitempty"`
	ChannelCall  *ExtensionChannelCall  `json:"channelCall,omitempty"`
}

func (s *ExtensionExternalCall) GetFunctionCall() *ExtensionFunctionCall {
	return s.FunctionCall
}
func (s *ExtensionExternalCall) GetHttpCall() *ExtensionHttpCall {
	return s.HttpCall
}
func (s *ExtensionExternalCall) GetChannelCall() *ExtensionChannelCall {
	return s.ChannelCall
}

type ExtensionEventSelector struct {
	Actions        []EventAction               `json:"actions,omitempty"`
	RecordSelector *ExtensionBooleanExpression `json:"recordSelector,omitempty"`
	Namespaces     []string                    `json:"namespaces,omitempty"`
	Resources      []string                    `json:"resources,omitempty"`
	Ids            []string                    `json:"ids,omitempty"`
	Annotations    map[string]string           `json:"annotations,omitempty"`
}

func (s *ExtensionEventSelector) GetActions() []EventAction {
	return s.Actions
}
func (s *ExtensionEventSelector) GetRecordSelector() *ExtensionBooleanExpression {
	return s.RecordSelector
}
func (s *ExtensionEventSelector) GetNamespaces() []string {
	return s.Namespaces
}
func (s *ExtensionEventSelector) GetResources() []string {
	return s.Resources
}
func (s *ExtensionEventSelector) GetIds() []string {
	return s.Ids
}
func (s *ExtensionEventSelector) GetAnnotations() map[string]string {
	return s.Annotations
}

type ExtensionRecordSearchParams struct {
	Query             *ExtensionBooleanExpression `json:"query,omitempty"`
	Limit             *int32                      `json:"limit,omitempty"`
	Offset            *int32                      `json:"offset,omitempty"`
	ResolveReferences []string                    `json:"resolveReferences,omitempty"`
}

func (s *ExtensionRecordSearchParams) GetQuery() *ExtensionBooleanExpression {
	return s.Query
}
func (s *ExtensionRecordSearchParams) GetLimit() *int32 {
	return s.Limit
}
func (s *ExtensionRecordSearchParams) GetOffset() *int32 {
	return s.Offset
}
func (s *ExtensionRecordSearchParams) GetResolveReferences() []string {
	return s.ResolveReferences
}

type ExtensionEvent struct {
	Id                 string                       `json:"id,omitempty"`
	Action             ExtensionAction              `json:"action,omitempty"`
	RecordSearchParams *ExtensionRecordSearchParams `json:"recordSearchParams,omitempty"`
	ActionSummary      *string                      `json:"actionSummary,omitempty"`
	ActionDescription  *string                      `json:"actionDescription,omitempty"`
	Resource           *Resource                    `json:"resource,omitempty"`
	Records            []*Record                    `json:"records,omitempty"`
	Ids                []string                     `json:"ids,omitempty"`
	Finalizes          *bool                        `json:"finalizes,omitempty"`
	Sync               *bool                        `json:"sync,omitempty"`
	Time               *time.Time                   `json:"time,omitempty"`
	Annotations        map[string]string            `json:"annotations,omitempty"`
	Error              *ExtensionError              `json:"error,omitempty"`
}

func (s *ExtensionEvent) GetId() string {
	return s.Id
}
func (s *ExtensionEvent) GetAction() ExtensionAction {
	return s.Action
}
func (s *ExtensionEvent) GetRecordSearchParams() *ExtensionRecordSearchParams {
	return s.RecordSearchParams
}
func (s *ExtensionEvent) GetActionSummary() *string {
	return s.ActionSummary
}
func (s *ExtensionEvent) GetActionDescription() *string {
	return s.ActionDescription
}
func (s *ExtensionEvent) GetResource() *Resource {
	return s.Resource
}
func (s *ExtensionEvent) GetRecords() []*Record {
	return s.Records
}
func (s *ExtensionEvent) GetIds() []string {
	return s.Ids
}
func (s *ExtensionEvent) GetFinalizes() *bool {
	return s.Finalizes
}
func (s *ExtensionEvent) GetSync() *bool {
	return s.Sync
}
func (s *ExtensionEvent) GetTime() *time.Time {
	return s.Time
}
func (s *ExtensionEvent) GetAnnotations() map[string]string {
	return s.Annotations
}
func (s *ExtensionEvent) GetError() *ExtensionError {
	return s.Error
}

type ExtensionErrorField struct {
	RecordId *string     `json:"recordId,omitempty"`
	Property *string     `json:"property,omitempty"`
	Message  *string     `json:"message,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

func (s *ExtensionErrorField) GetRecordId() *string {
	return s.RecordId
}
func (s *ExtensionErrorField) GetProperty() *string {
	return s.Property
}
func (s *ExtensionErrorField) GetMessage() *string {
	return s.Message
}
func (s *ExtensionErrorField) GetValue() interface{} {
	return s.Value
}

type ExtensionError struct {
	Code    *ExtensionCode        `json:"code,omitempty"`
	Message *string               `json:"message,omitempty"`
	Fields  []ExtensionErrorField `json:"fields,omitempty"`
}

func (s *ExtensionError) GetCode() *ExtensionCode {
	return s.Code
}
func (s *ExtensionError) GetMessage() *string {
	return s.Message
}
func (s *ExtensionError) GetFields() []ExtensionErrorField {
	return s.Fields
}

type EventAction string

const (
	EventAction_CREATE  EventAction = "CREATE"
	EventAction_UPDATE  EventAction = "UPDATE"
	EventAction_DELETE  EventAction = "DELETE"
	EventAction_GET     EventAction = "GET"
	EventAction_LIST    EventAction = "LIST"
	EventAction_OPERATE EventAction = "OPERATE"
)

type ExtensionAction string

const (
	ExtensionAction_CREATE  ExtensionAction = "CREATE"
	ExtensionAction_UPDATE  ExtensionAction = "UPDATE"
	ExtensionAction_DELETE  ExtensionAction = "DELETE"
	ExtensionAction_GET     ExtensionAction = "GET"
	ExtensionAction_LIST    ExtensionAction = "LIST"
	ExtensionAction_OPERATE ExtensionAction = "OPERATE"
)

type ExtensionCode string

const (
	ExtensionCode_UNKNOWNERROR                      ExtensionCode = "UNKNOWN_ERROR"
	ExtensionCode_RECORDNOTFOUND                    ExtensionCode = "RECORD_NOT_FOUND"
	ExtensionCode_UNABLETOLOCATEPRIMARYKEY          ExtensionCode = "UNABLE_TO_LOCATE_PRIMARY_KEY"
	ExtensionCode_INTERNALERROR                     ExtensionCode = "INTERNAL_ERROR"
	ExtensionCode_PROPERTYNOTFOUND                  ExtensionCode = "PROPERTY_NOT_FOUND"
	ExtensionCode_RECORDVALIDATIONERROR             ExtensionCode = "RECORD_VALIDATION_ERROR"
	ExtensionCode_RESOURCEVALIDATIONERROR           ExtensionCode = "RESOURCE_VALIDATION_ERROR"
	ExtensionCode_AUTHENTICATIONFAILED              ExtensionCode = "AUTHENTICATION_FAILED"
	ExtensionCode_ALREADYEXISTS                     ExtensionCode = "ALREADY_EXISTS"
	ExtensionCode_ACCESSDENIED                      ExtensionCode = "ACCESS_DENIED"
	ExtensionCode_BACKENDERROR                      ExtensionCode = "BACKEND_ERROR"
	ExtensionCode_UNIQUEVIOLATION                   ExtensionCode = "UNIQUE_VIOLATION"
	ExtensionCode_REFERENCEVIOLATION                ExtensionCode = "REFERENCE_VIOLATION"
	ExtensionCode_RESOURCENOTFOUND                  ExtensionCode = "RESOURCE_NOT_FOUND"
	ExtensionCode_UNSUPPORTEDOPERATION              ExtensionCode = "UNSUPPORTED_OPERATION"
	ExtensionCode_EXTERNALBACKENDCOMMUNICATIONERROR ExtensionCode = "EXTERNAL_BACKEND_COMMUNICATION_ERROR"
	ExtensionCode_EXTERNALBACKENDERROR              ExtensionCode = "EXTERNAL_BACKEND_ERROR"
)
