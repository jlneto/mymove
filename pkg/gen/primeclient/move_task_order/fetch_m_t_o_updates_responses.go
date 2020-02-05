// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// FetchMTOUpdatesReader is a Reader for the FetchMTOUpdates structure.
type FetchMTOUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchMTOUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFetchMTOUpdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFetchMTOUpdatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFetchMTOUpdatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFetchMTOUpdatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFetchMTOUpdatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFetchMTOUpdatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFetchMTOUpdatesOK creates a FetchMTOUpdatesOK with default headers values
func NewFetchMTOUpdatesOK() *FetchMTOUpdatesOK {
	return &FetchMTOUpdatesOK{}
}

/*FetchMTOUpdatesOK handles this case with default header values.

Successfully retrieved all move task orders
*/
type FetchMTOUpdatesOK struct {
	Payload primemessages.MoveTaskOrders
}

func (o *FetchMTOUpdatesOK) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesOK  %+v", 200, o.Payload)
}

func (o *FetchMTOUpdatesOK) GetPayload() primemessages.MoveTaskOrders {
	return o.Payload
}

func (o *FetchMTOUpdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchMTOUpdatesBadRequest creates a FetchMTOUpdatesBadRequest with default headers values
func NewFetchMTOUpdatesBadRequest() *FetchMTOUpdatesBadRequest {
	return &FetchMTOUpdatesBadRequest{}
}

/*FetchMTOUpdatesBadRequest handles this case with default header values.

The request payload is invalid
*/
type FetchMTOUpdatesBadRequest struct {
	Payload interface{}
}

func (o *FetchMTOUpdatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesBadRequest  %+v", 400, o.Payload)
}

func (o *FetchMTOUpdatesBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *FetchMTOUpdatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchMTOUpdatesUnauthorized creates a FetchMTOUpdatesUnauthorized with default headers values
func NewFetchMTOUpdatesUnauthorized() *FetchMTOUpdatesUnauthorized {
	return &FetchMTOUpdatesUnauthorized{}
}

/*FetchMTOUpdatesUnauthorized handles this case with default header values.

The request was denied
*/
type FetchMTOUpdatesUnauthorized struct {
	Payload interface{}
}

func (o *FetchMTOUpdatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesUnauthorized  %+v", 401, o.Payload)
}

func (o *FetchMTOUpdatesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *FetchMTOUpdatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchMTOUpdatesForbidden creates a FetchMTOUpdatesForbidden with default headers values
func NewFetchMTOUpdatesForbidden() *FetchMTOUpdatesForbidden {
	return &FetchMTOUpdatesForbidden{}
}

/*FetchMTOUpdatesForbidden handles this case with default header values.

The request was denied
*/
type FetchMTOUpdatesForbidden struct {
	Payload interface{}
}

func (o *FetchMTOUpdatesForbidden) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesForbidden  %+v", 403, o.Payload)
}

func (o *FetchMTOUpdatesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *FetchMTOUpdatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchMTOUpdatesNotFound creates a FetchMTOUpdatesNotFound with default headers values
func NewFetchMTOUpdatesNotFound() *FetchMTOUpdatesNotFound {
	return &FetchMTOUpdatesNotFound{}
}

/*FetchMTOUpdatesNotFound handles this case with default header values.

The requested resource wasn't found
*/
type FetchMTOUpdatesNotFound struct {
	Payload interface{}
}

func (o *FetchMTOUpdatesNotFound) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesNotFound  %+v", 404, o.Payload)
}

func (o *FetchMTOUpdatesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *FetchMTOUpdatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchMTOUpdatesInternalServerError creates a FetchMTOUpdatesInternalServerError with default headers values
func NewFetchMTOUpdatesInternalServerError() *FetchMTOUpdatesInternalServerError {
	return &FetchMTOUpdatesInternalServerError{}
}

/*FetchMTOUpdatesInternalServerError handles this case with default header values.

A server error occurred
*/
type FetchMTOUpdatesInternalServerError struct {
	Payload interface{}
}

func (o *FetchMTOUpdatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /move-task-orders][%d] fetchMTOUpdatesInternalServerError  %+v", 500, o.Payload)
}

func (o *FetchMTOUpdatesInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *FetchMTOUpdatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
