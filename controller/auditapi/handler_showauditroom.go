package auditapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/inventoryP/application/apperror"
	"github.com/raismaulana/inventoryP/infrastructure/log"
	"github.com/raismaulana/inventoryP/infrastructure/util"
	"github.com/raismaulana/inventoryP/usecase/showauditroom"
)

// showAuditRoomHandler ...
func (r *Controller) showAuditRoomHandler(inputPort showauditroom.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithLogGroupID(c.Request.Context())

		var req showauditroom.InportRequest
		if err := c.ShouldBindUri(&req); err != nil {
			newErr := util.GetValidationErrorMessage(err)
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.ERR400.Var(newErr)))
			return
		}

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
