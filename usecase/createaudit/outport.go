package createaudit

import (
	"github.com/raismaulana/inventoryP/domain/repository"
	"github.com/raismaulana/inventoryP/domain/service"
)

// Outport of CreateAudit
type Outport interface {
	repository.WithTransactionDB
	service.GenerateUUIDService
	repository.CreateAuditRepo
}
