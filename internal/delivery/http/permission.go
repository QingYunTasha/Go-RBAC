package http

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"

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
	permissions, err := pmh.PermissionUsecase.GetAll(context.TODO())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, permissions)
}

func (pmh *PermissionHandler) GetByResource(c *gin.Context) {
	permissions, err := pmh.PermissionUsecase.GetByResource(context.TODO(), c.Param("resourcename"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, permissions)
}

func (pmh *PermissionHandler) Create(c *gin.Context) {
	permission := RepoDomain.Permission{}
	if err := c.BindJSON(&permission); err != nil {
		c.JSON(400, err.Error())
	}

	if err := pmh.PermissionUsecase.Create(context.TODO(), &permission); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (pmh *PermissionHandler) Delete(c *gin.Context) {
	if err := pmh.PermissionUsecase.Delete(context.TODO(), c.Param("resourcename"), c.Param("operation")); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}
