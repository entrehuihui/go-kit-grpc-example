//go:generate go-bindata -prefix ./migrations/ -pkg orderpostgres -o ./migrations_gen.go ./migrations/
package orderpostgres

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// MigrateDB ..
func MigrateDB(db *sqlx.DB) error {
	m := &migrate.AssetMigrationSource{
		Asset:    Asset,
		Dir:      "",
		AssetDir: AssetDir,
	}

	_, err := migrate.Exec(db.DB, "postgres", m, migrate.Up)
	if err != nil {
		return err
	}

	return nil
}
