package handler

import (
	"articles_psql/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) create(c *gin.Context) {
	var article models.Article

	if err := c.BindJSON(&article); err != nil {
		log.Println(err.Error())
		return
	}

	if err := h.services.Articles.Create(c.Request.Context(), article); err != nil {
		log.Println(err.Error())
		return
	}

	c.Redirect(http.StatusCreated, "/articles/getAll")
}

func (h *Handler) getAll(c *gin.Context) {
	articles, err := h.services.Articles.GetAll(c.Request.Context())
	if err != nil {
		log.Println(err.Error())
	}

	//c.JSON(http.StatusOK, articles)
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":    "Articles",
		"articles": articles,
	})

}

func (h *Handler) getByID(c *gin.Context) {
	//id := c.GetInt("id")
	id, _ := strconv.Atoi(c.Param("id"))
	log.Println(id)

	article, err := h.services.Articles.GetByID(c.Request.Context(), id)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/articles/getAll")
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) delete(c *gin.Context) {
	id := c.GetInt("id")

	err := h.services.Articles.Delete(c.Request.Context(), id)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
