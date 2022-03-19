package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckPerformanceController interface {
	LoadPosts(context *gin.Context)
}

type checkPerformanceController struct {
}

func (c *checkPerformanceController) LoadPosts(context *gin.Context) {
	context.JSON(http.StatusOK, "ok")
}

func NewCheckPerformanceController() CheckPerformanceController {
	return &checkPerformanceController{}
}
