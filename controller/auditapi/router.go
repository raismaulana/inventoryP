package auditapi

import (
	"github.com/raismaulana/inventoryP/controller"
	"github.com/raismaulana/inventoryP/usecase/addaudititem"
	"github.com/raismaulana/inventoryP/usecase/createaudit"
	"github.com/raismaulana/inventoryP/usecase/showaudit"
	"github.com/raismaulana/inventoryP/usecase/showauditroom"
	"github.com/raismaulana/inventoryP/usecase/showinventoriesinauditroom"
)

type Controller struct {
	controller.BaseController
	CreateAuditInport                createaudit.Inport
	AddAuditItemInport               addaudititem.Inport
	ShowAuditInport                  showaudit.Inport
	ShowAuditRoomInport              showauditroom.Inport
	ShowInventoriesInAuditRoomInport showinventoriesinauditroom.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/audit", r.authorized(), r.createAuditHandler(r.CreateAuditInport))
	r.Router.POST("/audit/item", r.authorized(), r.addAuditItemHandler(r.AddAuditItemInport))
	r.Router.GET("/audit", r.authorized(), r.showAuditHandler(r.ShowAuditInport))
	r.Router.GET("/audit/:audit_id/room", r.authorized(), r.showAuditRoomHandler(r.ShowAuditRoomInport))
	r.Router.GET("/audit/:audit_id/room/:room_id/item", r.authorized(), r.showInventoriesInAuditRoomHandler(r.ShowInventoriesInAuditRoomInport))
}
