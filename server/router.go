package server

import (
	"backend-app/domain/movie/handler"
	"github.com/gin-gonic/gin"
)

func registerMovieRoute(r *gin.Engine, movieController handler.HTTPController) {
	movieRouter := r.Group("/v1/movies")
	movieRouter.GET("/", movieController.List)
	movieRouter.POST("/", movieController.Save)
	movieRouter.GET("/:id", movieController.FindById)
	movieRouter.PUT("/:id", movieController.Update)
	movieRouter.DELETE("/:id", movieController.DeleteById)

}
