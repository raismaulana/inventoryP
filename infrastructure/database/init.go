package database

import (
	"fmt"

	"github.com/raismaulana/inventoryP/domain/entity"
	"github.com/raismaulana/inventoryP/infrastructure/env"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initGorm(dialector gorm.Dialector, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		panic(err.Error())
	}
	if err = autoMigrate(db); err != nil {
		panic(err.Error())
	}
	if env.GetBool("SEED", false) {
		if err = seeder(db); err != nil {
			panic(err.Error())
		}
	}
	return db
}

// NewGormDefault ...
func NewGormDefault() *gorm.DB {
	return initGorm(sqlite.Open("default.db"), &gorm.Config{})
}

func NewGormPostgres() *gorm.DB {
	var config gorm.Config
	if env.Var().Production {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.Var().DBHost,
		env.Var().DBUser,
		env.Var().DBPassword,
		env.Var().DBName,
		env.Var().DBPort,
	)
	return initGorm(postgres.Open(dsn), &config)
}

func autoMigrate(db *gorm.DB) error {
	if err := db.SetupJoinTable(&entity.Audit{}, "Items", &entity.AuditItem{}); err != nil {
		return err
	}
	err := db.AutoMigrate(&entity.Room{}, &entity.Inventory{}, &entity.RoomInventory{}, &entity.Item{}, &entity.Audit{})
	if err != nil {
		return err
	}
	return nil
}

func seeder(db *gorm.DB) error {
	rooms := [...]entity.Room{
		{
			ID:   1,
			Name: "Kelas X",
		},
		{
			ID:   2,
			Name: "Kelas XI",
		},
		{
			ID:   3,
			Name: "Kelas XII",
		},
	}
	if err := db.Create(&rooms).Error; err != nil {
		return err
	}

	inventorys := [...]entity.Inventory{
		{
			ID:       1,
			Name:     "Meja",
			Unit:     "Buah",
			Quantity: 60,
		},
		{
			ID:       2,
			Name:     "Kursi",
			Unit:     "Buah",
			Quantity: 120,
		},
		{
			ID:       3,
			Name:     "Papan Tulis",
			Unit:     "Buah",
			Quantity: 3,
		},
	}
	if err := db.Create(&inventorys).Error; err != nil {
		return err
	}

	var room_inventorys []entity.RoomInventory
	for _, r := range rooms {
		for _, i := range inventorys {
			room_inventorys = append(room_inventorys, entity.RoomInventory{
				RoomID:      r.ID,
				InventoryID: i.ID,
				Quantity:    i.Quantity / len(rooms),
			})
		}
	}
	if err := db.Create(&room_inventorys).Error; err != nil {
		return err
	}

	var items []entity.Item
	for _, v := range inventorys {
		for i := 0; i < v.Quantity; i++ {
			items = append(items, entity.Item{
				InventoryID: v.ID,
			})
		}
	}
	if err := db.Create(&items).Error; err != nil {
		return err
	}

	return nil
}
