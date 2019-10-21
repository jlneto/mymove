// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAdminUserHandlerFunc turns a function with the right signature into a get admin user handler
type GetAdminUserHandlerFunc func(GetAdminUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAdminUserHandlerFunc) Handle(params GetAdminUserParams) middleware.Responder {
	return fn(params)
}

// GetAdminUserHandler interface for that can handle valid get admin user params
type GetAdminUserHandler interface {
	Handle(GetAdminUserParams) middleware.Responder
}

// NewGetAdminUser creates a new http.Handler for the get admin user operation
func NewGetAdminUser(ctx *middleware.Context, handler GetAdminUserHandler) *GetAdminUser {
	return &GetAdminUser{Context: ctx, Handler: handler}
}

/*GetAdminUser swagger:route GET /admin_users/{adminUserId} admin_users getAdminUser

Fetch a specific admin user

Returns a single admin user

*/
type GetAdminUser struct {
	Context *middleware.Context
	Handler GetAdminUserHandler
}

func (o *GetAdminUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAdminUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}