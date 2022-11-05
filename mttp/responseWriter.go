package mttp

import "net/http"

// mttpResponseWriter is a struct that holds additional information on the
type mttpResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newMttpResponseWriter(rw http.ResponseWriter) *mttpResponseWriter {
	return &mttpResponseWriter{rw, http.StatusOK}
}

func (mrw *mttpResponseWriter) WriteHeader(statusCode int) {
	mrw.statusCode = statusCode
	mrw.ResponseWriter.WriteHeader(statusCode)
}

func (mrw *mttpResponseWriter) GetStatusCode() int {
	return mrw.statusCode
}
