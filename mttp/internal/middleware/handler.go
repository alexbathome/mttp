package middleware

import (
	"net/http"

	"github.com/alexbathome/mttp/mttp/internal/prom"
)

// MttpMiddlewareHandler interface is implemented by the mttpMiddlewareHandle
type MttpMiddlewareHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// mttpMiddleWareHandle
type mttpMiddleWareHandle struct {
	handlerFunc     http.HandlerFunc
	counters        prom.MttpRouteMetricer
	acceptedMethods []string
}

// NewMttpMiddlewareHandler constructs a new mttpMidleWareHandle object to be used
// by the mttpServer to handle traffic on specific routes
func NewMttpMiddlewareHandler(handlerFunc http.HandlerFunc, routeCounter prom.MttpRouteMetricer, acceptedMethods []string) *mttpMiddleWareHandle {
	return &mttpMiddleWareHandle{
		handlerFunc:     handlerFunc,
		counters:        routeCounter,
		acceptedMethods: acceptedMethods,
	}
}

// ServeHTTP is required to implement the http.Handler interface
func (m *mttpMiddleWareHandle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	middlewareResponseWriter := newMttpResponseWriter(rw)
	m.handlerFunc(middlewareResponseWriter, r)
	m.counters.IncremementStatusCode(middlewareResponseWriter.statusCode)
}
