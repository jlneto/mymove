// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CancelMoveHandlerFunc turns a function with the right signature into a cancel move handler
type CancelMoveHandlerFunc func(CancelMoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CancelMoveHandlerFunc) Handle(params CancelMoveParams) middleware.Responder {
	return fn(params)
}

// CancelMoveHandler interface for that can handle valid cancel move params
type CancelMoveHandler interface {
	Handle(CancelMoveParams) middleware.Responder
}

// NewCancelMove creates a new http.Handler for the cancel move operation
func NewCancelMove(ctx *middleware.Context, handler CancelMoveHandler) *CancelMove {
	return &CancelMove{Context: ctx, Handler: handler}
}

/*CancelMove swagger:route POST /moves/{moveId}/cancel office cancelMove

Cancels a move

Cancels the basic details of a move. The status of the move will be updated to CANCELED

*/
type CancelMove struct {
	Context *middleware.Context
	Handler CancelMoveHandler
}

func (o *CancelMove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCancelMoveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}