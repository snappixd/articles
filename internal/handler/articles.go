package handler

import (
	"articles_psql/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) create(c *gin.Context) {
	c.HTML(http.StatusOK, "article_create.html", gin.H{})
}

func (h *Handler) saveArticle(c *gin.Context) {
	var article models.Article

	article.Author, _ = c.GetPostForm("author")
	article.Title, _ = c.GetPostForm("title")
	article.Anons, _ = c.GetPostForm("anons")
	article.Text, _ = c.GetPostForm("text")

	if err := h.services.Articles.Create(c.Request.Context(), article); err != nil {
		log.Println(err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/articles/getAll")
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

	c.HTML(http.StatusOK, "article_show.html", gin.H{
		"title":   article.Title,
		"article": article,
	})
}

func (h *Handler) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	err = h.services.Articles.Delete(c.Request.Context(), id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c.Redirect(http.StatusSeeOther, "/articles/getAll")
}
