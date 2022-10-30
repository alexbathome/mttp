package mttp

type Server interface {
	Start() int
}

type server struct {
	name        string
	address     string
	port        string
	useMetrics  bool
	metricsPort string
	routes      []RouteBuilder
}

func (s *server) Start() int {
	err := startServer(*s)
	if err != nil {
		return 1
	}
	return 0
}

type ServerBuilder interface {
	WithRoutes(...RouteBuilder) ServerBuilder
	WithMetrics(metricsPort string) ServerBuilder
	Build() Server
}

type serverBuilder struct {
	name        string
	address     string
	port        string
	useMetrics  bool
	metricsPort string
	routes      []RouteBuilder
}

func NewServer(name, address string) ServerBuilder {
	return &serverBuilder{
		name:    name,
		address: address,
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

func (s *serverBuilder) Build() Server {
	return &server{
		name:        s.name,
		address:     s.address,
		port:        s.port,
		useMetrics:  s.useMetrics,
		metricsPort: s.metricsPort,
		routes:      s.routes,
	}
}
