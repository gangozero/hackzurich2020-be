// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gangozero/hackzurich2020-be/generated/restapi/operations"
	"github.com/gangozero/hackzurich2020-be/generated/restapi/operations/user"
)

//go:generate swagger generate server --target ../../generated --name Hackzurich2020Be --spec ../../openapi/swagger.yaml --principal interface{}

func configureFlags(api *operations.Hackzurich2020BeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Hackzurich2020BeAPI) http.Handler {
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

	api.JSONProducer = runtime.JSONProducer()

	if api.UserDeleteProductHandler == nil {
		api.UserDeleteProductHandler = user.DeleteProductHandlerFunc(func(params user.DeleteProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteProduct has not yet been implemented")
		})
	}
	if api.UserGetChatItemsHandler == nil {
		api.UserGetChatItemsHandler = user.GetChatItemsHandlerFunc(func(params user.GetChatItemsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetChatItems has not yet been implemented")
		})
	}
	if api.UserGetProductListHandler == nil {
		api.UserGetProductListHandler = user.GetProductListHandlerFunc(func(params user.GetProductListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetProductList has not yet been implemented")
		})
	}
	if api.UserGetRecipeDetailsHandler == nil {
		api.UserGetRecipeDetailsHandler = user.GetRecipeDetailsHandlerFunc(func(params user.GetRecipeDetailsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetRecipeDetails has not yet been implemented")
		})
	}
	if api.UserGetRecipeMatchListHandler == nil {
		api.UserGetRecipeMatchListHandler = user.GetRecipeMatchListHandlerFunc(func(params user.GetRecipeMatchListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetRecipeMatchList has not yet been implemented")
		})
	}
	if api.UserInitiateCookingHandler == nil {
		api.UserInitiateCookingHandler = user.InitiateCookingHandlerFunc(func(params user.InitiateCookingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.InitiateCooking has not yet been implemented")
		})
	}
	if api.UserLoginHandler == nil {
		api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		})
	}
	if api.UserPostChatItemHandler == nil {
		api.UserPostChatItemHandler = user.PostChatItemHandlerFunc(func(params user.PostChatItemParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.PostChatItem has not yet been implemented")
		})
	}
	if api.UserRejectCookingHandler == nil {
		api.UserRejectCookingHandler = user.RejectCookingHandlerFunc(func(params user.RejectCookingParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.RejectCooking has not yet been implemented")
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
