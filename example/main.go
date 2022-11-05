package main

import (
	"net/http"

	"github.com/alexbathome/mttp/mttp"
)

func webHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello from WEB"))
}

func apiHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello From API"))
}

func apiTestHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("This also works"))
}

var server = mttp.NewServer("my-first-api", "127.0.0.1", "8080").
	WithMetrics("9090").
	WithRoutes(
		mttp.NewRoute("/web").
			AcceptMethods(http.MethodGet).
			RespondedToBy(webHandler),
		mttp.NewRoute("/api").
			AcceptMethods(http.MethodPost, http.MethodGet).
			RespondedToBy(apiHandler),
		mttp.NewRoute("/api/test").
			AcceptMethods(http.MethodGet).
			RespondedToBy(apiTestHandler),
	)

func main() {
	s, err := server.Build()
	if err != nil {
		print(err)
	}
	s.Start()
}
