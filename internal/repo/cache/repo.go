package cache

import (
	"context"
	"errors"
	"url_shortener/internal/storage/cache"
	"url_shortener/pkg/logging"
)

type Repository struct {
	c *cache.Cache
	l *logging.Logger
}

func NewRepository(cache *cache.Cache, l logging.Logger) *Repository {
	return &Repository{c: cache, l: &l}
}

func (r *Repository) GetUrl(ctx context.Context, alias string) (string, error) {
	r.c.Mu.RLock()
	defer r.c.Mu.RUnlock()

	url, ok := r.c.AliasMap[alias]
	if !ok {
		r.l.Errorf("url %s doesn't exist", url)
		return "", errors.New("url doesn't exist")
	}

	return url, nil
}

func (r *Repository) CreateUrl(ctx context.Context, url string, alias string) error {
	r.c.Mu.Lock()
	defer r.c.Mu.Unlock()

	r.c.UrlMap[url] = alias
	r.c.AliasMap[alias] = url
	r.l.Infof("insert url %s and alias %s to cache", url, alias)

	return nil
}

func (r *Repository) ContainsAlias(ctx context.Context, alias string) bool {
	r.c.Mu.RLock()
	defer r.c.Mu.RUnlock()

	if _, ok := r.c.AliasMap[alias]; ok {
		return true
	}

	return false
}

func (r *Repository) ContainsUrl(ctx context.Context, url string) (bool, string) {
	r.c.Mu.RLock()
	defer r.c.Mu.RUnlock()

	if alias, ok := r.c.UrlMap[url]; ok {
		return true, alias
	}

	return false, ""
}
