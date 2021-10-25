package prod

import (
	"context"

	"github.com/google/uuid"
	"github.com/raismaulana/inventoryP/infrastructure/log"
)

func (r *prodGateway) GenerateUUID(ctx context.Context) (string, error) {
	log.Info(ctx, "called")
	a, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return a.String(), nil
}
