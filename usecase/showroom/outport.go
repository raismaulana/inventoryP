package showroom

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of ShowRoom
type Outport interface {
	repository.FetchRoomRepo
	repository.WithoutTransactionDB
}
