package router

import (
	"fmt"
	"server2/cmd/handler"
	"server2/internal/middleware"

	"github.com/gin-gonic/gin"
)

// API generate api
func API(handler handler.Handler, serverPort string) {
	tokenController := handler.TokenController
	storageController := handler.StoragesController
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/storages/size", storageController.ListSizeOption)
	r.GET("/storages/area", storageController.ListAreaOption)

	officeAuthorized := r.Group("/")
	officeAuthorized.Use(middleware.AuthRequired())
	{
		officeAuthorized.GET("/token/check", tokenController.Check)
		officeAuthorized.GET("/storages/list", storageController.ListStorage)

	}

	if serverPort == "" {
		serverPort = "8080"
		fmt.Print("using default port : 8080")
	}
	r.Run(`:` + serverPort)
}
