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

type UserParams struct {
	Email    string `url:"email"`
	RoleName string `url:"rolename"`
}

func NewUserHandler(server *gin.Engine, usecase UsecaseDomain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: usecase,
	}

	user := server.Group("/users")
	user.GET("/", handler.GetAll)
	user.GET("/:email/:rolename", handler.Get)
	user.POST("/", handler.Create)
	user.PUT("/", handler.Update)
	user.DELETE("/:email", handler.Delete)
}

func (ush *UserHandler) GetAll(c *gin.Context) {
	res, err := ush.UserUsecase.GetAll(context.TODO())
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, res)
}

func (ush *UserHandler) Get(c *gin.Context) {
	var userParams UserParams
	if err := c.ShouldBindUri(&userParams); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if userParams.Email != "" && userParams.RoleName != "" {
		c.JSON(400, gin.H{"msg": "must use only one query parameter"})
		return
	} else if userParams.Email != "" {
		res, err := ush.UserUsecase.Get(context.TODO(), userParams.Email)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, res)
	} else if userParams.RoleName != "" {
		res, err := ush.UserUsecase.GetByRole(context.TODO(), userParams.RoleName)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, res)
	} else {
		c.JSON(400, gin.H{"msg": "query parameter is empty"})
	}
}

func (ush *UserHandler) Create(c *gin.Context) {
	rolename := c.PostForm("rolename")
	user := &RepoDomain.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		RoleName: &rolename,
	}
	if err := ush.UserUsecase.Create(context.TODO(), user); err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Update(c *gin.Context) {
	email := c.PostForm("email")
	rolename := c.PostForm("rolename")
	user := &RepoDomain.User{
		Name:     c.PostForm("name"),
		Email:    email,
		RoleName: &rolename,
	}
	if err := ush.UserUsecase.Update(context.TODO(), email, user); err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Delete(c *gin.Context) {
	if err := ush.UserUsecase.Delete(context.TODO(), c.PostForm("email")); err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, "success")
}
