package prom

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// mttpRouteMetricer is an interface that is implmeneted by the mttpRouteMetrics struct
// it provides a means to create and increment prometheus counters for a given server, and
// route identifier
type MttpRouteMetricer interface {
	IncremementStatusCode(statusCode int)
}

type mttpRouteMetrics struct {
	counters map[int]prometheus.Counter
}

func NewMttpRouteMetrics(serverName, routeIdentifier string) *mttpRouteMetrics {
	return &mttpRouteMetrics{
		counters: createDefaultPrometheusCounters(serverName, routeIdentifier),
	}
}

// IncremementStatusCode incremements a route's counter for a status code.
func (mrm *mttpRouteMetrics) IncremementStatusCode(statusCode int) {
	mrm.counters[mrm.getStatusFamily(statusCode)].Inc()
}

// getStatusFamily is a helper function that returns a given status code's "family"
// for example; providing 404 will result in 400
func (mrm *mttpRouteMetrics) getStatusFamily(statusCode int) int {
	return statusCode - statusCode%100
}

// createDefaultPrometheusCounters creates a default set of prometheus metrics that mttp provides for each of it's routes
func createDefaultPrometheusCounters(serverName, routeId string) map[int]prometheus.Counter {
	return map[int]prometheus.Counter{
		200: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_2XX_total", serverName, routeId),
			Help: fmt.Sprintf("The total number HTTP 2XX status code returned from the %s route", routeId),
		}),
		400: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_4XX_total", serverName, routeId),
			Help: fmt.Sprintf("The total number HTTP 4XX status code returned from the %s route", routeId),
		}),
		500: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_5XX_total", serverName, routeId),
			Help: fmt.Sprintf("The total number HTTP 5XX status code returned from the %s route", routeId),
		}),
	}
}
