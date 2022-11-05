package mttp

type ServerBuilder interface {
	WithRoutes(...RouteBuilder) ServerBuilder
	WithMetrics(metricsPort string) ServerBuilder
	Build() (Server, error)
}

type Server interface {
	Start() error
}
