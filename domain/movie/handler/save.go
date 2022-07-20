package handler

import (
	"backend-app/domain/movie/model"
	"backend-app/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *httpController) Save(c *gin.Context) {
	var payload model.Movie
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}
	result, err := controller.movieUseCase.Save(c.Request.Context(), payload)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, result)
}