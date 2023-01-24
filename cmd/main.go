package main

import (
	Config "go-authorization/config"
	OrmFactory "go-authorization/internal/repository/orm/factory"
	Db "go-authorization/model/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	db, err := Db.InitDb(Config.DB_DSN)
	if err != nil {
		panic("database connect fail")
	}
	OrmRepository, err := OrmFactory.InitOrmRepository(db)
	if err != nil {
		panic("orm init fail")
	}

	ApiRepository, err := ApiFactory.InitApi(OrmRepository)
	resourceApi := Api.NewResourceApi(OrmRepository)

	server := gin.Default()
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "ok",
		})
	})

	resource := server.Group("/resource")
	{
		resource.GET("/", resourceApi.GetAll)
		resource.GET("/:name", resourceApi.Get)
		resource.POST("/", resourceApi.Create)
		resource.DELETE("/:name", resourceApi.Delete)
		resource.PUT("/:title", resourceApi.Update)
	}

	server.Run(":8080")
}
