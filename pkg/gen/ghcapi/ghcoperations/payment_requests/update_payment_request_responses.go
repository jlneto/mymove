// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdatePaymentRequestOKCode is the HTTP code returned for type UpdatePaymentRequestOK
const UpdatePaymentRequestOKCode int = 200

/*UpdatePaymentRequestOK updated payment request

swagger:response updatePaymentRequestOK
*/
type UpdatePaymentRequestOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.PaymentRequest `json:"body,omitempty"`
}

// NewUpdatePaymentRequestOK creates UpdatePaymentRequestOK with default headers values
func NewUpdatePaymentRequestOK() *UpdatePaymentRequestOK {

	return &UpdatePaymentRequestOK{}
}

// WithPayload adds the payload to the update payment request o k response
func (o *UpdatePaymentRequestOK) WithPayload(payload *ghcmessages.PaymentRequest) *UpdatePaymentRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request o k response
func (o *UpdatePaymentRequestOK) SetPayload(payload *ghcmessages.PaymentRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestBadRequestCode is the HTTP code returned for type UpdatePaymentRequestBadRequest
const UpdatePaymentRequestBadRequestCode int = 400

/*UpdatePaymentRequestBadRequest The request payload is invalid

swagger:response updatePaymentRequestBadRequest
*/
type UpdatePaymentRequestBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestBadRequest creates UpdatePaymentRequestBadRequest with default headers values
func NewUpdatePaymentRequestBadRequest() *UpdatePaymentRequestBadRequest {

	return &UpdatePaymentRequestBadRequest{}
}

// WithPayload adds the payload to the update payment request bad request response
func (o *UpdatePaymentRequestBadRequest) WithPayload(payload interface{}) *UpdatePaymentRequestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request bad request response
func (o *UpdatePaymentRequestBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestUnauthorizedCode is the HTTP code returned for type UpdatePaymentRequestUnauthorized
const UpdatePaymentRequestUnauthorizedCode int = 401

/*UpdatePaymentRequestUnauthorized The request was denied

swagger:response updatePaymentRequestUnauthorized
*/
type UpdatePaymentRequestUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestUnauthorized creates UpdatePaymentRequestUnauthorized with default headers values
func NewUpdatePaymentRequestUnauthorized() *UpdatePaymentRequestUnauthorized {

	return &UpdatePaymentRequestUnauthorized{}
}

// WithPayload adds the payload to the update payment request unauthorized response
func (o *UpdatePaymentRequestUnauthorized) WithPayload(payload interface{}) *UpdatePaymentRequestUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request unauthorized response
func (o *UpdatePaymentRequestUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestForbiddenCode is the HTTP code returned for type UpdatePaymentRequestForbidden
const UpdatePaymentRequestForbiddenCode int = 403

/*UpdatePaymentRequestForbidden The request was denied

swagger:response updatePaymentRequestForbidden
*/
type UpdatePaymentRequestForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestForbidden creates UpdatePaymentRequestForbidden with default headers values
func NewUpdatePaymentRequestForbidden() *UpdatePaymentRequestForbidden {

	return &UpdatePaymentRequestForbidden{}
}

// WithPayload adds the payload to the update payment request forbidden response
func (o *UpdatePaymentRequestForbidden) WithPayload(payload interface{}) *UpdatePaymentRequestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request forbidden response
func (o *UpdatePaymentRequestForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestNotFoundCode is the HTTP code returned for type UpdatePaymentRequestNotFound
const UpdatePaymentRequestNotFoundCode int = 404

/*UpdatePaymentRequestNotFound The requested resource wasn't found

swagger:response updatePaymentRequestNotFound
*/
type UpdatePaymentRequestNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestNotFound creates UpdatePaymentRequestNotFound with default headers values
func NewUpdatePaymentRequestNotFound() *UpdatePaymentRequestNotFound {

	return &UpdatePaymentRequestNotFound{}
}

// WithPayload adds the payload to the update payment request not found response
func (o *UpdatePaymentRequestNotFound) WithPayload(payload interface{}) *UpdatePaymentRequestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request not found response
func (o *UpdatePaymentRequestNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestInternalServerErrorCode is the HTTP code returned for type UpdatePaymentRequestInternalServerError
const UpdatePaymentRequestInternalServerErrorCode int = 500

/*UpdatePaymentRequestInternalServerError A server error occurred

swagger:response updatePaymentRequestInternalServerError
*/
type UpdatePaymentRequestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestInternalServerError creates UpdatePaymentRequestInternalServerError with default headers values
func NewUpdatePaymentRequestInternalServerError() *UpdatePaymentRequestInternalServerError {

	return &UpdatePaymentRequestInternalServerError{}
}

// WithPayload adds the payload to the update payment request internal server error response
func (o *UpdatePaymentRequestInternalServerError) WithPayload(payload interface{}) *UpdatePaymentRequestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request internal server error response
func (o *UpdatePaymentRequestInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
