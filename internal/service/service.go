package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"url_shortener/internal/repo"
	"url_shortener/internal/service/utils"
	"url_shortener/pkg/logging"
)

type UrlService interface {
	AddUrl(ctx context.Context, url string) (string, error)
	GetUrl(ctx context.Context, alias string) (string, error)
}

type Service struct {
	repo repo.Repository
	l    *logging.Logger
}

func NewService(repo repo.Repository, l logging.Logger) *Service {
	return &Service{repo: repo, l: &l}
}

func (s *Service) AddUrl(ctx context.Context, url string) (string, error) {
	if !utils.IsValid(url) {
		s.l.Errorf("url %s is invalid", url)
		return "", fmt.Errorf("url %s is invalid", url)
	}

	if s.repo.ContainsUrl(ctx, url) {
		s.l.Errorf("url %s already exist in storage", url)
		return "", fmt.Errorf("url %s already exist in storage", url)
	}

	rand.Seed(time.Now().UnixNano())
	alias := utils.GenerateAlias()

	s.l.Debug(alias)

	for s.repo.ContainsAlias(ctx, alias) {
		s.l.Infof("alias %s already exists in storage", alias)
		alias = utils.GenerateAlias()
	}

	err := s.repo.CreateUrl(ctx, url, alias)
	if err != nil {
		return "", err
	}

	return alias, nil
}

func (s *Service) GetUrl(ctx context.Context, alias string) (string, error) {
	url, err := s.repo.GetUrl(ctx, alias)
	if err != nil {
		return "", err
	}

	return url, nil
}
