package handler

import (
	"ContainerMonitor-API/internall/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	logrus.Debug("Initializing new handler with provided service")
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	logrus.Debug("Initializing routes")

	router := gin.New()
	api := router.Group("/api")
	{
		lists := api.Group("containers")
		{
			lists.POST("/", h.createContainer)
			lists.POST("/ping-result", h.savePingResult)
			lists.GET("/", h.getAllContainers)
			lists.GET("/:id", h.getContainerById)
		}
	}

	logrus.Info("Routes initialized successfully")

	return router
}
