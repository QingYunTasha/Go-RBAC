package http

import (
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	PermissionUsecase UsecaseDomain.PermissionUsecase
}

type PermissionParams struct {
	Email    string `url:"email"`
	RoleName string `url:"rolename"`
}

func NewPermissionHandler(server *gin.Engine, usecase UsecaseDomain.PermissionUsecase) {
	handler := &PermissionHandler{
		PermissionUsecase: usecase,
	}

	user := server.Group("/permissions")
	user.GET("/", handler.GetAll)
	user.GET("/:resourcename", handler.GetByResource)
	user.POST("/", handler.Create)
	user.DELETE("/:resourcename/:operation", handler.Delete)
}

func (pmh *PermissionHandler) GetAll(c *gin.Context) {
	permissions, err := pmh.PermissionUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func (pmh *PermissionHandler) GetByResource(c *gin.Context) {
	permissions, err := pmh.PermissionUsecase.GetByResource(c.Request.Context(), c.Param("resourcename"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func (pmh *PermissionHandler) Create(c *gin.Context) {
	permission := RepoDomain.Permission{}
	if err := c.BindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pmh.PermissionUsecase.Create(c.Request.Context(), &permission); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, "success")
}

func (pmh *PermissionHandler) Delete(c *gin.Context) {
	if err := pmh.PermissionUsecase.Delete(c.Request.Context(), c.Param("resourcename"), c.Param("action")); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, "success")
}
