package middleware

import (
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/appcontext"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/constant"
	"github.com/satori/go.uuid"
	"github.com/urfave/negroni"
	"net/http"
)

func CorrelationID() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		correlationID := r.Header.Get(constant.CorrelationIDHeader)
		if len(correlationID) == 0 {
			correlationID = uuid.NewV4().String()
		}
		next(rw, appcontext.WithCorrelationID(r, correlationID))
	})
}
