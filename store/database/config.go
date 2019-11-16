package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Config struct {
	Dialect  string `json:"dialect,omitempty"` // mysql,postgres,sqlite3
	Host     string `json:"host,omitempty"`    // if Dialect is `sqlite3`, host should be db file path
	ReadHost string `json:"read_host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
	Debug    bool   `json:"debug,omitempty"`
}

func SqliteInMemory() Config {
	return Config{
		Dialect: "sqlite3",
		Host:    ":memory:",
	}
}

func open(dialect, host string, port int, user, password, database string) (*gorm.DB, error) {
	var uri string
	switch dialect {
	case "mysql":
		uri = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&charset=utf8mb4",
			user,
			password,
			"tcp",
			host,
			port,
			database,
		)
	case "postgres":
		uri = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			host,
			port,
			user,
			database,
			password,
		)
	case "sqlite3", "sqlite":
		dialect = "sqlite3"
		uri = host
	default:
		return nil, fmt.Errorf("unkonow db dialect: %s", dialect)
	}

	db, err := gorm.Open(dialect, uri)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	return db, nil
}

func Open(cfg Config) (*DB, error) {
	write, err := open(cfg.Dialect, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	if err != nil {
		return nil, err
	}

	db := &DB{
		write: write,
		read:  write,
	}

	if cfg.ReadHost != "" && cfg.ReadHost != cfg.Host {
		db.read, err = open(cfg.Dialect, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Debug {
		db = db.Debug()
	}

	return db, nil
}

func MustOpen(cfg Config) *DB {
	db, err := Open(cfg)
	if err != nil {
		panic(fmt.Errorf("open db failed: %w", err))
	}

	return db
}
