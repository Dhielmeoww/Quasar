package controller

import (
	"net/http"
	"phynatureApi/entity"
	"phynatureApi/service"

	"github.com/gin-gonic/gin"
)

type KategoriController struct {
	kategoriService service.KategoriService
}

func NewKategoriController(kategoriService service.KategoriService) *KategoriController {
	return &KategoriController{kategoriService: kategoriService}
}

func (c *KategoriController) GetAllKategori(ctx *gin.Context) {
	article, err := c.kategoriService.GetAllKategori()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": article})
}

func (c *KategoriController) CreateKategori(ctx *gin.Context) {
	var kategori entity.Kategori
	err := ctx.ShouldBindJSON(&kategori)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	err = c.kategoriService.CreateKategori(&kategori)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": kategori})
}
