// Code generated by go-swagger; DO NOT EDIT.

package backup_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateServiceMemberBackupContactHandlerFunc turns a function with the right signature into a create service member backup contact handler
type CreateServiceMemberBackupContactHandlerFunc func(CreateServiceMemberBackupContactParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceMemberBackupContactHandlerFunc) Handle(params CreateServiceMemberBackupContactParams) middleware.Responder {
	return fn(params)
}

// CreateServiceMemberBackupContactHandler interface for that can handle valid create service member backup contact params
type CreateServiceMemberBackupContactHandler interface {
	Handle(CreateServiceMemberBackupContactParams) middleware.Responder
}

// NewCreateServiceMemberBackupContact creates a new http.Handler for the create service member backup contact operation
func NewCreateServiceMemberBackupContact(ctx *middleware.Context, handler CreateServiceMemberBackupContactHandler) *CreateServiceMemberBackupContact {
	return &CreateServiceMemberBackupContact{Context: ctx, Handler: handler}
}

/*CreateServiceMemberBackupContact swagger:route POST /service_members/{serviceMemberId}/backup_contacts backup_contacts createServiceMemberBackupContact

Submits backup contact for a logged-in user

Creates an instance of a backup contact tied to a service member user

*/
type CreateServiceMemberBackupContact struct {
	Context *middleware.Context
	Handler CreateServiceMemberBackupContactHandler
}

func (o *CreateServiceMemberBackupContact) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateServiceMemberBackupContactParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
