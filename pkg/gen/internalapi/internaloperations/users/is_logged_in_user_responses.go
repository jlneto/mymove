// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// IsLoggedInUserOKCode is the HTTP code returned for type IsLoggedInUserOK
const IsLoggedInUserOKCode int = 200

/*IsLoggedInUserOK Currently logged in user

swagger:response isLoggedInUserOK
*/
type IsLoggedInUserOK struct {

	/*
	  In: Body
	*/
	Payload bool `json:"body,omitempty"`
}

// NewIsLoggedInUserOK creates IsLoggedInUserOK with default headers values
func NewIsLoggedInUserOK() *IsLoggedInUserOK {

	return &IsLoggedInUserOK{}
}

// WithPayload adds the payload to the is logged in user o k response
func (o *IsLoggedInUserOK) WithPayload(payload bool) *IsLoggedInUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is logged in user o k response
func (o *IsLoggedInUserOK) SetPayload(payload bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsLoggedInUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IsLoggedInUserBadRequestCode is the HTTP code returned for type IsLoggedInUserBadRequest
const IsLoggedInUserBadRequestCode int = 400

/*IsLoggedInUserBadRequest invalid request

swagger:response isLoggedInUserBadRequest
*/
type IsLoggedInUserBadRequest struct {
}

// NewIsLoggedInUserBadRequest creates IsLoggedInUserBadRequest with default headers values
func NewIsLoggedInUserBadRequest() *IsLoggedInUserBadRequest {

	return &IsLoggedInUserBadRequest{}
}

// WriteResponse to the client
func (o *IsLoggedInUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IsLoggedInUserInternalServerErrorCode is the HTTP code returned for type IsLoggedInUserInternalServerError
const IsLoggedInUserInternalServerErrorCode int = 500

/*IsLoggedInUserInternalServerError server error

swagger:response isLoggedInUserInternalServerError
*/
type IsLoggedInUserInternalServerError struct {
}

// NewIsLoggedInUserInternalServerError creates IsLoggedInUserInternalServerError with default headers values
func NewIsLoggedInUserInternalServerError() *IsLoggedInUserInternalServerError {

	return &IsLoggedInUserInternalServerError{}
}

// WriteResponse to the client
func (o *IsLoggedInUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}