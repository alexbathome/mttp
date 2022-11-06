package middleware

import "net/http"

// mttpResponseWriter is a struct that implements the http.ResponseWriter interface
// it allows us to also hold additional information on the status code that the
// handler might return, this is exclusively used in order for us to inject
// this information into our prometheus middleware.
type mttpResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// constructs a new mttpResponseWriter, it defaults to the http.StatusOk statuscode.
func newMttpResponseWriter(rw http.ResponseWriter) *mttpResponseWriter {
	return &mttpResponseWriter{rw, http.StatusOK}
}

// WriteHeader is a function that completes the http.ResponseWriter interface implemenentation
// it additionally sets the mttpResponseWriter's statusCode field
func (mrw *mttpResponseWriter) WriteHeader(statusCode int) {
	mrw.statusCode = statusCode
	mrw.ResponseWriter.WriteHeader(statusCode)
}
