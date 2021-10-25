package showinventoriesinauditroom

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of ShowAuditRoomItem
type Outport interface {
	repository.WithoutTransactionDB
	repository.FindInventoriesByAuditAndRoomRepo
}
