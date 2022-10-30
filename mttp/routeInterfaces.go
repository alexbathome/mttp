package mttp

type Router interface {
	Route() (*route, error)
}

type RouteBuilder interface {
	AcceptMethods(...string) RouteBuilder
	RespondedToBy(routeHandler) RouteBuilder
	Build() Router
}
