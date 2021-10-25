package roomapi

import (
	"github.com/raismaulana/inventoryP/controller"
	"github.com/raismaulana/inventoryP/usecase/showroom"
	"github.com/raismaulana/inventoryP/usecase/showroominventory"
)

type Controller struct {
	controller.BaseController
	ShowRoomInport          showroom.Inport
	ShowRoomInventoryInport showroominventory.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/room", r.authorized(), r.showRoomHandler(r.ShowRoomInport))
	r.Router.GET("/room/:room_id/inventory", r.authorized(), r.showRoomInventoryHandler(r.ShowRoomInventoryInport))
}
