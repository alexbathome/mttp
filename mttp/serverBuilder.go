package mttp

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type serverBuilder struct {
	name    string
	address string
	port    string

	useMetrics  bool
	metricsPort string

	routes []RouteBuilder

	serverMux *http.ServeMux

	promServer *http.Server
}

func NewServer(name, address, port string) ServerBuilder {
	return &serverBuilder{
		name:    name,
		address: address,
		port:    port,
	}
}

func (s *serverBuilder) WithRoutes(routes ...RouteBuilder) ServerBuilder {
	s.routes = append(s.routes, routes...)
	return s
}

func (s *serverBuilder) WithMetrics(metricsPort string) ServerBuilder {
	s.metricsPort = metricsPort
	s.useMetrics = true
	return s
}

func (s *serverBuilder) Build() (Server, error) {

	// create the server
	appServer, err := s.createServer()
	if err != nil {
		return nil, err
	}

	// create the prom server
	var promServer *http.Server
	if s.useMetrics {
		promServer, err = s.createPromServer()
		if err != nil {
			return nil, err
		}
	}

	return &server{
		server:     appServer,
		promServer: promServer,
	}, nil
}

func (s *serverBuilder) createServer() (*http.Server, error) {
	err := s.createServerHandler()
	if err != nil {
		return nil, err
	}
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.address, s.port),
		Handler: s.serverMux,
	}
	return &server, nil
}

// createPromServer - This method creates our prometheus server.
// It is also required to register all of our prometheus metric/counters
func (s *serverBuilder) createPromServer() (*http.Server, error) {
	promHandler := http.DefaultServeMux
	promHandler.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.address, s.metricsPort),
		Handler: promHandler,
	}
	return &server, nil
}

// createServerHandler - Is responsible creating the HTTP server's mux. This method
// additionally also parses the handles for each route provided to the server.
func (s *serverBuilder) createServerHandler() error {
	s.serverMux = http.DefaultServeMux
	for _, r := range s.routes {
		route, err := r.Build().Route()
		if err != nil {
			return err
		}

		// if we are running prometheus also, we need to register the routes with the
		// counters
		if s.useMetrics {
			counters := registerRouteWithPrometheusCounters(s.name, route.canonicaliseRoute())
			s.serverMux.Handle(route.routePath, createRouteHandlerWithMetrics(*s.serverMux, *route, counters))
		} else {
			// Using else instead of "continue" here just to make it slightly more readable
			s.serverMux.Handle(route.routePath, createRouteHandler(*s.serverMux, *route))
		}
	}
	return nil
}
