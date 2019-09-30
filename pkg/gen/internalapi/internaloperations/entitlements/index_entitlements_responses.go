// Code generated by go-swagger; DO NOT EDIT.

package entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// IndexEntitlementsOKCode is the HTTP code returned for type IndexEntitlementsOK
const IndexEntitlementsOKCode int = 200

/*IndexEntitlementsOK List of weights allotted entitlement

swagger:response indexEntitlementsOK
*/
type IndexEntitlementsOK struct {

	/*
	  In: Body
	*/
	Payload internalmessages.IndexEntitlements `json:"body,omitempty"`
}

// NewIndexEntitlementsOK creates IndexEntitlementsOK with default headers values
func NewIndexEntitlementsOK() *IndexEntitlementsOK {

	return &IndexEntitlementsOK{}
}

// WithPayload adds the payload to the index entitlements o k response
func (o *IndexEntitlementsOK) WithPayload(payload internalmessages.IndexEntitlements) *IndexEntitlementsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index entitlements o k response
func (o *IndexEntitlementsOK) SetPayload(payload internalmessages.IndexEntitlements) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexEntitlementsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = internalmessages.IndexEntitlements{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
