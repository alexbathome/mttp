package mttp

import "net/http"

type server struct {
	server     *http.Server
	promServer *http.Server
}

func (s *server) Start() error {
	if s.promServer != nil {
		go s.promServer.ListenAndServe()
	}
	return s.server.ListenAndServe()
}
