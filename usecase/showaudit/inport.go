package showaudit

import (
	"context"
	"time"
)

// Inport of ShowAudit
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAudit
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAudit
type InportResponse struct {
	Audits []AuditResponse `json:"audits"` //
}

type AuditResponse struct {
	ID   uint      `json:"audit_id"` //
	Date time.Time `json:"date"`     //
}
