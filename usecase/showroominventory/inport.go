package showroominventory

import (
	"context"
)

// Inport of ShowRoomInventory
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowRoomInventory
type InportRequest struct {
	RoomID uint `uri:"room_id" binding:"required"`
}

// InportResponse is response payload after running the usecase ShowRoomInventory
type InportResponse struct {
	Inventory []InventoryResponse `` //
}

type InventoryResponse struct {
	InventoryID uint   `json:"inventory_id"` //
	Name        string `json:"name"`         //
	Quantity    int    `json:"quantity"`     //
}
