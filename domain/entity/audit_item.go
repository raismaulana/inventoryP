package entity

import (
	"time"

	"github.com/raismaulana/inventoryP/application/apperror"
)

type AuditItem struct {
	AuditID   uint      `gorm:"primary_key"`                        //
	ItemID    uint      `gorm:"primary_key"`                        //
	RoomID    uint      ``                                          //
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` //
}

type AuditItemRequest struct {
	AuditID uint `` //
	ItemID  uint `` //
	RoomID  uint `` //
}

func NewAuditItem(req AuditItemRequest) (*AuditItem, error) {
	if req.AuditID == 0 {
		return nil, apperror.AuditIDMustNotEmpty
	}
	if req.ItemID == 0 {
		return nil, apperror.ItemIDMustNotEmpty
	}
	if req.RoomID == 0 {
		return nil, apperror.RoomIDMustNotEmpty
	}

	return &AuditItem{
		AuditID: req.AuditID,
		ItemID:  req.ItemID,
		RoomID:  req.RoomID,
	}, nil
}
