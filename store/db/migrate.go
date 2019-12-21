package db

import (
	"github.com/hashicorp/go-multierror"
)

type MigrateFunc func(*DB) error

var migrateFuncs []MigrateFunc

func RegisterMigrate(fn MigrateFunc) {
	migrateFuncs = append(migrateFuncs, fn)
}

func Migrate(db *DB) error {
	db = &DB{
		read:  db.read,
		write: db.write.Set("gorm:table_options", "CHARSET=utf8mb4"),
	}

	var err *multierror.Error

	for _, fn := range migrateFuncs {
		err = multierror.Append(err, fn(db))
	}

	return err.ErrorOrNil()
}
