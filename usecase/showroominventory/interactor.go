package showroominventory

import (
	"context"

	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showRoomInventoryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowRoomInventory
func NewUsecase(outputPort Outport) Inport {
	return &showRoomInventoryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowRoomInventory
func (r *showRoomInventoryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithoutTrx(ctx, r.outport, func(dbCtx context.Context) error {

		roomObjs, err := r.outport.FindInventoryByRoomID(dbCtx, req.RoomID)
		if err != nil {
			return err
		}
		for _, v := range roomObjs.RoomInventories {
			res.Inventory = append(res.Inventory, InventoryResponse{
				InventoryID: v.InventoryID,
				Name:        v.Inventory.Name,
				Quantity:    v.Quantity,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
