package restapi

import (
	"crypto/tls"
	"github.com/ildomm/zus/handlers"
	handler_tokens "github.com/ildomm/zus/handlers/tokens"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
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
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.TokensCreateHashHandler = tokens.CreateHashHandlerFunc(handler_tokens.TokensCreateHandlerResponder)
	api.TokensGetHashesHandler = tokens.GetHashesHandlerFunc(handler_tokens.GetHashesHandlerResponder)
	api.TokensGetHashHandler = tokens.GetHashHandlerFunc(handler_tokens.GetHashHandlerResponder)

	api.OptionsAllowHandler = operations.OptionsAllowHandlerFunc(handlers.OptionsHandlerResponder)
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
