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

func (r *Repository) CreateUrl(ctx context.Context, url string, alias string) error {
	_, err := r.db.Exec(ctx, insertUrlQuery, alias, url)
	if err != nil {
		r.l.Error(err)
		return err
	}
	r.l.Infof("inserting url %s with alias %s to db", url, alias)

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

func (r *Repository) ContainsAlias(ctx context.Context, alias string) bool {
	rows, err := r.db.Query(ctx, checkAliasQuery, alias)
	if err != nil {
		r.l.Error(err)
		return false
	}
	if rows.Next() {
		return true
	}
	return false
}

func (r *Repository) ContainsUrl(ctx context.Context, url string) (bool, string) {
	rows, err := r.db.Query(ctx, checkUrlQuery, url)
	if err != nil {
		r.l.Error(err)
		return false, ""
	}
	var alias string
	if rows.Next() {
		err = rows.Scan(&alias)
		if err != nil {
			r.l.Error(err)
		}
		return true, alias
	}
	return false, ""
}
