package db

import (
	"fmt"

	"github.com/fox-one/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect(dialect, uri string) (*DB, error) {
	coon, err := gorm.Open(dialect, uri)
	if err != nil {
		return nil, err
	}

	return &DB{
		write: coon,
		read:  coon,
	}, nil
}

func open(dialect, dsn string) (*gorm.DB, error) {
	if dialect == "sqlite" {
		dialect = "sqlite3"
	}

	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Open(cfg db.Config) (*DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}

	write, err := open(cfg.Dialect, dsn)
	if err != nil {
		return nil, err
	}

	db := &DB{
		write: write,
		read:  write,
	}

	if cfg.ReadHost != "" && cfg.ReadHost != cfg.Host {
		readHostDSN, err := cfg.ReadHostDSN()
		if err != nil {
			return nil, err
		}
		db.read, err = open(cfg.Dialect, readHostDSN)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Debug {
		db = db.Debug()
	}

	return db, nil
}

func MustOpen(cfg db.Config) *DB {
	db, err := Open(cfg)
	if err != nil {
		panic(fmt.Errorf("open db failed: %w", err))
	}

	return db
}
