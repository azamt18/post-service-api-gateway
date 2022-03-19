package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Cors() func(ginContext *gin.Context) {
	corsConfig := cors.Config{}

	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	corsConfig.MaxAge = 24 * time.Hour

	return cors.New(corsConfig)
}
