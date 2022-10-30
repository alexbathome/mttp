package mttp

import (
	"fmt"
	"net/http"
)

type mttpHandler struct {
	mux             http.ServeMux
	handlerFunc     http.HandlerFunc
	acceptedMethods []string
}

func (m *mttpHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if !routeHasMethod(r.Method, m.acceptedMethods) {
		methodNotAllowedResponse(rw, r)
		return
	}
	m.handlerFunc(rw, r)
}

func startServer(s server) error {
	listenAndServeAddr := fmt.Sprintf("%s", s.address)
	fmt.Printf("Server address is %s\n", listenAndServeAddr)
	handler, err := createMuxHandler(s)
	if err != nil {
		return err
	}
	err = http.ListenAndServe(listenAndServeAddr, handler)
	if err != nil {
		return err
	}
	return nil
}

func createMuxHandler(s server) (http.Handler, error) {
	myMux := http.NewServeMux()
	for _, r := range s.routes {
		rBuilt, err := r.Build().Route()
		if err != nil {
			return nil, err
		}
		myMux.Handle(rBuilt.routePath, createRouteHandler(*myMux, *rBuilt))
	}
	return myMux, nil
}

func createRouteHandler(m http.ServeMux, r route) http.Handler {
	h := mttpHandler{
		mux:             m,
		handlerFunc:     http.HandlerFunc(r.handlerFunc),
		acceptedMethods: r.acceptedMethods,
	}
	return &h
}
