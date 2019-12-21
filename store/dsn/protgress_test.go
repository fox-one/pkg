package dsn

import (
	"testing"
)

func TestPostgres(t *testing.T) {
	d := Postgres("localhost", 0, "user", "pwd", "db", "a", "b")
	t.Log(d)
}
