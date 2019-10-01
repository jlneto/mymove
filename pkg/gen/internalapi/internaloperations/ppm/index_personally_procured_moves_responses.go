// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// IndexPersonallyProcuredMovesOKCode is the HTTP code returned for type IndexPersonallyProcuredMovesOK
const IndexPersonallyProcuredMovesOKCode int = 200

/*IndexPersonallyProcuredMovesOK returns list of personally_procured_move

swagger:response indexPersonallyProcuredMovesOK
*/
type IndexPersonallyProcuredMovesOK struct {

	/*
	  In: Body
	*/
	Payload internalmessages.IndexPersonallyProcuredMovePayload `json:"body,omitempty"`
}

// NewIndexPersonallyProcuredMovesOK creates IndexPersonallyProcuredMovesOK with default headers values
func NewIndexPersonallyProcuredMovesOK() *IndexPersonallyProcuredMovesOK {

	return &IndexPersonallyProcuredMovesOK{}
}

// WithPayload adds the payload to the index personally procured moves o k response
func (o *IndexPersonallyProcuredMovesOK) WithPayload(payload internalmessages.IndexPersonallyProcuredMovePayload) *IndexPersonallyProcuredMovesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index personally procured moves o k response
func (o *IndexPersonallyProcuredMovesOK) SetPayload(payload internalmessages.IndexPersonallyProcuredMovePayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexPersonallyProcuredMovesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = internalmessages.IndexPersonallyProcuredMovePayload{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexPersonallyProcuredMovesBadRequestCode is the HTTP code returned for type IndexPersonallyProcuredMovesBadRequest
const IndexPersonallyProcuredMovesBadRequestCode int = 400

/*IndexPersonallyProcuredMovesBadRequest invalid request

swagger:response indexPersonallyProcuredMovesBadRequest
*/
type IndexPersonallyProcuredMovesBadRequest struct {
}

// NewIndexPersonallyProcuredMovesBadRequest creates IndexPersonallyProcuredMovesBadRequest with default headers values
func NewIndexPersonallyProcuredMovesBadRequest() *IndexPersonallyProcuredMovesBadRequest {

	return &IndexPersonallyProcuredMovesBadRequest{}
}

// WriteResponse to the client
func (o *IndexPersonallyProcuredMovesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexPersonallyProcuredMovesUnauthorizedCode is the HTTP code returned for type IndexPersonallyProcuredMovesUnauthorized
const IndexPersonallyProcuredMovesUnauthorizedCode int = 401

/*IndexPersonallyProcuredMovesUnauthorized request requires user authentication

swagger:response indexPersonallyProcuredMovesUnauthorized
*/
type IndexPersonallyProcuredMovesUnauthorized struct {
}

// NewIndexPersonallyProcuredMovesUnauthorized creates IndexPersonallyProcuredMovesUnauthorized with default headers values
func NewIndexPersonallyProcuredMovesUnauthorized() *IndexPersonallyProcuredMovesUnauthorized {

	return &IndexPersonallyProcuredMovesUnauthorized{}
}

// WriteResponse to the client
func (o *IndexPersonallyProcuredMovesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexPersonallyProcuredMovesForbiddenCode is the HTTP code returned for type IndexPersonallyProcuredMovesForbidden
const IndexPersonallyProcuredMovesForbiddenCode int = 403

/*IndexPersonallyProcuredMovesForbidden user is not authorized

swagger:response indexPersonallyProcuredMovesForbidden
*/
type IndexPersonallyProcuredMovesForbidden struct {
}

// NewIndexPersonallyProcuredMovesForbidden creates IndexPersonallyProcuredMovesForbidden with default headers values
func NewIndexPersonallyProcuredMovesForbidden() *IndexPersonallyProcuredMovesForbidden {

	return &IndexPersonallyProcuredMovesForbidden{}
}

// WriteResponse to the client
func (o *IndexPersonallyProcuredMovesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}