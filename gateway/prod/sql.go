package prod

import (
	"context"

	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/domain/entity"
	"github.com/raismaulana/inventoryP/domain/repository"
	"github.com/raismaulana/inventoryP/infrastructure/database"
	"github.com/raismaulana/inventoryP/infrastructure/log"
)

type sqlGateway struct {
	database.GormWithTrxImpl
	database.GormWithoutTrxImpl
}

func (r *prodGateway) CreateAudit(ctx context.Context, obj *entity.Audit) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return apperror.DatabaseNotFoundInContextError
	}
	err = db.Create(obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	return nil
}
func (r *prodGateway) SaveAudit(ctx context.Context, obj *entity.Audit) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return apperror.DatabaseNotFoundInContextError
	}
	err = db.Save(obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (r *prodGateway) CreateAuditItem(ctx context.Context, obj *entity.AuditItem) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return apperror.DatabaseNotFoundInContextError
	}
	err = db.Create(obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return apperror.ItemIsAlreadyScannedInThisSession
	}

	return nil
}

func (r *prodGateway) FetchRoom(ctx context.Context) ([]*entity.Room, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, apperror.DatabaseNotFoundInContextError
	}
	var objs []*entity.Room
	err = db.Preload("RoomInventories.Inventory").Find(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *prodGateway) FetchAudit(ctx context.Context) ([]*entity.Audit, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, apperror.DatabaseNotFoundInContextError
	}
	var objs []*entity.Audit
	err = db.Find(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *prodGateway) FindRoomByAuditID(ctx context.Context, auditID uint) ([]*entity.Room, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, apperror.DatabaseNotFoundInContextError
	}
	var objs []*entity.Room
	err = db.Joins("JOIN audit_items ai ON rooms.room_id = ai.room_id AND ai.audit_id = ?", auditID).Group("rooms.room_id").Find(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *prodGateway) FindInventoriesByAuditAndRoom(ctx context.Context, auditID uint, roomID uint) ([]*repository.FindInventoriesByAuditAndRoomResponse, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, apperror.DatabaseNotFoundInContextError
	}
	var objs []*repository.FindInventoriesByAuditAndRoomResponse
	err = db.Raw("select i.inventory_id, i.name, count(i.inventory_id) quantity from inventories i join items i2 on i.inventory_id = i2.inventory_id join audit_items ai on i2.item_id = ai.item_id and ai.audit_id = 1 and ai.room_id = 1 group by i.name, i.inventory_id").Scan(&objs).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return objs, nil
}

func (r *prodGateway) FindInventoryByRoomID(ctx context.Context, roomID uint) (*entity.Room, error) {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, apperror.DatabaseNotFoundInContextError
	}
	var obj *entity.Room
	err = db.Preload("RoomInventories.Inventory").First(&obj, roomID).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}
	return obj, nil
}
