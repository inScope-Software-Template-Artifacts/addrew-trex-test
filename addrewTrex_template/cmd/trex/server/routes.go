package server

import (
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/openshift-online/addrewTrex/cmd/addrewTrex-service/server/logging"
	"github.com/openshift-online/addrewTrex/pkg/api"
	"github.com/openshift-online/addrewTrex/pkg/auth"
	"github.com/openshift-online/addrewTrex/pkg/db"
	"github.com/openshift-online/addrewTrex/pkg/handlers"
	"github.com/openshift-online/addrewTrex/pkg/logger"
)

func (s *apiServer) routes() *mux.Router {
	services := &env().Services

	openAPIDefinitions, err := s.loadOpenAPISpec("openapi.yaml")
	if err != nil {
		check(err, "Can't load OpenAPI specification")
	}

	dinosaurHandler := handlers.NewDinosaurHandler(services.Dinosaurs(), services.Generic())
	errorsHandler := handlers.NewErrorsHandler()

	var authMiddleware auth.JWTMiddleware
	authMiddleware = &auth.AuthMiddlewareMock{}
	if env().Config.Server.EnableJWT {
		var err error
		authMiddleware, err = auth.NewAuthMiddleware()
		check(err, "Unable to create auth middleware")
	}
	if authMiddleware == nil {
		check(err, "Unable to create auth middleware: missing middleware")
	}

	authzMiddleware := auth.NewAuthzMiddlewareMock()
	if env().Config.Server.EnableAuthz {
		// TODO: authzMiddleware, err = auth.NewAuthzMiddleware()
		check(err, "Unable to create authz middleware")
	}

	// mainRouter is top level "/"
	mainRouter := mux.NewRouter()
	mainRouter.NotFoundHandler = http.HandlerFunc(api.SendNotFound)

	// Operation ID middleware sets a relatively unique operation ID in the context of each request for debugging purposes
	mainRouter.Use(logger.OperationIDMiddleware)

	// Request logging middleware logs pertinent information about the request and response
	mainRouter.Use(logging.RequestLoggingMiddleware)

	//  /api/addrewTrex
	apiRouter := mainRouter.PathPrefix("/api/addrewTrex").Subrouter()
	apiRouter.HandleFunc("", api.SendAPI).Methods(http.MethodGet)

	//  /api/addrewTrex/v1
	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()
	apiV1Router.HandleFunc("", api.SendAPIV1).Methods(http.MethodGet)
	apiV1Router.HandleFunc("/", api.SendAPIV1).Methods(http.MethodGet)

	//  /api/addrewTrex/v1/openapi
	apiV1Router.HandleFunc("/openapi", handlers.NewOpenAPIHandler(openAPIDefinitions).Get).Methods(http.MethodGet)
	registerApiMiddleware(apiV1Router)

	//  /api/addrewTrex/v1/errors
	apiV1ErrorsRouter := apiV1Router.PathPrefix("/errors").Subrouter()
	apiV1ErrorsRouter.HandleFunc("", errorsHandler.List).Methods(http.MethodGet)
	apiV1ErrorsRouter.HandleFunc("/{id}", errorsHandler.Get).Methods(http.MethodGet)

	//  /api/addrewTrex/v1/dinosaurs
	apiV1DinosaursRouter := apiV1Router.PathPrefix("/dinosaurs").Subrouter()
	apiV1DinosaursRouter.HandleFunc("", dinosaurHandler.List).Methods(http.MethodGet)
	apiV1DinosaursRouter.HandleFunc("/{id}", dinosaurHandler.Get).Methods(http.MethodGet)
	apiV1DinosaursRouter.HandleFunc("", dinosaurHandler.Create).Methods(http.MethodPost)
	apiV1DinosaursRouter.HandleFunc("/{id}", dinosaurHandler.Patch).Methods(http.MethodPatch)
	apiV1DinosaursRouter.HandleFunc("/{id}", dinosaurHandler.Delete).Methods(http.MethodDelete)
	apiV1DinosaursRouter.Use(authMiddleware.AuthenticateAccountJWT)

	apiV1DinosaursRouter.Use(authzMiddleware.AuthorizeApi)

	return mainRouter
}

func registerApiMiddleware(router *mux.Router) {
	router.Use(MetricsMiddleware)

	router.Use(
		func(next http.Handler) http.Handler {
			return db.TransactionMiddleware(next, env().Database.SessionFactory)
		},
	)

	router.Use(gorillahandlers.CompressHandler)
}
