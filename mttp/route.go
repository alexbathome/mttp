package mttp

type route struct {
	routePath       string
	handlerFunc     routeHandler
	acceptedMethods []string
}

func (r *route) Route() (*route, error) {
	return r, nil
}
