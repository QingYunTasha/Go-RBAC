package http

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	RoleUsecase UsecaseDomain.RoleUsecase
}

func NewRoleHandler(server *gin.Engine, usecase UsecaseDomain.RoleUsecase) {
	handler := &RoleHandler{
		RoleUsecase: usecase,
	}

	user := server.Group("/roles")
	user.GET("/", handler.GetAll)
	user.GET("/:name", handler.Get)
	user.POST("/", handler.Create)
	user.PUT("/:name", handler.Update)
	user.DELETE("/:name", handler.Delete)
}

func (rlh *RoleHandler) GetAll(c *gin.Context) {
	res, err := rlh.RoleUsecase.GetAll(context.TODO())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (rlh *RoleHandler) Get(c *gin.Context) {
	role, err := rlh.RoleUsecase.Get(context.TODO(), c.Param("name"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, role)
}

func (rlh *RoleHandler) Create(c *gin.Context) {
	user := &RepoDomain.Role{
		Name: c.PostForm("name"),
	}
	if err := rlh.RoleUsecase.Create(context.TODO(), user); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (rlh *RoleHandler) Update(c *gin.Context) {
	user := &RepoDomain.Role{
		Name: c.PostForm("name"),
	}
	if err := rlh.RoleUsecase.Update(context.TODO(), c.Param("name"), user); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (rlh *RoleHandler) Delete(c *gin.Context) {
	if err := rlh.RoleUsecase.Delete(context.TODO(), c.Param("name")); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}
