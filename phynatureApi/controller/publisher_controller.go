package controller

import (
	"net/http"
	"phynatureApi/entity"
	"phynatureApi/service"

	"github.com/gin-gonic/gin"
)

type PublisherController struct {
	publisherService service.PublisherService
}

func NewPublisherController(publisherService service.PublisherService) *PublisherController {
	return &PublisherController{publisherService: publisherService}
}

func (c *PublisherController) GetAllPublisher(ctx *gin.Context) {
	publisher, err := c.publisherService.GetAllPublisher()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": publisher})
}

func (c *PublisherController) CreatePublisher(ctx *gin.Context) {
	var publisher entity.Publisher
	err := ctx.ShouldBindJSON(&publisher)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	err = c.publisherService.CreatePublisher(&publisher)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": publisher})
}
