// Code generated by go-swagger; DO NOT EDIT.

package ghcoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/customer"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/entitlements"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/payment_requests"
	"github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/service_item"
)

// NewMymoveAPI creates a new Mymove instance
func NewMymoveAPI(spec *loads.Document) *MymoveAPI {
	return &MymoveAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		PaymentRequestsCreatePaymentRequestHandler: payment_requests.CreatePaymentRequestHandlerFunc(func(params payment_requests.CreatePaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsCreatePaymentRequest has not yet been implemented")
		}),
		ServiceItemCreateServiceItemHandler: service_item.CreateServiceItemHandlerFunc(func(params service_item.CreateServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemCreateServiceItem has not yet been implemented")
		}),
		MoveTaskOrderDeleteMoveTaskOrderHandler: move_task_order.DeleteMoveTaskOrderHandlerFunc(func(params move_task_order.DeleteMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderDeleteMoveTaskOrder has not yet been implemented")
		}),
		ServiceItemDeleteServiceItemHandler: service_item.DeleteServiceItemHandlerFunc(func(params service_item.DeleteServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemDeleteServiceItem has not yet been implemented")
		}),
		CustomerGetCustomerInfoHandler: customer.GetCustomerInfoHandlerFunc(func(params customer.GetCustomerInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CustomerGetCustomerInfo has not yet been implemented")
		}),
		EntitlementsGetEntitlementsHandler: entitlements.GetEntitlementsHandlerFunc(func(params entitlements.GetEntitlementsParams) middleware.Responder {
			return middleware.NotImplemented("operation EntitlementsGetEntitlements has not yet been implemented")
		}),
		MoveTaskOrderGetMoveTaskOrderHandler: move_task_order.GetMoveTaskOrderHandlerFunc(func(params move_task_order.GetMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderGetMoveTaskOrder has not yet been implemented")
		}),
		PaymentRequestsGetPaymentRequestHandler: payment_requests.GetPaymentRequestHandlerFunc(func(params payment_requests.GetPaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsGetPaymentRequest has not yet been implemented")
		}),
		ServiceItemGetServiceItemHandler: service_item.GetServiceItemHandlerFunc(func(params service_item.GetServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemGetServiceItem has not yet been implemented")
		}),
		MoveTaskOrderListMoveTaskOrdersHandler: move_task_order.ListMoveTaskOrdersHandlerFunc(func(params move_task_order.ListMoveTaskOrdersParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderListMoveTaskOrders has not yet been implemented")
		}),
		PaymentRequestsListPaymentRequestsHandler: payment_requests.ListPaymentRequestsHandlerFunc(func(params payment_requests.ListPaymentRequestsParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsListPaymentRequests has not yet been implemented")
		}),
		ServiceItemListServiceItemsHandler: service_item.ListServiceItemsHandlerFunc(func(params service_item.ListServiceItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemListServiceItems has not yet been implemented")
		}),
		MoveTaskOrderUpdateMoveTaskOrderHandler: move_task_order.UpdateMoveTaskOrderHandlerFunc(func(params move_task_order.UpdateMoveTaskOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderUpdateMoveTaskOrder has not yet been implemented")
		}),
		MoveTaskOrderUpdateMoveTaskOrderStatusHandler: move_task_order.UpdateMoveTaskOrderStatusHandlerFunc(func(params move_task_order.UpdateMoveTaskOrderStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation MoveTaskOrderUpdateMoveTaskOrderStatus has not yet been implemented")
		}),
		PaymentRequestsUpdatePaymentRequestHandler: payment_requests.UpdatePaymentRequestHandlerFunc(func(params payment_requests.UpdatePaymentRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsUpdatePaymentRequest has not yet been implemented")
		}),
		PaymentRequestsUpdatePaymentRequestStatusHandler: payment_requests.UpdatePaymentRequestStatusHandlerFunc(func(params payment_requests.UpdatePaymentRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation PaymentRequestsUpdatePaymentRequestStatus has not yet been implemented")
		}),
		ServiceItemUpdateServiceItemHandler: service_item.UpdateServiceItemHandlerFunc(func(params service_item.UpdateServiceItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemUpdateServiceItem has not yet been implemented")
		}),
		ServiceItemUpdateServiceItemStatusHandler: service_item.UpdateServiceItemStatusHandlerFunc(func(params service_item.UpdateServiceItemStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation ServiceItemUpdateServiceItemStatus has not yet been implemented")
		}),
	}
}

