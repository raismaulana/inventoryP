package showroom

import (
	"context"
)

// Inport of ShowRoom
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowRoom
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowRoom
type InportResponse struct {
	Rooms []RoomResponse `json:"room"`
}

type RoomResponse struct {
	ID          uint                `json:"room_id"`   //
	Name        string              `json:"name"`      //
	Inventories []InventoryResponse `json:"inventory"` //
}

type InventoryResponse struct {
	ID       uint   `json:"inventory_id"` //
	Name     string `json:"name"`         //
	Quantity int    `json:"quantity"`     //
	Unit     string `json:"unit"`         //
}
