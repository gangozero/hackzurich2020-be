// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetChatItemsHandlerFunc turns a function with the right signature into a get chat items handler
type GetChatItemsHandlerFunc func(GetChatItemsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetChatItemsHandlerFunc) Handle(params GetChatItemsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetChatItemsHandler interface for that can handle valid get chat items params
type GetChatItemsHandler interface {
	Handle(GetChatItemsParams, interface{}) middleware.Responder
}

// NewGetChatItems creates a new http.Handler for the get chat items operation
func NewGetChatItems(ctx *middleware.Context, handler GetChatItemsHandler) *GetChatItems {
	return &GetChatItems{Context: ctx, Handler: handler}
}

/*GetChatItems swagger:route GET /match/{id}/chat user getChatItems

Get items from the chat

*/
type GetChatItems struct {
	Context *middleware.Context
	Handler GetChatItemsHandler
}

func (o *GetChatItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetChatItemsParams()

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
