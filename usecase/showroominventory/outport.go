package showroominventory

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of ShowRoomInventory
type Outport interface {
	repository.WithoutTransactionDB
	repository.FindInventoryByRoomIDRepo
}
