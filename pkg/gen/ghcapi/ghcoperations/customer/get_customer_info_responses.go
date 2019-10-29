// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// GetCustomerInfoOKCode is the HTTP code returned for type GetCustomerInfoOK
const GetCustomerInfoOKCode int = 200

/*GetCustomerInfoOK Successfully retrieved information on an individual customer

swagger:response getCustomerInfoOK
*/
type GetCustomerInfoOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Customer `json:"body,omitempty"`
}

// NewGetCustomerInfoOK creates GetCustomerInfoOK with default headers values
func NewGetCustomerInfoOK() *GetCustomerInfoOK {

	return &GetCustomerInfoOK{}
}

// WithPayload adds the payload to the get customer info o k response
func (o *GetCustomerInfoOK) WithPayload(payload *ghcmessages.Customer) *GetCustomerInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info o k response
func (o *GetCustomerInfoOK) SetPayload(payload *ghcmessages.Customer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCustomerInfoBadRequestCode is the HTTP code returned for type GetCustomerInfoBadRequest
const GetCustomerInfoBadRequestCode int = 400

/*GetCustomerInfoBadRequest The request payload is invalid

swagger:response getCustomerInfoBadRequest
*/
type GetCustomerInfoBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetCustomerInfoBadRequest creates GetCustomerInfoBadRequest with default headers values
func NewGetCustomerInfoBadRequest() *GetCustomerInfoBadRequest {

	return &GetCustomerInfoBadRequest{}
}

// WithPayload adds the payload to the get customer info bad request response
func (o *GetCustomerInfoBadRequest) WithPayload(payload interface{}) *GetCustomerInfoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info bad request response
func (o *GetCustomerInfoBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCustomerInfoUnauthorizedCode is the HTTP code returned for type GetCustomerInfoUnauthorized
const GetCustomerInfoUnauthorizedCode int = 401

/*GetCustomerInfoUnauthorized The request was denied

swagger:response getCustomerInfoUnauthorized
*/
type GetCustomerInfoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetCustomerInfoUnauthorized creates GetCustomerInfoUnauthorized with default headers values
func NewGetCustomerInfoUnauthorized() *GetCustomerInfoUnauthorized {

	return &GetCustomerInfoUnauthorized{}
}

// WithPayload adds the payload to the get customer info unauthorized response
func (o *GetCustomerInfoUnauthorized) WithPayload(payload interface{}) *GetCustomerInfoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info unauthorized response
func (o *GetCustomerInfoUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCustomerInfoForbiddenCode is the HTTP code returned for type GetCustomerInfoForbidden
const GetCustomerInfoForbiddenCode int = 403

/*GetCustomerInfoForbidden The request was denied

swagger:response getCustomerInfoForbidden
*/
type GetCustomerInfoForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetCustomerInfoForbidden creates GetCustomerInfoForbidden with default headers values
func NewGetCustomerInfoForbidden() *GetCustomerInfoForbidden {

	return &GetCustomerInfoForbidden{}
}

// WithPayload adds the payload to the get customer info forbidden response
func (o *GetCustomerInfoForbidden) WithPayload(payload interface{}) *GetCustomerInfoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info forbidden response
func (o *GetCustomerInfoForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCustomerInfoNotFoundCode is the HTTP code returned for type GetCustomerInfoNotFound
const GetCustomerInfoNotFoundCode int = 404

/*GetCustomerInfoNotFound The requested resource wasn't found

swagger:response getCustomerInfoNotFound
*/
type GetCustomerInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetCustomerInfoNotFound creates GetCustomerInfoNotFound with default headers values
func NewGetCustomerInfoNotFound() *GetCustomerInfoNotFound {

	return &GetCustomerInfoNotFound{}
}

// WithPayload adds the payload to the get customer info not found response
func (o *GetCustomerInfoNotFound) WithPayload(payload interface{}) *GetCustomerInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info not found response
func (o *GetCustomerInfoNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCustomerInfoInternalServerErrorCode is the HTTP code returned for type GetCustomerInfoInternalServerError
const GetCustomerInfoInternalServerErrorCode int = 500

/*GetCustomerInfoInternalServerError A server error occurred

swagger:response getCustomerInfoInternalServerError
*/
type GetCustomerInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetCustomerInfoInternalServerError creates GetCustomerInfoInternalServerError with default headers values
func NewGetCustomerInfoInternalServerError() *GetCustomerInfoInternalServerError {

	return &GetCustomerInfoInternalServerError{}
}

// WithPayload adds the payload to the get customer info internal server error response
func (o *GetCustomerInfoInternalServerError) WithPayload(payload interface{}) *GetCustomerInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get customer info internal server error response
func (o *GetCustomerInfoInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCustomerInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}