package controller

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	Enforcer *casbin.Enforcer
	Router   gin.IRouter
}

func (r *BaseController) Authorization() {

}
