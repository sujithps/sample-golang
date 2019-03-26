package middleware

import (
	"github.com/urfave/negroni"
	"net/http"
	"spikes/sample-golang/pkg/appcontext"
	"spikes/sample-golang/pkg/constant"
	"spikes/sample-golang/pkg/logger"
	"time"
)

const (
	requestMethod        = "HTTPMethod"
	requestPath          = "RequestURL"
	requestURLQueryParam = "RequestQueryParam"
	responseStatusCode   = "ResponseCode"
	responseStatusText   = "ResponseCodeText"
	responseTimeTaken    = "ResponseTimeTaken"
)

func HTTPStatLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path != "/ping" {
			correlationID := appcontext.GetCorrelationID(r.Context())

			start := time.Now()
			next(rw, r)
			duration := time.Since(start)
			res := rw.(negroni.ResponseWriter)

			logger.NonContext.Info("StatMiddleware", "Completed HTTP request", map[string]interface{}{
				requestMethod:                r.Method,
				requestPath:                  r.URL.Path,
				requestURLQueryParam:         r.URL.Query(),
				responseStatusCode:           res.Status(),
				responseStatusText:           http.StatusText(res.Status()),
				responseTimeTaken:            duration.Seconds(),
				constant.CorrelationIDHeader: correlationID,
			})
		} else {
			next(rw, r)
		}
	})
}
