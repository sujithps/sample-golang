package middleware

import (
	"net/http"

	"fmt"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/logger"
	sentry "github.com/getsentry/raven-go"
	"github.com/urfave/negroni"
)

func Recover() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		defer func() {
			if err := recover(); err != nil {
				logger.NewContextLogger(r.Context()).Error("Recover", fmt.Sprintf("Recovered from panic: %+v", err), nil)
				sentry.CaptureError(err.(error), map[string]string{})
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next(w, r)
	})
}
