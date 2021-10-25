package showroom

import (
	"context"

	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showRoomInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowRoom
func NewUsecase(outputPort Outport) Inport {
	return &showRoomInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowRoom
func (r *showRoomInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithoutTrx(ctx, r.outport, func(dbCtx context.Context) error {

		roomObj, err := r.outport.FetchRoom(dbCtx)
		if err != nil {
			return err
		}
		if roomObj == nil {
			return apperror.ObjectNotFound.Var(roomObj)
		}
		for _, r := range roomObj {
			roomRes := RoomResponse{
				ID:   r.ID,
				Name: r.Name,
			}
			for _, ri := range r.RoomInventories {
				roomRes.Inventories = append(roomRes.Inventories, InventoryResponse{
					ID:       ri.InventoryID,
					Name:     ri.Inventory.Name,
					Quantity: ri.Quantity,
					Unit:     ri.Inventory.Unit,
				})
			}
			res.Rooms = append(res.Rooms, roomRes)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
