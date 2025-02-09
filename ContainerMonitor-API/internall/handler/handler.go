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

	router.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Debug("CORS Middleware triggered")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		logrus.Debug("Responding with CORS headers")

		c.Next()
	}
}
