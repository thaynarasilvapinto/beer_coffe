package main

import (
	"database/sql"
	"fmt"
	"src/db"
	"src/monitoring"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	postgresConnect           *sql.DB               = nil
	newrelicAplicationConnect *newrelic.Application = nil
	beers                     []db.Beer
	cofeers                   []db.Coffee
)

func main() {

	postgresConnect, _ = db.PostgresConnect()
	newrelicAplicationConnect, _ = monitoring.NewRelicConnect()

	beers = db.SelectAllBeer(postgresConnect, newrelicAplicationConnect)

	fmt.Println(beers)

	cofeers = db.SelectAllCoffee(postgresConnect, newrelicAplicationConnect)

	fmt.Println(cofeers)
}
