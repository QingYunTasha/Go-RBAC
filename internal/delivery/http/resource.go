package http

import (
	RepoDomain "go-authorization/domain/repository"
	UsecaseDomain "go-authorization/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceHandler struct {
	ResourceUsecase UsecaseDomain.ResourceUsecase
}

func NewResourceHandler(server *gin.Engine, usecase UsecaseDomain.ResourceUsecase) {
	handler := &ResourceHandler{
		ResourceUsecase: usecase,
	}

	res := server.Group("/resources")
	res.GET("/", handler.GetAll)
	res.GET("/:name", handler.Get)
	res.POST("/", handler.Create)
	res.PUT("/:name", handler.Update)
	res.DELETE("/:name", handler.Delete)
}

func (rsh *ResourceHandler) GetAll(c *gin.Context) {
	res, err := rsh.ResourceUsecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rsh *ResourceHandler) Get(c *gin.Context) {
	resource, err := rsh.ResourceUsecase.Get(c.Request.Context(), c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, resource)
}

func (rsh *ResourceHandler) Create(c *gin.Context) {
	resource := RepoDomain.Resource{}
	if err := c.BindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rsh.ResourceUsecase.Create(c.Request.Context(), &resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (rsh *ResourceHandler) Update(c *gin.Context) {
	resource := RepoDomain.Resource{}
	if err := c.BindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := rsh.ResourceUsecase.Update(c.Request.Context(), c.Param("name"), &resource); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (rsh *ResourceHandler) Delete(c *gin.Context) {
	if err := rsh.ResourceUsecase.Delete(c.Request.Context(), c.Param("name")); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}
