package createaudit

import (
	"context"
	"time"

	"github.com/raismaulana/inventoryP/domain/entity"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type createAuditInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreateAudit
func NewUsecase(outputPort Outport) Inport {
	return &createAuditInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreateAudit
func (r *createAuditInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTrx(ctx, r.outport, func(dbCtx context.Context) error {
		audit, err := entity.NewAudit(entity.AuditRequest{
			Date: time.Now(),
		})
		if err != nil {
			return err
		}
		err = r.outport.CreateAudit(dbCtx, audit)
		if err != nil {
			return err
		}

		res.AuditID = audit.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
