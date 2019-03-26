package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func JSONAPI() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set(contentTypeHeader, jsonContentType)
		next(w, r)
	})
}
