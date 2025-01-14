package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-first-categories-msvc/cmd/api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
}
