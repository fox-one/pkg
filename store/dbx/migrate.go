package dbx

import (
	"net/url"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
)

// https://github.com/golang-migrate/migrate
// dir is the ddl sql folder path
func (db *DB) Migrate(dir string) (*migrate.Migrate, error) {
	var driver database.Driver

	switch db.Driver() {
	case Mysql:
		driver, _ = mysql.WithInstance(db.conn.DB, &mysql.Config{})
	case Postgres:
		driver, _ = postgres.WithInstance(db.conn.DB, &postgres.Config{})
	default:
		driver, _ = sqlite3.WithInstance(db.conn.DB, &sqlite3.Config{})
	}

	u := url.URL{Scheme: "file", Path: dir}
	source := u.String()
	return migrate.NewWithDatabaseInstance(source, string(db.Driver()), driver)
}

func (db *DB) MigrateUp(dir string) error {
	m, err := db.Migrate(dir)
	if err != nil {
		return err
	}

	return m.Up()
}
