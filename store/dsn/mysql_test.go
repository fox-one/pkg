package dsn

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysql(t *testing.T) {
	dsn := Mysql("localhost", 13306, "root", "yiplee", "fox-api")
	t.Log(dsn)

	db, err := sql.Open("mysql", dsn)
	if assert.Nil(t, err) {
		assert.Nil(t, db.Ping())
	}
}
