package mttp

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type mttpHandler struct {
	mux             http.ServeMux
	handlerFunc     http.HandlerFunc
	promCounters    map[string]prometheus.Counter
	acceptedMethods []string
}

func (m *mttpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// convert the httpResponseWriter to our managed mttpResponseWriter
	responseWriter := newMttpResponseWriter(rw)
	if !routeHasMethod(r.Method, m.acceptedMethods) {
		fmt.Println("Method not found!")
		methodNotAllowedResponse(responseWriter, r)
		fmt.Println("Modified!")
		if m.promCounters != nil {
			// We have prom metrics enabled.
			incrementStatusCounter(responseWriter, m.promCounters)
			return
		}
	}

	if m.promCounters != nil {
		m.handlerFunc(responseWriter, r)
		incrementStatusCounter(responseWriter, m.promCounters)
	}
}

func createRouteHandler(m http.ServeMux, r route) http.Handler {
	h := mttpHandler{
		mux:             m,
		handlerFunc:     http.HandlerFunc(r.handlerFunc),
		acceptedMethods: r.acceptedMethods,
	}
	return &h
}

func createRouteHandlerWithMetrics(m http.ServeMux, r route, counters map[string]prometheus.Counter) http.Handler {
	h := mttpHandler{
		mux:             m,
		handlerFunc:     http.HandlerFunc(r.handlerFunc),
		acceptedMethods: r.acceptedMethods,
		promCounters:    counters,
	}
	return &h
}
