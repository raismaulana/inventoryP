package createaudit

import (
	"context"
)

// Inport of CreateAudit
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase CreateAudit
type InportRequest struct {
}

// InportResponse is response payload after running the usecase CreateAudit
type InportResponse struct {
	AuditID uint `json:"audit_id"`
}
