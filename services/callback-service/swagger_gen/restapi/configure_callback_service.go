// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/sunbirdrc/callback-service/swagger_gen/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger_gen --name CallbackService --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.CallbackServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CallbackServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.PostG2pFamapOnSearchHandler == nil {
		api.PostG2pFamapOnSearchHandler = operations.PostG2pFamapOnSearchHandlerFunc(func(params operations.PostG2pFamapOnSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pFamapOnSearch has not yet been implemented")
		})
	}
	if api.PostG2pMapperOnLinkHandler == nil {
		api.PostG2pMapperOnLinkHandler = operations.PostG2pMapperOnLinkHandlerFunc(func(params operations.PostG2pMapperOnLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperOnLink has not yet been implemented")
		})
	}
	if api.PostG2pMapperOnResolveHandler == nil {
		api.PostG2pMapperOnResolveHandler = operations.PostG2pMapperOnResolveHandlerFunc(func(params operations.PostG2pMapperOnResolveParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperOnResolve has not yet been implemented")
		})
	}
	if api.PostG2pMapperOnTxnstatusHandler == nil {
		api.PostG2pMapperOnTxnstatusHandler = operations.PostG2pMapperOnTxnstatusHandlerFunc(func(params operations.PostG2pMapperOnTxnstatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperOnTxnstatus has not yet been implemented")
		})
	}
	if api.PostG2pMapperOnUnlinkHandler == nil {
		api.PostG2pMapperOnUnlinkHandler = operations.PostG2pMapperOnUnlinkHandlerFunc(func(params operations.PostG2pMapperOnUnlinkParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperOnUnlink has not yet been implemented")
		})
	}
	if api.PutG2pMapperOnUpdateHandler == nil {
		api.PutG2pMapperOnUpdateHandler = operations.PutG2pMapperOnUpdateHandlerFunc(func(params operations.PutG2pMapperOnUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutG2pMapperOnUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
