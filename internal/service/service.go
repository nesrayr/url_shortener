package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"url_shortener/internal/repo"
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
	if !IsValid(url) {
		s.l.Errorf("url %s is invalid", url)
		return "", fmt.Errorf("url %s is invalid", url)
	}

	check, alias := s.repo.ContainsUrl(ctx, url)

	if check {
		s.l.Infof("url %s already exist in storage", url)
		return alias, ErrUrlAlreadyExists
	}

	rand.Seed(time.Now().UnixNano())
	alias = GenerateAlias()

	s.l.Debug(alias)

	if s.repo.ContainsAlias(ctx, alias) {
		s.l.Infof("alias %s already exists in storage", alias)
		alias = GenerateAlias()
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
