package handler

import (
	"backend-app/domain/movie"
	"github.com/gin-gonic/gin"
)

type HTTPController interface {
	Save(c *gin.Context)
	List(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	DeleteById(c *gin.Context)
}
func NewHTTPController(movieUseCase movie.UseCase) HTTPController {
	return &httpController{
		movieUseCase: movieUseCase,
	}
}

type httpController struct {
	movieUseCase movie.UseCase
}