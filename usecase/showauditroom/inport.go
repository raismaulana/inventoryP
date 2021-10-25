package showauditroom

import (
	"context"
)

// Inport of ShowAuditRoom
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAuditRoom
type InportRequest struct {
	AuditID uint `json:"audit_id" uri:"audit_id" binding:"required"`
}

// InportResponse is response payload after running the usecase ShowAuditRoom
type InportResponse struct {
	Rooms []RoomResponse `json:"rooms"` //
}

type RoomResponse struct {
	ID   uint   `json:"room_id"` //
	Name string `json:"name"`    //
}
