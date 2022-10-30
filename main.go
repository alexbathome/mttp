package main

import (
	"net/http"
	"os"

	"github.com/alexbathome/mttp/mttp"
)

func webHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello from WEB"))
}

func apiHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello From API"))
}

var server = mttp.NewServer("my-first-api", "127.0.0.1:8080").
	WithRoutes(
		mttp.NewRoute("/web").
			AcceptMethods(http.MethodGet).
			RespondedToBy(webHandler),
		mttp.NewRoute("/api").
			AcceptMethods(http.MethodPost).
			RespondedToBy(apiHandler),
	).Build()

func main() {
	os.Exit(server.Start())
}
