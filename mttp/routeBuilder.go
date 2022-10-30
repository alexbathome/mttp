package mttp

import "net/http"

type routeBuilder struct {
	routePath       string
	handlerFunc     routeHandler
	acceptedMethods []string
}

func NewRoute(routePath string) RouteBuilder {
	return &routeBuilder{
		routePath: routePath,
	}
}

func (r *routeBuilder) AcceptMethods(methods ...string) RouteBuilder {
	r.acceptedMethods = append(r.acceptedMethods, methods...)
	return r
}

func (r *routeBuilder) RespondedToBy(handler routeHandler) RouteBuilder {
	r.handlerFunc = handler
	return r
}

func (r *routeBuilder) Build() Router {
	return &route{
		routePath:       r.routePath,
		handlerFunc:     r.handlerFunc,
		acceptedMethods: r.acceptedMethods,
	}
}

type routeHandler func(http.ResponseWriter, *http.Request)
