package entity

import (
	"time"

	"github.com/raismaulana/inventoryP/application/apperror"
)

type Audit struct {
	ID        uint      `gorm:"primary_key:auto_increment;column:audit_id"` //
	Date      time.Time ``                                                  //
	Items     []Item    `gorm:"many2many:audit_items;"`                     //
	CreatedAt time.Time ``                                                  //
	UpdatedAt time.Time ``                                                  //
}

type AuditRequest struct {
	ID   uint      `gorm:"primary_key:auto_increment;column:audit_id"` //
	Date time.Time ``                                                  //
}

func NewAudit(req AuditRequest) (*Audit, error) {

	if req.Date.IsZero() {
		return nil, apperror.DateMustNotEmpty
	}

	return &Audit{
		ID:   req.ID,
		Date: req.Date,
	}, nil
}
