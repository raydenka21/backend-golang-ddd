package handler

import (
	"backend-app/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *httpController) List(c *gin.Context) {
	result, err := controller.movieUseCase.List(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}
	response.Success(c, http.StatusOK, result)
}
