package http

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase UsecaseDomain.UserUsecase
}

func NewUserHandler(server *gin.Engine, usecase UsecaseDomain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	user := server.Group("/users")
	user.GET("/", handler.GetAll)
	user.GET("/:email", handler.Get)
	user.POST("/", handler.Create)
	user.PUT("/:email", handler.Update)
	user.DELETE("/:email", handler.Delete)
}

func (ush *UserHandler) GetAll(c *gin.Context) {
	res, err := ush.UserUsecase.GetAll(context.TODO())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (ush *UserHandler) Get(c *gin.Context) {
	res, err := ush.UserUsecase.Get(context.TODO(), c.Param("email"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (ush *UserHandler) Create(c *gin.Context) {
	user := &RepoDomain.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		RoleName: c.PostForm("rolename"),
	}
	if err := ush.UserUsecase.Create(context.TODO(), user); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Update(c *gin.Context) {
	email := c.Param("email")
	user := &RepoDomain.User{
		Name:     c.PostForm("name"),
		Email:    email,
		RoleName: c.Param("rolename"),
	}
	if err := ush.UserUsecase.Update(context.TODO(), email, user); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Delete(c *gin.Context) {
	if err := ush.UserUsecase.Delete(context.TODO(), c.Param("email")); err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}
