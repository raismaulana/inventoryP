package showaudit

import (
	"context"

	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAuditInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAudit
func NewUsecase(outputPort Outport) Inport {
	return &showAuditInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAudit
func (r *showAuditInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithoutTrx(ctx, r.outport, func(dbCtx context.Context) error {
		auditObj, err := r.outport.FetchAudit(dbCtx)
		if err != nil {
			return err
		}
		if len(auditObj) == 0 {
			return apperror.ObjectNotFound.Var(auditObj)
		}
		for _, v := range auditObj {
			res.Audits = append(res.Audits, AuditResponse{
				ID:   v.ID,
				Date: v.Date,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
