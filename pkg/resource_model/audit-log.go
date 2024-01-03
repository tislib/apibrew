// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package resource_model

import "github.com/google/uuid"
import "time"

type AuditLog struct {
	Id          *uuid.UUID        `json:"id,omitempty"`
	Version     int32             `json:"version,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Resource    string            `json:"resource,omitempty"`
	RecordId    string            `json:"recordId,omitempty"`
	Time        time.Time         `json:"time,omitempty"`
	Username    string            `json:"username,omitempty"`
	Operation   AuditLogOperation `json:"operation,omitempty"`
	Properties  interface{}       `json:"properties,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (s *AuditLog) GetId() *uuid.UUID {
	return s.Id
}
func (s *AuditLog) GetVersion() int32 {
	return s.Version
}
func (s *AuditLog) GetNamespace() string {
	return s.Namespace
}
func (s *AuditLog) GetResource() string {
	return s.Resource
}
func (s *AuditLog) GetRecordId() string {
	return s.RecordId
}
func (s *AuditLog) GetTime() time.Time {
	return s.Time
}
func (s *AuditLog) GetUsername() string {
	return s.Username
}
func (s *AuditLog) GetOperation() AuditLogOperation {
	return s.Operation
}
func (s *AuditLog) GetProperties() interface{} {
	return s.Properties
}
func (s *AuditLog) GetAnnotations() map[string]string {
	return s.Annotations
}

type AuditLogOperation string

const (
	AuditLogOperation_CREATE AuditLogOperation = "CREATE"
	AuditLogOperation_UPDATE AuditLogOperation = "UPDATE"
	AuditLogOperation_DELETE AuditLogOperation = "DELETE"
)
