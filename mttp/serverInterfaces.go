package mttp

type ServerBuilder interface {
	// WithRoutes is the builder pattern implementation to provide the mttp server
	// with routes to serve
	WithRoutes(...RouteBuilder) ServerBuilder
	// WithMetrics enables the MTTP library's middle-ware.
	//
	// By default, this middleware captures HTTP status codes from your
	// handlers, and publishes them with a prometheus server that runs
	// on the given port.
	WithMetrics(metricsPort string) ServerBuilder

	// Build constructs the mttp server, it returns an implementation of the mttp.Server
	// interface, and an error.
	//
	// Build internally wraps the route's handler funcs with mttp's middleware if mttp.WithMetrics()
	// has been called on the ServerBuilder
	Build() (Server, error)
}

// Server is the mttp.Server interface
type Server interface {
	// Start starts the mttp server, and begins listening to HTTP traffic
	// on the server's provided address
	Start() error
}
