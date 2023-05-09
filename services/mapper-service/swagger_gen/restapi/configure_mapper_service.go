// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger_gen --name MapperService --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.MapperServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MapperServiceAPI) http.Handler {
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

	if api.PostG2pFamapSearchHandler == nil {
		api.PostG2pFamapSearchHandler = operations.PostG2pFamapSearchHandlerFunc(func(params operations.PostG2pFamapSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pFamapSearch has not yet been implemented")
		})
	}
	if api.PostG2pMapperLinkHandler == nil {
		api.PostG2pMapperLinkHandler = operations.PostG2pMapperLinkHandlerFunc(func(params operations.PostG2pMapperLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperLink has not yet been implemented")
		})
	}
	if api.PostG2pMapperResolveHandler == nil {
		api.PostG2pMapperResolveHandler = operations.PostG2pMapperResolveHandlerFunc(func(params operations.PostG2pMapperResolveParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperResolve has not yet been implemented")
		})
	}
	if api.PostG2pMapperTxnstatusHandler == nil {
		api.PostG2pMapperTxnstatusHandler = operations.PostG2pMapperTxnstatusHandlerFunc(func(params operations.PostG2pMapperTxnstatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperTxnstatus has not yet been implemented")
		})
	}
	if api.PostG2pMapperUnlinkHandler == nil {
		api.PostG2pMapperUnlinkHandler = operations.PostG2pMapperUnlinkHandlerFunc(func(params operations.PostG2pMapperUnlinkParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostG2pMapperUnlink has not yet been implemented")
		})
	}
	if api.PutG2pMapperUpdateHandler == nil {
		api.PutG2pMapperUpdateHandler = operations.PutG2pMapperUpdateHandlerFunc(func(params operations.PutG2pMapperUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutG2pMapperUpdate has not yet been implemented")
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
