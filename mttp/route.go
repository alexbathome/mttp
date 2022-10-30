package mttp

import "net/http"

type Router interface {
	Route() (*route, error)
}

type route struct {
	routePath       string
	handlerFunc     routeHandler
	acceptedMethods []string
}

func (r *route) Route() (*route, error) {
	return r, nil
}

type RouteBuilder interface {
	AcceptMethods(...string) RouteBuilder
	RespondedToBy(routeHandler) RouteBuilder
	Build() Router
}

type routeBuilder struct {
	routePath       string
	handlerFunc     routeHandler
	acceptedMethods []string
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

func NewRoute(routePath string) RouteBuilder {
	return &routeBuilder{
		routePath: routePath,
	}
}

type routeHandler func(http.ResponseWriter, *http.Request)
