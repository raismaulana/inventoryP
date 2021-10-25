package showaudit

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of ShowAudit
type Outport interface {
	repository.WithoutTransactionDB
	repository.FetchAuditRepo
}
