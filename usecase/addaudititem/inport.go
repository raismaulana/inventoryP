package addaudititem

import (
	"context"
)

// Inport of AddAuditItem
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase AddAuditItem
type InportRequest struct {
	AuditID uint `json:"audit_id" binding:"required"`
	ItemID  uint `json:"item_id" binding:"required"`
	RoomID  uint `json:"room_id"  binding:"required"`
}

// InportResponse is response payload after running the usecase AddAuditItem
type InportResponse struct {
}
