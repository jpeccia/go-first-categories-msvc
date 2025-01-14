package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/jpeccia/go-first-categories-msvc/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var body createCategoryInput

	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase()

	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true})
}
