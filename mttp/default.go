package mttp

import "net/http"

// Default Responses
func methodNotAllowedResponse(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
