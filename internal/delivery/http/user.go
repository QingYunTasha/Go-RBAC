package http

import (
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
	res, err := ush.UserUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (ush *UserHandler) Get(c *gin.Context) {
	res, err := ush.UserUsecase.Get(c.Request.Context(), c.Param("email"))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

func (ush *UserHandler) Create(c *gin.Context) {
	user := RepoDomain.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := ush.UserUsecase.Create(c.Request.Context(), &user); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Update(c *gin.Context) {
	user := RepoDomain.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := ush.UserUsecase.Update(c.Request.Context(), c.Param("email"), &user); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}

func (ush *UserHandler) Delete(c *gin.Context) {
	if err := ush.UserUsecase.Delete(c.Request.Context(), c.Param("email")); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "success")
}
