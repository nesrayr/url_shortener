package postgres

const (
	insertUrlQuery        = `INSERT INTO urls(alias, url) VALUES ($1, $2)`
	selectUrlByAliasQuery = `SELECT url FROM urls WHERE alias=$1`
	selectAliasQuery      = `SELECT * FROM urls WHERE alias=$1`
)
