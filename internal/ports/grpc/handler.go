package grpc_

import (
	"context"
	protos "github.com/nesrayr/protos/gen"
	"url_shortener/internal/service"
	"url_shortener/pkg/logging"
)

type Handler struct {
	service service.UrlService
	l       *logging.Logger
	protos.UnimplementedUrlShortenerServer
}

func NewHandler(service service.UrlService, l logging.Logger) *Handler {
	return &Handler{service: service, l: &l}
}

func (h *Handler) ShortenUrl(ctx context.Context, request *protos.ShortenUrlRequest) (*protos.ShortenUrlResponse, error) {
	h.l.Info("grpc handler ShortenUrl")

	url := request.Url

	alias, err := h.service.AddUrl(ctx, url)
	if err != nil {
		h.l.Error(err)
		return &protos.ShortenUrlResponse{Alias: alias}, nil
	}

	return &protos.ShortenUrlResponse{Alias: alias}, nil
}

func (h *Handler) GetUrl(ctx context.Context, request *protos.GetUrlRequest) (*protos.GetUrlResponse, error) {
	h.l.Info("grpc handler GetUrl")

	alias := request.Alias

	url, err := h.service.GetUrl(ctx, alias)
	if err != nil {
		h.l.Error(err)
		return nil, err
	}

	return &protos.GetUrlResponse{Url: url}, nil
}
