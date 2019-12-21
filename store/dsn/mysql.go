package dsn

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func Mysql(host string, port int, user, password, database string, args ...string) string {
	cfg := mysql.NewConfig()
	cfg.Addr = host
	if port > 0 {
		cfg.Addr = fmt.Sprintf("%s:%d", host, port)
	}

	cfg.Net = "tcp"
	cfg.User = user
	cfg.Passwd = password
	cfg.DBName = database
	cfg.ParseTime = true

	// https://github.com/go-sql-driver/mysql#charset
	// https://github.com/go-sql-driver/mysql#collation
	cfg.Collation = "utf8mb4_general_ci"

	cfg.Params = make(map[string]string, len(args)/2)
	for idx := 0; idx < len(args)-1; idx += 2 {
		cfg.Params[args[idx]] = args[idx+1]
	}

	return cfg.FormatDSN()
}
