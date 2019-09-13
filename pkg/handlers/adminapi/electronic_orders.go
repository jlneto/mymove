package adminapi

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"

	electronicorderop "github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/electronic_order"
	"github.com/transcom/mymove/pkg/gen/adminmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

func payloadForElectronicOrderModel(o models.ElectronicOrder) *adminmessages.ElectronicOrder {
	return &adminmessages.ElectronicOrder{
		ID:        handlers.FmtUUID(o.ID),
		Issuer:    adminmessages.Issuer(o.Issuer),
		CreatedAt: handlers.FmtDateTime(o.CreatedAt),
		UpdatedAt: handlers.FmtDateTime(o.UpdatedAt),
	}
}

type IndexElectronicOrdersHandler struct {
	handlers.HandlerContext
	services.ElectronicOrderListFetcher
	services.NewQueryFilter
	services.NewPagination
}

func (h IndexElectronicOrdersHandler) Handle(params electronicorderop.IndexElectronicOrdersParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	queryFilters := []services.QueryFilter{}

	pagination := h.NewPagination(params.Page, params.PerPage)

	electronicOrders, err := h.ElectronicOrderListFetcher.FetchElectronicOrderList(queryFilters, pagination)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	totalElectronicOrdersCount, err := h.DB().Count(&models.ElectronicOrder{})
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	queriedOfficeUsersCount := len(electronicOrders)

	payload := make(adminmessages.ElectronicOrders, queriedOfficeUsersCount)
	for i, s := range electronicOrders {
		payload[i] = payloadForElectronicOrderModel(s)
	}

	return electronicorderop.NewIndexElectronicOrdersOK().WithContentRange(fmt.Sprintf("electronic_orders %d-%d/%d", pagination.Offset(), pagination.Offset()+queriedOfficeUsersCount, totalElectronicOrdersCount)).WithPayload(payload)
}

type GetElectronicOrdersTotalsHandler struct {
	handlers.HandlerContext
	services.ElectronicOrderCategoryCountFetcher
	services.NewQueryFilter
}

func split(r rune) bool {
	return r == '.' || r == ':'
}

func translateComparator(s string) string {
	s = strings.ToLower(s)
	switch s {
	case "eq":
		return "="
	case "gt":
		return ">"
	case "lt":
		return "<"
	case "neq":
		return "!="
	case "lte":
		return "<="
	case "gte":
		return ">="
	}
	return s
}

func (h GetElectronicOrdersTotalsHandler) Handle(params electronicorderop.GetElectronicOrdersTotalsParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	comparator := ""

	queryFilters := []services.QueryFilter{}
	andQueryFilters := []services.QueryFilter{}

	for _, filter := range params.Filter {
		queryFilterSplit := strings.FieldsFunc(filter, split)
		comparator = translateComparator(queryFilterSplit[1])
		queryFilters = append(queryFilters, h.NewQueryFilter(queryFilterSplit[0], comparator, queryFilterSplit[2]))
	}

	if params.AndFilter != nil {
		for _, andFilter := range params.AndFilter {
			andFilterSplit := strings.FieldsFunc(andFilter, split)
			comparator = translateComparator(andFilterSplit[1])
			andQueryFilters = append(andQueryFilters, h.NewQueryFilter(andFilterSplit[0], comparator, andFilterSplit[2]))
		}
	}
	counts, err := h.ElectronicOrderCategoryCountFetcher.FetchElectronicOrderCategoricalCounts(queryFilters, &andQueryFilters)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}
	payload := adminmessages.ElectronicOrdersTotals{}
	for key, count := range counts {
		count64 := int64(count)
		stringKey := fmt.Sprintf("%v", key)
		totalCount := adminmessages.ElectronicOrdersTotal{
			Category: stringKey,
			Count:    &count64,
		}
		payload = append(payload, &totalCount)
	}

	return electronicorderop.NewGetElectronicOrdersTotalsOK().WithPayload(payload)
}
