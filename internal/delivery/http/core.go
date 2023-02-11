package http

import (
	UsecaseDomain "go-authorization/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoreHandler struct {
	CoreUsecase UsecaseDomain.CoreUsecase
}

func NewCoreHandler(server *gin.Engine, usecase UsecaseDomain.CoreUsecase) {
	handler := &CoreHandler{
		CoreUsecase: usecase,
	}

	user := server.Group("/core")
	user.POST("/checkpermission", handler.CheckPermission)
}

type CheckPermissionParam struct {
	UserEmail string
	Action    string
	Resource  string
}

func (crh *CoreHandler) CheckPermission(c *gin.Context) {
	checkPermissionParam := CheckPermissionParam{}
	if err := c.BindJSON(&checkPermissionParam); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hasPermission, err := crh.CoreUsecase.CheckPermission(c.Request.Context(), checkPermissionParam.UserEmail, checkPermissionParam.Action, checkPermissionParam.Resource)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if hasPermission {
		c.JSON(http.StatusOK, "Yes")
	} else {
		c.JSON(http.StatusForbidden, "No")
	}
}