/*MymoveAPI The API for move.mil */
type MymoveAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PaymentRequestsCreatePaymentRequestHandler sets the operation handler for the create payment request operation
	PaymentRequestsCreatePaymentRequestHandler payment_requests.CreatePaymentRequestHandler
	// ServiceItemCreateServiceItemHandler sets the operation handler for the create service item operation
	ServiceItemCreateServiceItemHandler service_item.CreateServiceItemHandler
	// MoveTaskOrderDeleteMoveTaskOrderHandler sets the operation handler for the delete move task order operation
	MoveTaskOrderDeleteMoveTaskOrderHandler move_task_order.DeleteMoveTaskOrderHandler
	// ServiceItemDeleteServiceItemHandler sets the operation handler for the delete service item operation
	ServiceItemDeleteServiceItemHandler service_item.DeleteServiceItemHandler
	// CustomerGetCustomerInfoHandler sets the operation handler for the get customer info operation
	CustomerGetCustomerInfoHandler customer.GetCustomerInfoHandler
	// EntitlementsGetEntitlementsHandler sets the operation handler for the get entitlements operation
	EntitlementsGetEntitlementsHandler entitlements.GetEntitlementsHandler
	// MoveTaskOrderGetMoveTaskOrderHandler sets the operation handler for the get move task order operation
	MoveTaskOrderGetMoveTaskOrderHandler move_task_order.GetMoveTaskOrderHandler
	// PaymentRequestsGetPaymentRequestHandler sets the operation handler for the get payment request operation
	PaymentRequestsGetPaymentRequestHandler payment_requests.GetPaymentRequestHandler
	// ServiceItemGetServiceItemHandler sets the operation handler for the get service item operation
	ServiceItemGetServiceItemHandler service_item.GetServiceItemHandler
	// MoveTaskOrderListMoveTaskOrdersHandler sets the operation handler for the list move task orders operation
	MoveTaskOrderListMoveTaskOrdersHandler move_task_order.ListMoveTaskOrdersHandler
	// PaymentRequestsListPaymentRequestsHandler sets the operation handler for the list payment requests operation
	PaymentRequestsListPaymentRequestsHandler payment_requests.ListPaymentRequestsHandler
	// ServiceItemListServiceItemsHandler sets the operation handler for the list service items operation
	ServiceItemListServiceItemsHandler service_item.ListServiceItemsHandler
	// MoveTaskOrderUpdateMoveTaskOrderHandler sets the operation handler for the update move task order operation
	MoveTaskOrderUpdateMoveTaskOrderHandler move_task_order.UpdateMoveTaskOrderHandler
	// MoveTaskOrderUpdateMoveTaskOrderStatusHandler sets the operation handler for the update move task order status operation
	MoveTaskOrderUpdateMoveTaskOrderStatusHandler move_task_order.UpdateMoveTaskOrderStatusHandler
	// PaymentRequestsUpdatePaymentRequestHandler sets the operation handler for the update payment request operation
	PaymentRequestsUpdatePaymentRequestHandler payment_requests.UpdatePaymentRequestHandler
	// PaymentRequestsUpdatePaymentRequestStatusHandler sets the operation handler for the update payment request status operation
	PaymentRequestsUpdatePaymentRequestStatusHandler payment_requests.UpdatePaymentRequestStatusHandler
	// ServiceItemUpdateServiceItemHandler sets the operation handler for the update service item operation
	ServiceItemUpdateServiceItemHandler service_item.UpdateServiceItemHandler
	// ServiceItemUpdateServiceItemStatusHandler sets the operation handler for the update service item status operation
	ServiceItemUpdateServiceItemStatusHandler service_item.UpdateServiceItemStatusHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MymoveAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MymoveAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MymoveAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MymoveAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MymoveAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MymoveAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MymoveAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MymoveAPI
