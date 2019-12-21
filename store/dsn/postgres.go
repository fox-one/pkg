package dsn

import (
	"bytes"
	"strconv"
)

func Postgres(host string, port int, user, password, database string, args ...string) string {
	var buf bytes.Buffer

	if host != "" {
		buf.WriteString("host=")
		buf.WriteString(host)
		buf.WriteByte(' ')
	}

	if port > 0 {
		buf.WriteString("port=")
		buf.WriteString(strconv.Itoa(port))
		buf.WriteByte(' ')
	}

	if user != "" {
		buf.WriteString("user=")
		buf.WriteString(user)
		buf.WriteByte(' ')
	}

	if password != "" {
		buf.WriteString("password=")
		buf.WriteString(password)
		buf.WriteByte(' ')
	}

	if database != "" {
		buf.WriteString("dbname=")
		buf.WriteString(database)
		buf.WriteByte(' ')
	}

	for idx := 0; idx < len(args)-1; idx += 2 {
		buf.WriteString(args[idx])
		buf.WriteByte('=')
		buf.WriteString(args[idx+1])
		buf.WriteByte(' ')
	}

	return buf.String()
}
