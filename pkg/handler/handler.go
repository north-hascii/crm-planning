package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/north-hascii/crm-planning/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		authenticated := api.Group("/verified", h.userIdentity)
		{
			authenticated.GET("/", h.getUser)

			operations := api.Group("/operations", h.userIdentity)
			{
				operations.GET("/get", h.getOperationById)
				operations.GET("/get-all", h.getAllOperation)
				operations.POST("/create", h.createOperation)
				operations.PUT("/update", h.updateOperation)
				operations.DELETE("/delete", h.deleteOperation)
			}

		}
	}

	return router
}
