package prom

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func unregisterAllCounters(mrm mttpRouteMetrics) {
	for _, v := range mrm.counters {
		prometheus.Unregister(v)
	}
}

func TestGetStatusFamily(t *testing.T) {
	tt := map[string]struct {
		want int
		have int
	}{
		"Status code 201 returns 200": {
			want: 200,
			have: 201,
		},
		"Status code 404 returns 400": {
			want: 400,
			have: 404,
		},
	}

	for _, test := range tt {
		sut := NewMttpRouteMetrics("blah", "blah")
		if r := sut.getStatusFamily(test.have); r != test.want {
			t.Fail()
		}
		// since prometheus counters are global, we have to unregister them
		// for the next test
		unregisterAllCounters(*sut)
	}
}
