package registry

import (
	"github.com/raismaulana/inventoryP/application"
	"github.com/raismaulana/inventoryP/controller"
	"github.com/raismaulana/inventoryP/controller/auditapi"
	"github.com/raismaulana/inventoryP/controller/roomapi"
	"github.com/raismaulana/inventoryP/gateway/prod"
	"github.com/raismaulana/inventoryP/infrastructure/auth"
	"github.com/raismaulana/inventoryP/infrastructure/database"
	"github.com/raismaulana/inventoryP/infrastructure/env"
	"github.com/raismaulana/inventoryP/infrastructure/log"
	"github.com/raismaulana/inventoryP/infrastructure/server"
	"github.com/raismaulana/inventoryP/usecase/addaudititem"
	"github.com/raismaulana/inventoryP/usecase/createaudit"
	"github.com/raismaulana/inventoryP/usecase/showaudit"
	"github.com/raismaulana/inventoryP/usecase/showauditroom"
	"github.com/raismaulana/inventoryP/usecase/showinventoriesinauditroom"
	"github.com/raismaulana/inventoryP/usecase/showroom"
	"github.com/raismaulana/inventoryP/usecase/showroominventory"
)

type app struct {
	server.GinHTTPHandler
	auditapi auditapi.Controller
	roomapi  roomapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

func NewApp() func() application.RegistryContract {
	return func() application.RegistryContract {
		// setup logger write file
		log.UseRotateFile(".log", "default", 30)

		// setup db
		db := database.NewGormPostgres()

		enforcer := auth.NewCasbinEnforcerByDB(db)

		// setup server
		httpHandler := server.NewGinHTTPHandler(env.Var().AppPort)

		datasource := prod.NewProdGateway(db)

		return &app{
			GinHTTPHandler: httpHandler,
			auditapi: auditapi.Controller{
				BaseController:                   controller.BaseController{Enforcer: enforcer, Router: httpHandler.Router},
				CreateAuditInport:                createaudit.NewUsecase(datasource),
				AddAuditItemInport:               addaudititem.NewUsecase(datasource),
				ShowAuditInport:                  showaudit.NewUsecase(datasource),
				ShowAuditRoomInport:              showauditroom.NewUsecase(datasource),
				ShowInventoriesInAuditRoomInport: showinventoriesinauditroom.NewUsecase(datasource),
			},
			roomapi: roomapi.Controller{
				BaseController:          controller.BaseController{Enforcer: enforcer, Router: httpHandler.Router},
				ShowRoomInport:          showroom.NewUsecase(datasource),
				ShowRoomInventoryInport: showroominventory.NewUsecase(datasource),
			},
		}

	}
}

func (r *app) SetupController() {
	r.auditapi.RegisterRouter()
	r.roomapi.RegisterRouter()
	// TODO another router call will added here ... <<<<<<
}
