package router

import (
	"github.com/sujithps/sample-golang/internal/dependency"
	"github.com/sujithps/sample-golang/internal/handler"
	"github.com/sujithps/sample-golang/pkg/instrumentation"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(container *dependency.Container) http.Handler {
	router := mux.NewRouter()
	router.Path("/ping").HandlerFunc(handler.PingHandler).Methods(http.MethodGet)

	userHandler := handler.GetUser(container.GetUserService())
	router.Path("/user/{id}").Methods(http.MethodGet).Name("GetUserByID").HandlerFunc(handler.Wrap(userHandler))

	return instrumentation.InstrumentNewRelicOnRoutes(container.GetNewRelic(), router)
}
