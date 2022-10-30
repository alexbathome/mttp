package mttp

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
