package addaudititem

import (
	"context"

	"github.com/raismaulana/inventoryP/domain/entity"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type addAuditItemInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase AddAuditItem
func NewUsecase(outputPort Outport) Inport {
	return &addAuditItemInteractor{
		outport: outputPort,
	}
}

// Execute the usecase AddAuditItem
func (r *addAuditItemInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithTrx(ctx, r.outport, func(dbCtx context.Context) error {

		auditItemObj, err := entity.NewAuditItem(entity.AuditItemRequest{
			AuditID: req.AuditID,
			ItemID:  req.ItemID,
			RoomID:  req.RoomID,
		})
		if err != nil {
			return err
		}

		err = r.outport.CreateAuditItem(dbCtx, auditItemObj)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
