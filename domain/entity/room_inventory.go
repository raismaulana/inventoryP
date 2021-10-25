package entity

import "time"

type RoomInventory struct {
	RoomID      uint      `gorm:"primary_key"`                        //
	InventoryID uint      `gorm:"primary_key"`                        //
	Quantity    int       ``                                          //
	Room        Room      ``                                          //
	Inventory   Inventory ``                                          //
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"` //
}
