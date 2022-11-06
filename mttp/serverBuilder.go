package mttp

import (
	"fmt"
	"net/http"

	"github.com/alexbathome/mttp/mttp/internal/validator"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// serverBuilder is the internal struct used by the mttp server
type serverBuilder struct {
	name    string
	address string
	port    string

	useMetrics  bool
	metricsPort string

	routes []RouteBuilder

	serverMux *http.ServeMux
}

// NewServer is the entry point to the mttp library, this creates a ServerBuilder interface
// which can be used to configure the mttp server.
func NewServer(name, address, port string) ServerBuilder {
	return &serverBuilder{
		name:    name,
		address: address,
		port:    port,
	}
}

// WithRoutes is a builder method which adds routes to the mttp server.
func (s *serverBuilder) WithRoutes(routes ...RouteBuilder) ServerBuilder {
	s.routes = append(s.routes, routes...)
	return s
}

// WithMetrics enables prometheus metrics for our mttp server, it accepts a
// single parameter which is the port that the metrics server will be available
// on.
func (s *serverBuilder) WithMetrics(metricsPort string) ServerBuilder {
	s.metricsPort = metricsPort
	s.useMetrics = true
	return s
}

// Build is the function that ultimately constructs our mttp server as per the definition
// specified in the builder pattern.
func (s *serverBuilder) Build() (Server, error) {

	// validate!
	err := s.validate()
	if err != nil {
		return nil, err
	}

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

// createServer creates our internal http server
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
		if s.useMetrics {
			r.addMttpMiddleware(s.name)
		}
		r.Build(s.serverMux)
	}
	return nil
}

// validate method verifies that the configuration for the server is suitable for
// implementation. The checks it includes are:
// * Server name meets expectations
// TODO - add more validations if required
func (s *serverBuilder) validate() error {
	// check server name
	err := validator.ValidateServerName(s.name)
	if err != nil {
		return err
	}
	return nil
}
