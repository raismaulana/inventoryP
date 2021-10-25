package entity

import "time"

type Item struct {
	ID          uint      `gorm:"primary_key:auto_increment;column:item_id"` //
	InventoryID uint      ``                                                 //
	CreatedAt   time.Time ``                                                 //
	UpdatedAt   time.Time ``                                                 //
}
