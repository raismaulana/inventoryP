package showinventoriesinauditroom

import (
	"context"

	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showInventoriesInAuditRoomInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAuditRoomItem
func NewUsecase(outputPort Outport) Inport {
	return &showInventoriesInAuditRoomInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAuditRoomItem
func (r *showInventoriesInAuditRoomInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithoutTrx(ctx, r.outport, func(dbCtx context.Context) error {

		inventoryObjs, err := r.outport.FindInventoriesByAuditAndRoom(dbCtx, req.AuditID, req.RoomID)
		if err != nil {
			return err
		}
		if len(inventoryObjs) == 0 {
			return apperror.ObjectNotFound.Var(inventoryObjs)
		}
		for _, v := range inventoryObjs {
			res.Inventory = append(res.Inventory, InventoryResponse{
				InventoryID: v.InventoryID,
				Name:        v.Name,
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
