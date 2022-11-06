package mttp

import "net/http"

type route struct {
	routePath       string
	handlerFunc     http.HandlerFunc
	acceptedMethods []string
}
