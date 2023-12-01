package grpc_

import (
	"context"
	protos "github.com/nesrayr/protos/gen"
	"google.golang.org/grpc/peer"
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
	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()
	h.l.Debugf("grpc handling on %s", addr)

	url := request.Url

	alias, err := h.service.AddUrl(ctx, url)
	if err != nil {
		h.l.Error(err)
		return &protos.ShortenUrlResponse{Alias: alias}, nil
	}

	h.l.Infof("grpc responded on %s", addr)

	return &protos.ShortenUrlResponse{Alias: alias}, nil
}

func (h *Handler) GetUrl(ctx context.Context, request *protos.GetUrlRequest) (*protos.GetUrlResponse, error) {
	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()
	h.l.Debugf("grpc handling on %s", addr)

	alias := request.Alias

	url, err := h.service.GetUrl(ctx, alias)
	if err != nil {
		h.l.Error(err)
		return nil, err
	}

	h.l.Info("grpc responded on %s", addr)

	return &protos.GetUrlResponse{Url: url}, nil
}
