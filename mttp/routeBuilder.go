package mttp

import (
	"net/http"
	"strings"

	"github.com/alexbathome/mttp/mttp/internal/middleware"
	"github.com/alexbathome/mttp/mttp/internal/prom"
)

var defaultAcceptedMethods = []string{
	http.MethodGet,
	http.MethodDelete,
	http.MethodPut,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodOptions,
	http.MethodTrace,
}

type routeBuilder struct {
	routePath       string
	acceptedMethods []string

	handler     http.Handler
	handlerFunc http.HandlerFunc
}

func NewRoute(routePath string) RouteBuilder {
	return &routeBuilder{
		routePath:       routePath,
		acceptedMethods: defaultAcceptedMethods,
	}
}

func (r *routeBuilder) AcceptMethods(methods ...string) RouteBuilder {
	// empty out the default
	r.acceptedMethods = []string{}
	r.acceptedMethods = append(r.acceptedMethods, methods...)
	return r
}

func (r *routeBuilder) RespondedToBy(handler http.HandlerFunc) RouteBuilder {
	r.handlerFunc = handler
	return r
}

func (r *routeBuilder) addMttpMiddleware(serverName string) RouteBuilder {
	r.handler = middleware.NewMttpMiddlewareHandler(
		http.HandlerFunc(r.handlerFunc), prom.NewMttpRouteMetrics(serverName, r.canonicalRoutePath()), r.acceptedMethods)
	return r
}

// canonicalRoutePath converts a routePath into a canonical route identifier
// used for metric collection
func (r *routeBuilder) canonicalRoutePath() string {
	// TODO(alexbathome) - Perhaps this needs more thought,
	// however, we should just be able to do a simple string
	// replace here.

	// first strip the beginning "/" and ending "/"
	routePath := strings.TrimPrefix(r.routePath, "/")
	routePath = strings.TrimSuffix(routePath, "/")
	return strings.ReplaceAll(routePath, "/", "_")
}

func (r *routeBuilder) Build(mux *http.ServeMux) {
	rb := &route{
		routePath:       r.routePath,
		handlerFunc:     r.handlerFunc,
		acceptedMethods: r.acceptedMethods,
	}

	if r.handler != nil {
		mux.Handle(rb.routePath, r.handler)
	} else {
		mux.HandleFunc(rb.routePath, r.handlerFunc)
	}
}
