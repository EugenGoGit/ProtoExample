package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"../../server/restapi/operations"
	"../../server/restapi/operations/key_service"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../../swagger.json

func configureFlags(api *operations.KeyProtoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KeyProtoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.KeyServiceKeyServiceAddKeyHandler = key_service.KeyServiceAddKeyHandlerFunc(func(params key_service.KeyServiceAddKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation key_service.KeyServiceAddKey has not yet been implemented")
	})
	api.KeyServiceKeyServiceDeleteKeyHandler = key_service.KeyServiceDeleteKeyHandlerFunc(func(params key_service.KeyServiceDeleteKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation key_service.KeyServiceDeleteKey has not yet been implemented")
	})
	api.KeyServiceKeyServiceGetKeyHandler = key_service.KeyServiceGetKeyHandlerFunc(func(params key_service.KeyServiceGetKeyParams) middleware.Responder {
		return middleware.NotImplemented("operation key_service.KeyServiceGetKey has not yet been implemented")
	})
	api.KeyServiceKeyServiceGetPageHandler = key_service.KeyServiceGetPageHandlerFunc(func(params key_service.KeyServiceGetPageParams) middleware.Responder {
		return middleware.NotImplemented("operation key_service.KeyServiceGetPage has not yet been implemented")
	})
	api.KeyServiceKeyServiceGetPageContinuedHandler = key_service.KeyServiceGetPageContinuedHandlerFunc(func(params key_service.KeyServiceGetPageContinuedParams) middleware.Responder {
		return middleware.NotImplemented("operation key_service.KeyServiceGetPageContinued has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
