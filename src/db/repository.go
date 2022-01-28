package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func SelectAllBeer(postgresConnect *sql.DB, newRelicApplicationConnect *newrelic.Application) []Beer {

	newRelicApplicationConnect.WaitForConnection(5 * time.Second)
	txn := newRelicApplicationConnect.StartTransaction("postgresQuery")
	ctx := newrelic.NewContext(context.Background(), txn)

	rows, _ := postgresConnect.QueryContext(ctx, "SELECT * FROM beer")

	var beer Beer
	var beers []Beer

	for rows.Next() {
		rows.Scan(&beer.id, &beer.name, &beer.beer_type)
		beers = append(beers, beer)
	}

	txn.End()
	newRelicApplicationConnect.Shutdown(5 * time.Second)

	return beers

}

func SelectAllCoffee(postgresConnect *sql.DB, newRelicApplicationConnect *newrelic.Application) []Coffee {

	newRelicApplicationConnect.WaitForConnection(5 * time.Second)
	txn := newRelicApplicationConnect.StartTransaction("postgresQuery")
	ctx := newrelic.NewContext(context.Background(), txn)

	rows, _ := postgresConnect.QueryContext(ctx, "SELECT * FROM coffee")

	var coffee Coffee
	var coffeers []Coffee

	for rows.Next() {
		rows.Scan(&coffee.id, &coffee.name)
		coffeers = append(coffeers, coffee)
	}

	txn.End()
	newRelicApplicationConnect.Shutdown(5 * time.Second)

	return coffeers

}
