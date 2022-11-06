package mttp

import "net/http"

type RouteBuilder interface {
	// AcceptMethods is accepts a list of HTTP Methods that are allowed
	// to be used against this route
	//
	// by default, all HTTP Methods are allowed. Except for HTTP Connect.
	AcceptMethods(...string) RouteBuilder

	// RespondedToBy takes a http.HandlerFunc, this maps the route with a
	// given handlerFunc to run against that route
	RespondedToBy(http.HandlerFunc) RouteBuilder

	// addMttpMiddleware is an internal method that is called by the ServerBuilder
	// this adds the prometheus metrics collector to the route
	addMttpMiddleware(string) RouteBuilder

	// Build requires the mttp's http.ServerMux to be passed in.
	// it builds the route, and adds it to the server's multiplexer.
	Build(mux *http.ServeMux)
}
