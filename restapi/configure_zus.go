// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ildomm/zus/restapi/operations"
	"github.com/ildomm/zus/restapi/operations/tokens"
)

//go:generate swagger generate server --target ..\..\zus --name Zus --spec ..\spec\swagger.yaml

func configureFlags(api *operations.ZusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ZusAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.TokensCreateHashHandler == nil {
		api.TokensCreateHashHandler = tokens.CreateHashHandlerFunc(func(params tokens.CreateHashParams) middleware.Responder {
			return middleware.NotImplemented("operation tokens.CreateHash has not yet been implemented")
		})
	}
	if api.TokensGetHashsHandler == nil {
		api.TokensGetHashsHandler = tokens.GetHashsHandlerFunc(func(params tokens.GetHashsParams) middleware.Responder {
			return middleware.NotImplemented("operation tokens.GetHashs has not yet been implemented")
		})
	}
	if api.TokensGetURLStatsHandler == nil {
		api.TokensGetURLStatsHandler = tokens.GetURLStatsHandlerFunc(func(params tokens.GetURLStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation tokens.GetURLStats has not yet been implemented")
		})
	}
	if api.OptionsAllowHandler == nil {
		api.OptionsAllowHandler = operations.OptionsAllowHandlerFunc(func(params operations.OptionsAllowParams) middleware.Responder {
			return middleware.NotImplemented("operation .OptionsAllow has not yet been implemented")
		})
	}

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
func configureServer(s *http.Server, scheme, addr string) {
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
