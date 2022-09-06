package store2

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Update() *gorm.DB {
	return db.Clauses(dbresolver.Write)
}

func (db *DB) View() *gorm.DB {
	return db.Clauses(dbresolver.Read)
}
