package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/go-first-categories-msvc/cmd/api/controllers"
	"github.com/jpeccia/go-first-categories-msvc/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	InMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx,InMemoryCategoryRepository )
	})
}
