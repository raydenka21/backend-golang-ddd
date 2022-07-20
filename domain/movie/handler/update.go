package handler

import (
	"backend-app/domain/movie/model"
	"backend-app/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (controller *httpController) Update(c *gin.Context) {
	var spec model.Movie
	var err error
	err = c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	spec.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.Update(c.Request.Context(), spec)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}