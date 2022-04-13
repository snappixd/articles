package handler

import (
	"articles_psql/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("ui/html/*.html")

	router.GET("/", h.home)

	articles := router.Group("/articles")
	{
		articles.POST("/create", h.create)
		articles.GET("/getAll", h.getAll)
		articles.GET("/getByID/:id", h.getByID)
		articles.DELETE("/delete/:id", h.delete)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.sign_up)
		auth.POST("/sign-in", h.sign_in)
	}

	return router
}
