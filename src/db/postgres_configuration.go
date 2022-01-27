package db

import (
	"database/sql"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
)

var postgresConnect *sql.DB = nil

func PostgresConnect() (*sql.DB, error) {

	postgresConnect, err := sql.Open("nrpgx", "host=localhost port=5432 user=beer-coffee dbname=beer-coffee password=12345678 sslmode=disable")

	return postgresConnect, err

}
