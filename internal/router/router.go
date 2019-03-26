package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"spikes/sample-golang/internal/dependency"
	"spikes/sample-golang/internal/handler"
	"spikes/sample-golang/pkg/instrumentation"
)

func Router(container *dependency.Container) http.Handler {
	router := mux.NewRouter()
	router.Path("/ping").HandlerFunc(handler.PingHandler).Methods(http.MethodGet)

	userHandler := handler.GetUser(container.GetUserService())
	router.Path("/user/{id}").Methods(http.MethodGet).Name("GetUserByID").HandlerFunc(handler.Wrap(userHandler))

	return instrumentation.InstrumentNewRelicOnRoutes(container.GetNewRelic(), router)
}
