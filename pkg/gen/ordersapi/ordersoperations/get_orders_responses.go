// Code generated by go-swagger; DO NOT EDIT.

package ordersoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ordersmessages "github.com/transcom/mymove/pkg/gen/ordersmessages"
)

// GetOrdersOKCode is the HTTP code returned for type GetOrdersOK
const GetOrdersOKCode int = 200

/*GetOrdersOK Successful

swagger:response getOrdersOK
*/
type GetOrdersOK struct {

	/*
	  In: Body
	*/
	Payload *ordersmessages.Orders `json:"body,omitempty"`
}

// NewGetOrdersOK creates GetOrdersOK with default headers values
func NewGetOrdersOK() *GetOrdersOK {

	return &GetOrdersOK{}
}

// WithPayload adds the payload to the get orders o k response
func (o *GetOrdersOK) WithPayload(payload *ordersmessages.Orders) *GetOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get orders o k response
func (o *GetOrdersOK) SetPayload(payload *ordersmessages.Orders) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrdersBadRequestCode is the HTTP code returned for type GetOrdersBadRequest
const GetOrdersBadRequestCode int = 400

/*GetOrdersBadRequest Invalid

swagger:response getOrdersBadRequest
*/
type GetOrdersBadRequest struct {
}

// NewGetOrdersBadRequest creates GetOrdersBadRequest with default headers values
func NewGetOrdersBadRequest() *GetOrdersBadRequest {

	return &GetOrdersBadRequest{}
}

// WriteResponse to the client
func (o *GetOrdersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetOrdersUnauthorizedCode is the HTTP code returned for type GetOrdersUnauthorized
const GetOrdersUnauthorizedCode int = 401

/*GetOrdersUnauthorized must be authenticated to use this endpoint

swagger:response getOrdersUnauthorized
*/
type GetOrdersUnauthorized struct {
}

// NewGetOrdersUnauthorized creates GetOrdersUnauthorized with default headers values
func NewGetOrdersUnauthorized() *GetOrdersUnauthorized {

	return &GetOrdersUnauthorized{}
}

// WriteResponse to the client
func (o *GetOrdersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetOrdersForbiddenCode is the HTTP code returned for type GetOrdersForbidden
const GetOrdersForbiddenCode int = 403

/*GetOrdersForbidden Forbidden

swagger:response getOrdersForbidden
*/
type GetOrdersForbidden struct {
}

// NewGetOrdersForbidden creates GetOrdersForbidden with default headers values
func NewGetOrdersForbidden() *GetOrdersForbidden {

	return &GetOrdersForbidden{}
}

// WriteResponse to the client
func (o *GetOrdersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetOrdersNotFoundCode is the HTTP code returned for type GetOrdersNotFound
const GetOrdersNotFoundCode int = 404

/*GetOrdersNotFound Orders not found

swagger:response getOrdersNotFound
*/
type GetOrdersNotFound struct {
}

// NewGetOrdersNotFound creates GetOrdersNotFound with default headers values
func NewGetOrdersNotFound() *GetOrdersNotFound {

	return &GetOrdersNotFound{}
}

// WriteResponse to the client
func (o *GetOrdersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetOrdersInternalServerErrorCode is the HTTP code returned for type GetOrdersInternalServerError
const GetOrdersInternalServerErrorCode int = 500

/*GetOrdersInternalServerError Server error

swagger:response getOrdersInternalServerError
*/
type GetOrdersInternalServerError struct {
}

// NewGetOrdersInternalServerError creates GetOrdersInternalServerError with default headers values
func NewGetOrdersInternalServerError() *GetOrdersInternalServerError {

	return &GetOrdersInternalServerError{}
}

// WriteResponse to the client
func (o *GetOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
