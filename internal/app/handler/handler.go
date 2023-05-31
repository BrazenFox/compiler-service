package handler

import (
	"github.com/BrazenFox/compiler-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		program := api.Group("/program")
		{
			program.POST("/run", h.run)
		}
	}

	return router
}
