package mttp

import "strings"

type route struct {
	routePath       string
	handlerFunc     routeHandler
	acceptedMethods []string
}

func (r *route) Route() (*route, error) {
	return r, nil
}

// canonicaliseRoute converts a routePath into a canonical route identifier
// used for metric collection
func (r *route) canonicaliseRoute() string {
	// TODO(alexbathome) - Perhaps this needs more thought,
	// however, we should just be able to do a simple string
	// replace here.

	return strings.ReplaceAll(r.routePath, "/", "_")
}
