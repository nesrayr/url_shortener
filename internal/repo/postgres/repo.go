package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"url_shortener/pkg/logging"
)

type Repository struct {
	db *pgxpool.Pool
	l  *logging.Logger
}

func NewRepository(db *pgxpool.Pool, l logging.Logger) *Repository {
	return &Repository{db: db, l: &l}
}

func (r *Repository) AddUrl(ctx context.Context, url string, alias string) error {
	_, err := r.db.Exec(ctx, insertUrlQuery, alias, url)
	if err != nil {
		r.l.Error(err)
		return err
	}
	r.l.Infof("inserting url %s with alias %s in db", url, alias)

	return nil
}

func (r *Repository) GetUrl(ctx context.Context, alias string) (string, error) {
	var url string
	err := r.db.QueryRow(ctx, selectUrlByAliasQuery, alias).Scan(&url)
	if err != nil {
		r.l.Error(err)
		return "", err
	}
	r.l.Infof("selecting url %s by alias %s", url, alias)

	return url, nil
}

func (r *Repository) Contains(ctx context.Context, alias string) (bool, error) {
	rows, err := r.db.Query(ctx, selectAliasQuery, alias)
	if err != nil {
		r.l.Error(err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
