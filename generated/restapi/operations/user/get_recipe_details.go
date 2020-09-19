// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRecipeDetailsHandlerFunc turns a function with the right signature into a get recipe details handler
type GetRecipeDetailsHandlerFunc func(GetRecipeDetailsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecipeDetailsHandlerFunc) Handle(params GetRecipeDetailsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRecipeDetailsHandler interface for that can handle valid get recipe details params
type GetRecipeDetailsHandler interface {
	Handle(GetRecipeDetailsParams, interface{}) middleware.Responder
}

// NewGetRecipeDetails creates a new http.Handler for the get recipe details operation
func NewGetRecipeDetails(ctx *middleware.Context, handler GetRecipeDetailsHandler) *GetRecipeDetails {
	return &GetRecipeDetails{Context: ctx, Handler: handler}
}

/*GetRecipeDetails swagger:route GET /match/{id} user getRecipeDetails

Get details of the matched recipe

*/
type GetRecipeDetails struct {
	Context *middleware.Context
	Handler GetRecipeDetailsHandler
}

func (o *GetRecipeDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRecipeDetailsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}