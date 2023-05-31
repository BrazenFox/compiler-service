package handler

import (
	"github.com/BrazenFox/compiler-service/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) run(c *gin.Context) {
	var input entity.Program
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logrus.Info("input", input)
	result, err := h.services.ProgramHandler.HandleProgram(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
	})

}
