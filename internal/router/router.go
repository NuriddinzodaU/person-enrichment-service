package router

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"person-service/internal/config"
	"person-service/internal/handlers"
	"person-service/internal/pkg/service"
)

var EntryPoint = fx.Options(
	fx.Invoke(
		NewRouter,
	),
)

type dependencies struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.IConfig
	SVC       service.IService
	Handler   handlers.IHandler
}

func NewRouter(d dependencies) {
	server := mux.NewRouter()
	mainRoute := server.PathPrefix("/api").Subrouter()
	routeVer := mainRoute.PathPrefix("/v1").Subrouter()

	//edit/add person
	routeVer.HandleFunc("/create-person", d.Handler.CreatePerson()).Methods("POST", "OPTIONS")
	routeVer.HandleFunc("/update-person", d.Handler.UpdatePerson()).Methods("PUT", "OPTIONS")

	//get person
	routeVer.Path("/person").HandlerFunc(d.Handler.GetPerson()).Methods(http.MethodGet, http.MethodOptions).
		Queries("id", "{id}")
	routeVer.Path("/persons").HandlerFunc(d.Handler.GetAllPerson()).Methods(http.MethodGet, http.MethodOptions).
		Queries("count", "{count}").Queries("page", "{page}")

	//Delete person
	routeVer.Path("/persons").
		Queries("id", "{id}").
		HandlerFunc(d.Handler.DeletePerson()).Methods(http.MethodDelete, http.MethodOptions)
	routeVer.HandleFunc("/person", d.Handler.UpdatePerson()).Methods("DELETE", "OPTIONS")

	srv := http.Server{
		Addr:    d.SVC.ConfigInstance().GetString("api.server.port"),
		Handler: server,
	}

	d.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				d.SVC.LoggerInstance().Info("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)

}
