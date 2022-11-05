package mttp

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	twoHundred  = "2XX"
	fourHundred = "4XX"
	fiveHundred = "5XX"
)

func registerRouteWithPrometheusCounters(serverName, canonicalRoute string) map[string]prometheus.Counter {

	canonicalServerName := strings.ReplaceAll(serverName, "-", "_")
	return map[string]prometheus.Counter{
		"2XX": promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_2XX_total", canonicalServerName, canonicalRoute),
			Help: fmt.Sprintf("The total number HTTP 2XX status code returned from the %s route", canonicalRoute),
		}),
		"4XX": promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_4XX_total", canonicalServerName, canonicalRoute),
			Help: fmt.Sprintf("The total number HTTP 4XX status code returned from the %s route", canonicalRoute),
		}),
		"5XX": promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s_HTTP_5XX_total", canonicalServerName, canonicalRoute),
			Help: fmt.Sprintf("The total number HTTP 5XX status code returned from the %s route", canonicalRoute),
		}),
	}
}

func incrementStatusCounter(rw *mttpResponseWriter, counterMap map[string]prometheus.Counter) {
	counterMap[matchStatus(rw.statusCode)].Inc()
}

func matchStatus(statusCode int) string {
	switch {
	case isBetween(statusCode, 200, 299, true):
		return twoHundred
	case isBetween(statusCode, 400, 499, true):
		return fourHundred
	case isBetween(statusCode, 500, 599, true):
		return fiveHundred
	}
	// TODO(alexbathome) - fix this logic
	return twoHundred
}
