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
		articles.GET("/create", h.create)
		articles.GET("/error", h.error)
		articles.POST("/saveArticle", h.saveArticle)
		articles.GET("/getAll", h.getAll)
		articles.GET("/getByID/:id", h.getByID)
		articles.GET("/delete/:id", h.delete)
	}

	return router
}
