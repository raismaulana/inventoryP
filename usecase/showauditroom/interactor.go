package showauditroom

import (
	"context"

	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAuditRoomInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAuditRoom
func NewUsecase(outputPort Outport) Inport {
	return &showAuditRoomInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAuditRoom
func (r *showAuditRoomInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	err := repository.WithoutTrx(ctx, r.outport, func(dbCtx context.Context) error {

		roomObjs, err := r.outport.FindRoomByAuditID(dbCtx, req.AuditID)
		if err != nil {
			return err
		}
		if len(roomObjs) == 0 {
			return apperror.ObjectNotFound.Var(roomObjs)
		}
		for _, v := range roomObjs {
			res.Rooms = append(res.Rooms, RoomResponse{
				ID:   v.ID,
				Name: v.Name,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
