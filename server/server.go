package server

import (
	"backend-app/config"
	"backend-app/domain/movie"
	"backend-app/domain/movie/handler"
	"backend-app/domain/movie/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterAPIService(e *gin.Engine) {
	db = config.GetDBConnection()

	registerMovieAPIService(e)
}

func registerMovieAPIService(r *gin.Engine) {
	// Initialize Movie Service
	movieRepo := postgres.NewRepository(db)
	movieUseCase := movie.NewUseCase(movieRepo)
	movieController := handler.NewHTTPController(movieUseCase)
	// Start API
	registerMovieRoute(r, movieController)
}