func (o *MymoveAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PaymentRequestsCreatePaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.CreatePaymentRequestHandler")
	}

	if o.ServiceItemCreateServiceItemHandler == nil {
		unregistered = append(unregistered, "service_item.CreateServiceItemHandler")
	}

	if o.MoveTaskOrderDeleteMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.DeleteMoveTaskOrderHandler")
	}

	if o.ServiceItemDeleteServiceItemHandler == nil {
		unregistered = append(unregistered, "service_item.DeleteServiceItemHandler")
	}

	if o.CustomerGetCustomerInfoHandler == nil {
		unregistered = append(unregistered, "customer.GetCustomerInfoHandler")
	}

	if o.EntitlementsGetEntitlementsHandler == nil {
		unregistered = append(unregistered, "entitlements.GetEntitlementsHandler")
	}

	if o.MoveTaskOrderGetMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.GetMoveTaskOrderHandler")
	}

	if o.PaymentRequestsGetPaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.GetPaymentRequestHandler")
	}

	if o.ServiceItemGetServiceItemHandler == nil {
		unregistered = append(unregistered, "service_item.GetServiceItemHandler")
	}

	if o.MoveTaskOrderListMoveTaskOrdersHandler == nil {
		unregistered = append(unregistered, "move_task_order.ListMoveTaskOrdersHandler")
	}

	if o.PaymentRequestsListPaymentRequestsHandler == nil {
		unregistered = append(unregistered, "payment_requests.ListPaymentRequestsHandler")
	}

	if o.ServiceItemListServiceItemsHandler == nil {
		unregistered = append(unregistered, "service_item.ListServiceItemsHandler")
	}

	if o.MoveTaskOrderUpdateMoveTaskOrderHandler == nil {
		unregistered = append(unregistered, "move_task_order.UpdateMoveTaskOrderHandler")
	}

	if o.MoveTaskOrderUpdateMoveTaskOrderStatusHandler == nil {
		unregistered = append(unregistered, "move_task_order.UpdateMoveTaskOrderStatusHandler")
	}

	if o.PaymentRequestsUpdatePaymentRequestHandler == nil {
		unregistered = append(unregistered, "payment_requests.UpdatePaymentRequestHandler")
	}

	if o.PaymentRequestsUpdatePaymentRequestStatusHandler == nil {
		unregistered = append(unregistered, "payment_requests.UpdatePaymentRequestStatusHandler")
	}

	if o.ServiceItemUpdateServiceItemHandler == nil {
		unregistered = append(unregistered, "service_item.UpdateServiceItemHandler")
	}

	if o.ServiceItemUpdateServiceItemStatusHandler == nil {
		unregistered = append(unregistered, "service_item.UpdateServiceItemStatusHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MymoveAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MymoveAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *MymoveAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MymoveAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mymove API
func (o *MymoveAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MymoveAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/payment-requests"] = payment_requests.NewCreatePaymentRequest(o.context, o.PaymentRequestsCreatePaymentRequestHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/move-task-orders/{moveTaskOrderID}/service-items"] = service_item.NewCreateServiceItem(o.context, o.ServiceItemCreateServiceItemHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewDeleteMoveTaskOrder(o.context, o.MoveTaskOrderDeleteMoveTaskOrderHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/move-task-orders/{moveTaskOrderID}/service-items/{serviceItemID}"] = service_item.NewDeleteServiceItem(o.context, o.ServiceItemDeleteServiceItemHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/customer/{customerID}"] = customer.NewGetCustomerInfo(o.context, o.CustomerGetCustomerInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/entitlements"] = entitlements.NewGetEntitlements(o.context, o.EntitlementsGetEntitlementsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewGetMoveTaskOrder(o.context, o.MoveTaskOrderGetMoveTaskOrderHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/payment-requests/{paymentRequestID}"] = payment_requests.NewGetPaymentRequest(o.context, o.PaymentRequestsGetPaymentRequestHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/service-items/{serviceItemID}"] = service_item.NewGetServiceItem(o.context, o.ServiceItemGetServiceItemHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders"] = move_task_order.NewListMoveTaskOrders(o.context, o.MoveTaskOrderListMoveTaskOrdersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/payment-requests"] = payment_requests.NewListPaymentRequests(o.context, o.PaymentRequestsListPaymentRequestsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/move-task-orders/{moveTaskOrderID}/service-items"] = service_item.NewListServiceItems(o.context, o.ServiceItemListServiceItemsHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}"] = move_task_order.NewUpdateMoveTaskOrder(o.context, o.MoveTaskOrderUpdateMoveTaskOrderHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/status"] = move_task_order.NewUpdateMoveTaskOrderStatus(o.context, o.MoveTaskOrderUpdateMoveTaskOrderStatusHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/{paymentRequestID}"] = payment_requests.NewUpdatePaymentRequest(o.context, o.PaymentRequestsUpdatePaymentRequestHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/payment-requests/{paymentRequestID}/status"] = payment_requests.NewUpdatePaymentRequestStatus(o.context, o.PaymentRequestsUpdatePaymentRequestStatusHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/service-items/{serviceItemID}"] = service_item.NewUpdateServiceItem(o.context, o.ServiceItemUpdateServiceItemHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/move-task-orders/{moveTaskOrderID}/service-items/{serviceItemID}/status"] = service_item.NewUpdateServiceItemStatus(o.context, o.ServiceItemUpdateServiceItemStatusHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MymoveAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MymoveAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MymoveAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MymoveAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
