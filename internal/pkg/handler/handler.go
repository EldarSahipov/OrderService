package handler

import (
	_ "OrderService/docs"
	"OrderService/internal/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() (router *gin.Engine) {
	router = gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("api")
	{
		orders := api.Group("/orders")
		{
			orders.GET("/:uid", h.GetOrderByUIDFromCache)
			orders.GET("/all", h.getAll)
		}
	}

	return
}
