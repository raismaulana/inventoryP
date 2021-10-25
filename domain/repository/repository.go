package repository

import (
	"context"

	"github.com/raismaulana/inventoryP/domain/entity"
)

type SaveAuditRepo interface {
	SaveAudit(ctx context.Context, obj *entity.Audit) error
}

type CreateAuditRepo interface {
	CreateAudit(ctx context.Context, obj *entity.Audit) error
}

type CreateAuditItemRepo interface {
	CreateAuditItem(ctx context.Context, obj *entity.AuditItem) error
}

type FetchRoomRepo interface {
	FetchRoom(ctx context.Context) ([]*entity.Room, error)
}

type FetchAuditRepo interface {
	FetchAudit(ctx context.Context) ([]*entity.Audit, error)
}

type FindRoomByAuditIDRepo interface {
	FindRoomByAuditID(ctx context.Context, auditID uint) ([]*entity.Room, error)
}

type FindInventoriesByAuditAndRoomRepo interface {
	FindInventoriesByAuditAndRoom(ctx context.Context, auditID uint, roomID uint) ([]*FindInventoriesByAuditAndRoomResponse, error)
}
type FindInventoriesByAuditAndRoomResponse struct {
	InventoryID uint   `` //
	Name        string `` //
	Quantity    int    `` //
}

type FindInventoryByRoomIDRepo interface {
	FindInventoryByRoomID(ctx context.Context, roomID uint) (*entity.Room, error)
}
