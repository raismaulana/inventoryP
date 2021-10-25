package auditapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/infrastructure/log"
	"github.com/raismaulana/inventoryP/infrastructure/util"
	"github.com/raismaulana/inventoryP/usecase/createaudit"
)

// createAuditHandler ...
func (r *Controller) createAuditHandler(inputPort createaudit.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithLogGroupID(c.Request.Context())

		var req createaudit.InportRequest

		log.InfoRequest(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(apperror.GetErrorCode(err), NewErrorResponse(err))
			return
		}

		log.InfoResponse(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
