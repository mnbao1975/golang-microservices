package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GzipHandler handles something
type GzipHandler struct {
}

// GzipMiddleware does gzip
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// Create a gzipped response
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()

			return
		}

		// Handle normal
		next.ServeHTTP(rw, r)
	})
}

// WrappedResponseWriter is a wrapper of http.ResponseWriter
type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter construct WrappedResponseWriter
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw, gw}
}

// Satisfy ResposeWriter's interface

// Header func implementation
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

// Write func implementation
func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

//WriteHeader func implementation
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

// Flush flused any pending compressed data to the underlying writer
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
