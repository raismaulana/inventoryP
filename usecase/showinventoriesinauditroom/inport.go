package showinventoriesinauditroom

import (
	"context"
)

// Inport of ShowAuditRoomItem
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAuditRoomItem
type InportRequest struct {
	AuditID uint `uri:"audit_id" binding:"required"` //
	RoomID  uint `uri:"room_id" binding:"required"`  //
}

// InportResponse is response payload after running the usecase ShowAuditRoomItem
type InportResponse struct {
	Inventory []InventoryResponse `` //
}

type InventoryResponse struct {
	InventoryID uint   `json:"inventory_id"` //
	Name        string `json:"name"`         //
	Quantity    int    `json:"quantity"`     //
}
