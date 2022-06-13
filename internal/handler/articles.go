package handler

import (
	"articles_psql/internal/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Article
// @Tags articles
// @Description To Create Articles
// @ID create-article
// @Accept json
// @Produce json
// @Param input body models.Article.Author true models.Article.Title true models.Article.Anons true models.Article.Text  true
// @Success 200
// @Failure 400
// @Router /articles/create [get]

// @Summary Create article
// @Tags articles
// @Description create article pls
// @ID create-article
// @Accept  json
// @Produce  json
// @Param input body models.Article true "list info"
// @Success 200
// @Failure default
// @Router /articles/create [get]

func (h *Handler) create(c *gin.Context) {
	c.HTML(http.StatusOK, "article_create.html", gin.H{
		"title": "Create",
	})
}

func (h *Handler) error(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{})
}

func (h *Handler) saveArticle(c *gin.Context) {
	var article models.Article

	article.Author, _ = c.GetPostForm("author")
	article.Title, _ = c.GetPostForm("title")
	article.Anons, _ = c.GetPostForm("anons")
	article.Text, _ = c.GetPostForm("text")
	//photo, _ := c.GetPostForm("photo")
	//photo, _, _ := c.Request.FormFile("upload")

	// fInfo, _ := photo.Stat()
	// var size int64 = fInfo.Size()
	// buf := make([]byte, size)

	// fReader := bufio.NewReader(photo)
	// fReader.Read(buf)

	// convertedPhoto := base64.StdEncoding.EncodeToString(buf)
	// article.Photo = convertedPhoto

	photo, header, _ := c.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./ui/img/" + filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, photo)
	if err != nil {
		log.Fatal(err)
	}

	article.Photo = filename

	if article.Author == "" || article.Title == "" || article.Anons == "" || article.Text == "" {
		c.Redirect(http.StatusSeeOther, "/articles/error")
	} else {
		if err := h.services.Articles.Create(c.Request.Context(), article); err != nil {
			log.Println(err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/articles/getAll")
	}
}

func (h *Handler) getAll(c *gin.Context) {
	articles, err := h.services.Articles.GetAll(c.Request.Context())
	if err != nil {
		log.Println(err.Error())
	}

	//c.JSON(http.StatusOK, articles)
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":    "Чудовий сайт",
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
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.services.Articles.Delete(c.Request.Context(), id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	c.Redirect(http.StatusSeeOther, "/articles/getAll")
}
