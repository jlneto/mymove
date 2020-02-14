// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// NewUpdateMoveTaskOrderDestinationAddressParams creates a new UpdateMoveTaskOrderDestinationAddressParams object
// no default values defined in spec.
func NewUpdateMoveTaskOrderDestinationAddressParams() UpdateMoveTaskOrderDestinationAddressParams {

	return UpdateMoveTaskOrderDestinationAddressParams{}
}

// UpdateMoveTaskOrderDestinationAddressParams contains all the bound params for the update move task order destination address operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMoveTaskOrderDestinationAddress
type UpdateMoveTaskOrderDestinationAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	DestinationAddress *primemessages.Address
	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMoveTaskOrderDestinationAddressParams() beforehand.
func (o *UpdateMoveTaskOrderDestinationAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body primemessages.Address
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("destinationAddress", "body"))
			} else {
				res = append(res, errors.NewParseError("destinationAddress", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.DestinationAddress = &body
			}
		}
	} else {
		res = append(res, errors.Required("destinationAddress", "body"))
	}
	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *UpdateMoveTaskOrderDestinationAddressParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}
