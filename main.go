package main

import (
	"github.com/gin-gonic/gin"
	"mobile_storage_test_service/Services"
	"mobile_storage_test_service/logger"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.Use(logger.LoggerMiddleware())

	v1 := router.Group("/")
	{
		v1.POST("AcceptApplicationForm", Services.AcceptApplicationForm)
	}
	router.Run(":8080")
}
