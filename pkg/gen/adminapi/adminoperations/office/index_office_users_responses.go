// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"
)

// IndexOfficeUsersOKCode is the HTTP code returned for type IndexOfficeUsersOK
const IndexOfficeUsersOKCode int = 200

/*IndexOfficeUsersOK success

swagger:response indexOfficeUsersOK
*/
type IndexOfficeUsersOK struct {
	/*Used for pagination

	 */
	ContentRange string `json:"Content-Range"`

	/*
	  In: Body
	*/
	Payload adminmessages.OfficeUsers `json:"body,omitempty"`
}

// NewIndexOfficeUsersOK creates IndexOfficeUsersOK with default headers values
func NewIndexOfficeUsersOK() *IndexOfficeUsersOK {

	return &IndexOfficeUsersOK{}
}

// WithContentRange adds the contentRange to the index office users o k response
func (o *IndexOfficeUsersOK) WithContentRange(contentRange string) *IndexOfficeUsersOK {
	o.ContentRange = contentRange
	return o
}

// SetContentRange sets the contentRange to the index office users o k response
func (o *IndexOfficeUsersOK) SetContentRange(contentRange string) {
	o.ContentRange = contentRange
}

// WithPayload adds the payload to the index office users o k response
func (o *IndexOfficeUsersOK) WithPayload(payload adminmessages.OfficeUsers) *IndexOfficeUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index office users o k response
func (o *IndexOfficeUsersOK) SetPayload(payload adminmessages.OfficeUsers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexOfficeUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Range

	contentRange := o.ContentRange
	if contentRange != "" {
		rw.Header().Set("Content-Range", contentRange)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = adminmessages.OfficeUsers{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexOfficeUsersBadRequestCode is the HTTP code returned for type IndexOfficeUsersBadRequest
const IndexOfficeUsersBadRequestCode int = 400

/*IndexOfficeUsersBadRequest invalid request

swagger:response indexOfficeUsersBadRequest
*/
type IndexOfficeUsersBadRequest struct {
}

// NewIndexOfficeUsersBadRequest creates IndexOfficeUsersBadRequest with default headers values
func NewIndexOfficeUsersBadRequest() *IndexOfficeUsersBadRequest {

	return &IndexOfficeUsersBadRequest{}
}

// WriteResponse to the client
func (o *IndexOfficeUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexOfficeUsersUnauthorizedCode is the HTTP code returned for type IndexOfficeUsersUnauthorized
const IndexOfficeUsersUnauthorizedCode int = 401

/*IndexOfficeUsersUnauthorized request requires user authentication

swagger:response indexOfficeUsersUnauthorized
*/
type IndexOfficeUsersUnauthorized struct {
}

// NewIndexOfficeUsersUnauthorized creates IndexOfficeUsersUnauthorized with default headers values
func NewIndexOfficeUsersUnauthorized() *IndexOfficeUsersUnauthorized {

	return &IndexOfficeUsersUnauthorized{}
}

// WriteResponse to the client
func (o *IndexOfficeUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexOfficeUsersNotFoundCode is the HTTP code returned for type IndexOfficeUsersNotFound
const IndexOfficeUsersNotFoundCode int = 404

/*IndexOfficeUsersNotFound office not found

swagger:response indexOfficeUsersNotFound
*/
type IndexOfficeUsersNotFound struct {
}

// NewIndexOfficeUsersNotFound creates IndexOfficeUsersNotFound with default headers values
func NewIndexOfficeUsersNotFound() *IndexOfficeUsersNotFound {

	return &IndexOfficeUsersNotFound{}
}

// WriteResponse to the client
func (o *IndexOfficeUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndexOfficeUsersInternalServerErrorCode is the HTTP code returned for type IndexOfficeUsersInternalServerError
const IndexOfficeUsersInternalServerErrorCode int = 500

/*IndexOfficeUsersInternalServerError server error

swagger:response indexOfficeUsersInternalServerError
*/
type IndexOfficeUsersInternalServerError struct {
}

// NewIndexOfficeUsersInternalServerError creates IndexOfficeUsersInternalServerError with default headers values
func NewIndexOfficeUsersInternalServerError() *IndexOfficeUsersInternalServerError {

	return &IndexOfficeUsersInternalServerError{}
}

// WriteResponse to the client
func (o *IndexOfficeUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
