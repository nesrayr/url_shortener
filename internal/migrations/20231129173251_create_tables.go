package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTables, downCreateTables)
}

func upCreateTables(tx *sql.Tx) error {
	//_, err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	//if err != nil {
	//	return err
	//}

	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS urls(
    alias varchar(10) primary key , 
	url varchar(1024) not null unique 
)`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateTables(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE urls`)
	if err != nil {
		return err
	}

	return nil
}
