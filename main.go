package main

import (
	"github.com/gin-gonic/gin"
	"mobile_storage_test_service/Logger"
	"mobile_storage_test_service/Services"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	//router.Use(Logger.LoggerMiddleware())
	router.Use(Logger.LogrusLogger())
	v1 := router.Group("/")
	{
		v1.POST("CreateApplicationForm", Services.CreateApplicationForm)
		v1.POST("BatchGetApplicationForm", Services.BatchGetApplicationForm)
		v1.POST("UpdateApplicationFormTestV1", Services.UpdateApplicationFormTestV1)
	}
	router.Run(":8080")
}
