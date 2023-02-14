package main

import (
	"fmt"
	HttpDelivery "go-authorization/internal/delivery/http"

	OrmFactory "go-authorization/internal/repository/database/factory"
	UsecaseFactory "go-authorization/internal/usecase/factory"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	db, err := OrmFactory.InitDb(viper.GetString("DB_DSN"))
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
