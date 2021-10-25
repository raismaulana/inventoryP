package addaudititem

import "github.com/raismaulana/inventoryP/domain/repository"

// Outport of AddAuditItem
type Outport interface {
	repository.WithTransactionDB
	repository.CreateAuditItemRepo
}
