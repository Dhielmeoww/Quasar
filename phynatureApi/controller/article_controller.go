package controller

import (
	"net/http"
	"phynatureApi/entity"
	"phynatureApi/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

func (c *ArticleController) GetAllArticle(ctx *gin.Context) {
	article, err := c.articleService.GetAllArticle()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": article})
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var article entity.Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"data": article})
}
