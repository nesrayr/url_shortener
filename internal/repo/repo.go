package repo

import "context"

type Repository interface {
	GetUrl(ctx context.Context, alias string) (string, error)
	AddUrl(ctx context.Context, url string, alias string) error
	Contains(ctx context.Context, alias string) (bool, error)
}
