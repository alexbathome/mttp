package mttp

type ServerBuilder interface {
	WithRoutes(...RouteBuilder) ServerBuilder
	WithMetrics(metricsPort string) ServerBuilder
	Build() Server
}

type Server interface {
	Start() int
}
