package handler

import (
	"backend-app/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func (controller *httpController) DeleteById(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	err = controller.movieUseCase.DeleteById(c.Request.Context(), reqID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, nil)
}