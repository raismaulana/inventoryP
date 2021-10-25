package showauditroom

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of ShowAuditRoom
type Outport interface {
	repository.WithoutTransactionDB
	repository.FindRoomByAuditIDRepo
}
