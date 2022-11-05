package mttp

func routeHasMethod(requestMethod string, routeMethods []string) bool {
	for _, method := range routeMethods {
		if method == requestMethod {
			return true
		}
	}
	return false
}

func isBetween(point, start, end int, inclusive bool) bool {
	if !inclusive {
		return start < point && point < end
	}
	return start <= point && point <= end
}
