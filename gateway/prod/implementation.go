package prod

import (
	"github.com/raismaulana/inventoryP/infrastructure/database"
	"gorm.io/gorm"
)

type prodGateway struct {
	sqlGateway
}

// NewProdGateway ...
func NewProdGateway(db *gorm.DB) *prodGateway {
	return &prodGateway{
		sqlGateway: sqlGateway{
			GormWithTrxImpl: database.GormWithTrxImpl{
				DB: db,
			},
			GormWithoutTrxImpl: database.GormWithoutTrxImpl{
				DB: db,
			},
		},
	}
}
