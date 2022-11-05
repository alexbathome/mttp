package mttp

import "net/http"

// Default Responses
func methodNotAllowedResponse(rw http.ResponseWriter, r *http.Request) http.ResponseWriter {
	rw.WriteHeader(http.StatusMethodNotAllowed)
	return rw
}
