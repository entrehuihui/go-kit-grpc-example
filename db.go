package ordering

import (
	"./orderpostgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //
)

// Connect connect postgresql
func Connect(dsn string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = orderpostgres.MigrateDB(db)
	if err != nil {
		return nil, err
	}
	return
}
