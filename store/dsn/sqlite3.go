package dsn

const sqlite3InMemory = ":memory:"

func Sqlite3(path string) string {
	return path
}

func Sqlite3InMemory() string {
	return sqlite3InMemory
}
