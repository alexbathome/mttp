/*
Package mttp implements simple HTTP server routing. Package mttp aims to provide a simple, and effective
method of creating a HTTP router, all the while keeping the configuration clear and concise.

The mttp package leans on the builder pattern to exhibit a declarative way of configuring your HTTP server
and it's routes.

To get started, you can simply define a simple server such as the example below

	import (
		"net/http"

		"github.com/alexbathome/mttp/mttp"
	)

	func webHandler(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello from WEB"))
	}

	var server = mttp.NewServer("myapi", "127.0.0.1", "8080").
		WithMetrics("9090").
		WithRoutes(
			mttp.NewRoute("/web").
				AcceptMethods(http.MethodGet).
				RespondedToBy(webHandler),
		)

	func main() {
		s, err := server.Build()
		if err != nil {
			panic(err)
		}
		err = s.Start()
		if err != nil {
			panic(err)
		}
	}
*/
package mttp
