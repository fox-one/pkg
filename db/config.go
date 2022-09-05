package db

type Config struct {
	Dialect  string `json:"dialect"` // mysql,postgres,sqlite3
	Host     string `json:"host"`    // if Dialect is `sqlite3`, host should be db file path
	ReadHost string `json:"read_host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Location string `json:"location,omitempty"`
	SSLMode  string `json:"sslmode,omitempty"` // postgres has this property, which could be "disable"
	Debug    bool   `json:"debug,omitempty"`
}

func SqliteInMemory() Config {
	return Config{
		Dialect: "sqlite3",
		Host:    ":memory:",
	}
}
