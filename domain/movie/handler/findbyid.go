package handler

import (
	"backend-app/http/response"
	"backend-app/domain/movie/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (controller *httpController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.FindById(c.Request.Context(), id)
	defaultMovie := model.Movie{}
	if result == defaultMovie {
		response.Notfound(c, http.StatusNotFound)
		return
	}

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}