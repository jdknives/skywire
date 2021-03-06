package setupmetrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/skycoin/skywire/pkg/routing"
)

// NewEmpty creates a new metrics implementation that does nothing.
func NewEmpty() Metrics {
	return empty{}
}

type empty struct{}

func (empty) Collectors() []prometheus.Collector {
	return nil
}
func (empty) RecordRequest(routing.BidirectionalRoute) func(*routing.EdgeRules, *error) {
	return func(*routing.EdgeRules, *error) {}
}
