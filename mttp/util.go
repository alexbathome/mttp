package mttp

func routeHasMethod(requestMethod string, routeMethods []string) bool {
	for _, method := range routeMethods {
		if method == requestMethod {
			return true
		}
	}
	return false
}
