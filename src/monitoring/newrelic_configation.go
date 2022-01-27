package monitoring

import "github.com/newrelic/go-agent/v3/newrelic"

func NewRelicConnect() (*newrelic.Application, error) {

	newRelicApplicationConnect, err := newrelic.NewApplication(
		newrelic.ConfigAppName("PostgreSQL App"),
		newrelic.ConfigLicense("NEW_RELIC_LICENSE_KEY"),
	)

	return newRelicApplicationConnect, err
}
