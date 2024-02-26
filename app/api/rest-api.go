package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"hello-k8s/app"
	"hello-k8s/env"
)

// RESTServer manages the REST API for this microservice
type RESTServer struct {
	config env.RESTConfig
	router *mux.Router
	app    *app.Hello
	server *http.Server
}

// Initialize configures the REST API Server
func (rest *RESTServer) Initialize(config *env.Config) {
	rest.config = config.REST

	rest.app = new(app.Hello)
	rest.app.Initialize(config)

	rest.initializeRoutes()
}

// Run starts the REST API server
func (rest *RESTServer) Run() {
	// Use the custom not found handler
	rest.router.NotFoundHandler = new(notFoundHandler)

	// Enable CORS for all origins, methods and common headers
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPatch, http.MethodPut})

	// Configure the HTTP server
	rest.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", rest.config.Port),
		Handler: handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(rest.router),
	}

	// Start handling requests
	err := rest.server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

// Destroy stops the REST server and performs any necessary cleanup on nested structures
func (rest *RESTServer) Destroy() {
	rest.server.Shutdown(context.TODO())
	rest.app.Destroy()
}

func (rest *RESTServer) initializeRoutes() {
	rest.router = mux.NewRouter()
	rest.router.HandleFunc(rest.config.BasePath, rest.handleHello).Methods(http.MethodGet)
	rest.router.HandleFunc(rest.config.BasePath+"/hello", rest.handleHello).Methods(http.MethodGet)
	rest.router.HandleFunc(rest.config.BasePath+"/attack", rest.handleAttack).Methods(http.MethodGet)
	rest.enableProfiling()
}

func (rest *RESTServer) enableProfiling() {
	rest.router.HandleFunc(rest.config.BasePath+"/debug/pprof/", pprof.Index)
	rest.router.HandleFunc(rest.config.BasePath+"/debug/pprof/cmdline", pprof.Cmdline)
	rest.router.HandleFunc(rest.config.BasePath+"/debug/pprof/profile", pprof.Profile)
	rest.router.HandleFunc(rest.config.BasePath+"/debug/pprof/symbol", pprof.Symbol)
	rest.router.HandleFunc(rest.config.BasePath+"/debug/pprof/trace", pprof.Trace)
	rest.router.Handle(rest.config.BasePath+"/debug/pprof/goroutine", pprof.Handler("goroutine"))
	rest.router.Handle(rest.config.BasePath+"/debug/pprof/heap", pprof.Handler("heap"))
	rest.router.Handle(rest.config.BasePath+"/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	rest.router.Handle(rest.config.BasePath+"/debug/pprof/block", pprof.Handler("block"))
	rest.router.Handle(rest.config.BasePath+"/debug/vars", http.DefaultServeMux)
}
