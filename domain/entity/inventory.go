package entity

import "time"

type Inventory struct {
	ID              uint            `gorm:"primary_key:auto_increment;column:inventory_id"` //
	Name            string          ``                                                      //
	Unit            string          ``                                                      //
	Quantity        int             ``                                                      //
	RoomInventories []RoomInventory ``                                                      //
	Items           []Item          ``                                                      //
	CreatedAt       time.Time       ``                                                      //
	UpdatedAt       time.Time       ``                                                      //
}
