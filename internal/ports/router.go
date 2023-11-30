package ports

import (
	"github.com/gin-gonic/gin"
	"url_shortener/internal/service"
	"url_shortener/pkg/logging"
)

func SetupRouter(service service.UrlService, logger logging.Logger) *gin.Engine {
	router := gin.Default()

	handler := NewHandler(service, logger)

	router.POST("/api/shorten", handler.AddUrl)
	router.GET("/api/get-url/:alias", handler.GetUrl)

	return router
}
