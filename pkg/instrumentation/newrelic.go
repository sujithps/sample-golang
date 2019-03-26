package instrumentation

import (
	"github.com/newrelic/go-agent"
	"net/http"
	"github.com/gorilla/mux"
	"spikes/sample-golang/pkg/appcontext"
)

func InstrumentNewRelicOnRoutes(app newrelic.Application, r *mux.Router) *mux.Router {
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if n, _ := route.GetPathTemplate(); n == "/ping" {
			return nil
		}
		name := routeName(route)
		route.Handler(instrumentHandler(app, route.GetHandler(), name))
		return nil
	})
	return r
}

func instrumentHandler(app newrelic.Application, h http.Handler, transactionName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(transactionName, w, r)
		defer txn.End()
		r = appcontext.WithNewRelicTransaction(r, txn)
		h.ServeHTTP(txn, r)
	})
}

func routeName(route *mux.Route) string {
	if nil == route {
		return ""
	}
	if n := route.GetName(); n != "" {
		return n
	}
	if n, _ := route.GetPathTemplate(); n != "" {
		return n
	}
	n, _ := route.GetHostTemplate()
	return n
}
