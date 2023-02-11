package main

import (
	Config "go-authorization/config"
	HttpDelivery "go-authorization/internal/delivery/http"
	Db "go-authorization/internal/repository/db"
	OrmFactory "go-authorization/internal/repository/orm/factory"
	UsecaseFactory "go-authorization/internal/usecase/factory"
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

	UseCaseRepository := UsecaseFactory.InitUsecaseRepository(OrmRepository)

	server := gin.Default()
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "ok",
		})
	})

	HttpDelivery.NewUserHandler(server, UseCaseRepository.User)
	HttpDelivery.NewRoleHandler(server, UseCaseRepository.Role)
	HttpDelivery.NewPermissionHandler(server, UseCaseRepository.Permission)
	HttpDelivery.NewResourceHandler(server, UseCaseRepository.Resource)
	HttpDelivery.NewCoreHandler(server, UseCaseRepository.Core)

	server.Run(":8080")
}
