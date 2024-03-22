// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type ResourceAction struct {
	Id          *uuid.UUID               `json:"id,omitempty"`
	Version     int32                    `json:"version,omitempty"`
	AuditData   *ResourceActionAuditData `json:"auditData,omitempty"`
	Resource    *Resource                `json:"resource,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Title       *string                  `json:"title,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Internal    bool                     `json:"internal,omitempty"`
	Types       []SubType                `json:"types,omitempty"`
	Input       map[string]Property      `json:"input,omitempty"`
	Output      *Property                `json:"output,omitempty"`
	Annotations map[string]string        `json:"annotations,omitempty"`
}

func (s *ResourceAction) GetId() *uuid.UUID {
	return s.Id
}
func (s *ResourceAction) GetVersion() int32 {
	return s.Version
}
func (s *ResourceAction) GetAuditData() *ResourceActionAuditData {
	return s.AuditData
}
func (s *ResourceAction) GetResource() *Resource {
	return s.Resource
}
func (s *ResourceAction) GetName() string {
	return s.Name
}
func (s *ResourceAction) GetTitle() *string {
	return s.Title
}
func (s *ResourceAction) GetDescription() *string {
	return s.Description
}
func (s *ResourceAction) GetInternal() bool {
	return s.Internal
}
func (s *ResourceAction) GetTypes() []SubType {
	return s.Types
}
func (s *ResourceAction) GetInput() map[string]Property {
	return s.Input
}
func (s *ResourceAction) GetOutput() *Property {
	return s.Output
}
func (s ResourceAction) GetAnnotations() map[string]string {
	return s.Annotations
}

type ResourceActionAuditData struct {
	CreatedBy *string    `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}

func (s *ResourceActionAuditData) GetCreatedBy() *string {
	return s.CreatedBy
}
func (s *ResourceActionAuditData) GetUpdatedBy() *string {
	return s.UpdatedBy
}
func (s *ResourceActionAuditData) GetCreatedOn() *time.Time {
	return s.CreatedOn
}
func (s *ResourceActionAuditData) GetUpdatedOn() *time.Time {
	return s.UpdatedOn
}

type ResourceActionType string

const (
	ResourceActionType_BOOL      ResourceActionType = "BOOL"
	ResourceActionType_STRING    ResourceActionType = "STRING"
	ResourceActionType_FLOAT32   ResourceActionType = "FLOAT32"
	ResourceActionType_FLOAT64   ResourceActionType = "FLOAT64"
	ResourceActionType_INT32     ResourceActionType = "INT32"
	ResourceActionType_INT64     ResourceActionType = "INT64"
	ResourceActionType_BYTES     ResourceActionType = "BYTES"
	ResourceActionType_UUID      ResourceActionType = "UUID"
	ResourceActionType_DATE      ResourceActionType = "DATE"
	ResourceActionType_TIME      ResourceActionType = "TIME"
	ResourceActionType_TIMESTAMP ResourceActionType = "TIMESTAMP"
	ResourceActionType_OBJECT    ResourceActionType = "OBJECT"
	ResourceActionType_MAP       ResourceActionType = "MAP"
	ResourceActionType_LIST      ResourceActionType = "LIST"
	ResourceActionType_REFERENCE ResourceActionType = "REFERENCE"
	ResourceActionType_ENUM      ResourceActionType = "ENUM"
	ResourceActionType_STRUCT    ResourceActionType = "STRUCT"
)
