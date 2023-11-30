package repo

import "context"

type Repository interface {
	GetUrl(ctx context.Context, alias string) (string, error)
	CreateUrl(ctx context.Context, url string, alias string) error
	ContainsAlias(ctx context.Context, alias string) bool
	ContainsUrl(ctx context.Context, url string) bool
}
