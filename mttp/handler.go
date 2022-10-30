package mttp

import "net/http"

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

func methodNotAllowedResponse(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func routeHasMethod(requestMethod string, routeMethods []string) bool {
	for _, method := range routeMethods {
		if method == requestMethod {
			return true
		}
	}
	return false
}
