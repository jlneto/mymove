// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOShipmentReader is a Reader for the UpdateMTOShipment structure.
type UpdateMTOShipmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMTOShipmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMTOShipmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMTOShipmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMTOShipmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMTOShipmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMTOShipmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateMTOShipmentPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMTOShipmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMTOShipmentOK creates a UpdateMTOShipmentOK with default headers values
func NewUpdateMTOShipmentOK() *UpdateMTOShipmentOK {
	return &UpdateMTOShipmentOK{}
}

/*UpdateMTOShipmentOK handles this case with default header values.

updated instance of mto shipment
*/
type UpdateMTOShipmentOK struct {
	Payload *primemessages.MTOShipment
}

func (o *UpdateMTOShipmentOK) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOShipmentOK) GetPayload() *primemessages.MTOShipment {
	return o.Payload
}

func (o *UpdateMTOShipmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MTOShipment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentBadRequest creates a UpdateMTOShipmentBadRequest with default headers values
func NewUpdateMTOShipmentBadRequest() *UpdateMTOShipmentBadRequest {
	return &UpdateMTOShipmentBadRequest{}
}

/*UpdateMTOShipmentBadRequest handles this case with default header values.

invalid request
*/
type UpdateMTOShipmentBadRequest struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMTOShipmentBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentUnauthorized creates a UpdateMTOShipmentUnauthorized with default headers values
func NewUpdateMTOShipmentUnauthorized() *UpdateMTOShipmentUnauthorized {
	return &UpdateMTOShipmentUnauthorized{}
}

/*UpdateMTOShipmentUnauthorized handles this case with default header values.

The request was denied
*/
type UpdateMTOShipmentUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOShipmentUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentForbidden creates a UpdateMTOShipmentForbidden with default headers values
func NewUpdateMTOShipmentForbidden() *UpdateMTOShipmentForbidden {
	return &UpdateMTOShipmentForbidden{}
}

/*UpdateMTOShipmentForbidden handles this case with default header values.

The request was denied
*/
type UpdateMTOShipmentForbidden struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentForbidden) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOShipmentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentNotFound creates a UpdateMTOShipmentNotFound with default headers values
func NewUpdateMTOShipmentNotFound() *UpdateMTOShipmentNotFound {
	return &UpdateMTOShipmentNotFound{}
}

/*UpdateMTOShipmentNotFound handles this case with default header values.

The requested resource wasn't found
*/
type UpdateMTOShipmentNotFound struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentNotFound) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOShipmentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentPreconditionFailed creates a UpdateMTOShipmentPreconditionFailed with default headers values
func NewUpdateMTOShipmentPreconditionFailed() *UpdateMTOShipmentPreconditionFailed {
	return &UpdateMTOShipmentPreconditionFailed{}
}

/*UpdateMTOShipmentPreconditionFailed handles this case with default header values.

precondition failed
*/
type UpdateMTOShipmentPreconditionFailed struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOShipmentPreconditionFailed) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOShipmentInternalServerError creates a UpdateMTOShipmentInternalServerError with default headers values
func NewUpdateMTOShipmentInternalServerError() *UpdateMTOShipmentInternalServerError {
	return &UpdateMTOShipmentInternalServerError{}
}

/*UpdateMTOShipmentInternalServerError handles this case with default header values.

internal server error
*/
type UpdateMTOShipmentInternalServerError struct {
	Payload interface{}
}

func (o *UpdateMTOShipmentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /move-task-orders/{moveTaskOrderID}/mto-shipments/{mtoShipmentID}][%d] updateMTOShipmentInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOShipmentInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOShipmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
