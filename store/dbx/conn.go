// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package dbx

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Connect to a database and verify with a ping.
func Connect(driver, dataSource string) (*DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	switch driver {
	case "mysql":
		// default
		db.SetMaxIdleConns(2)
	}
	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	var engine string
	var locker Locker
	switch driver {
	case Mysql:
		engine = Mysql
		locker = &nopLocker{}
	case Postgres:
		engine = Postgres
		locker = &nopLocker{}
	default:
		engine = Sqlite
		locker = &sync.RWMutex{}
	}

	return &DB{
		conn:   sqlx.NewDb(db, driver),
		lock:   locker,
		driver: engine,
	}, nil
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}
