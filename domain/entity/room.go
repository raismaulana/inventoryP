package entity

import (
	"time"
)

type Room struct {
	ID              uint            `gorm:"primary_key:auto_increment;column:room_id"` //
	Name            string          ``                                                 //
	RoomInventories []RoomInventory ``                                                 //
	AuditItems      []AuditItem     ``                                                 //
	CreatedAt       time.Time       ``                                                 //
	UpdatedAt       time.Time       ``                                                 //
}
