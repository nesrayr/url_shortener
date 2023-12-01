package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shortener/internal/service"
	"url_shortener/pkg/logging"
)

type AddUrlRequest struct {
	Url string `json:"url"`
}

type Handler struct {
	service service.UrlService
	l       *logging.Logger
}

func NewHandler(service service.UrlService, l logging.Logger) *Handler {
	return &Handler{service: service, l: &l}
}

func (h *Handler) ShortenUrl(ctx *gin.Context) {
	var request AddUrlRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.l.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alias, err := h.service.AddUrl(ctx, request.Url)
	if errors.Is(err, service.ErrUrlAlreadyExists) {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error(),
			"alias": alias})
		return
	} else if err != nil {
		h.l.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"alias": alias})
}

func (h *Handler) GetUrl(ctx *gin.Context) {
	alias := ctx.Param("alias")

	url, err := h.service.GetUrl(ctx, alias)
	if err != nil {
		h.l.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"url": url})
}
