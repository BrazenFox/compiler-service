package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		program := api.Group("/program")
		{
			program.POST("/run")
		}
	}

	return router
}