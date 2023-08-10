module github.com/fox-one/pkg/store2

go 1.18

replace github.com/fox-one/pkg/db => ../db

replace github.com/fox-one/pkg/property => ../property

require (
	github.com/fox-one/pkg/db v0.0.0-20230711064542-e002c9aad80a
	github.com/fox-one/pkg/property v0.0.2
	github.com/hashicorp/go-multierror v1.1.1
	gorm.io/driver/mysql v1.3.6
	gorm.io/driver/postgres v1.3.9
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
	gorm.io/plugin/dbresolver v1.2.3
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/yiplee/go-cache v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/text v0.3.7 // indirect
)
