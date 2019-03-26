package middleware

import (
	"github.com/satori/go.uuid"
	"github.com/urfave/negroni"
	"net/http"
	"spikes/sample-golang/pkg/appcontext"
	"spikes/sample-golang/pkg/constant"
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
