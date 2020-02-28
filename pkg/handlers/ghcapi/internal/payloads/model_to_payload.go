package payloads

import (
	"encoding/base64"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrder payload
func MoveTaskOrder(moveTaskOrder *models.MoveTaskOrder) *ghcmessages.MoveTaskOrder {
	if moveTaskOrder == nil {
		return nil
	}

	payload := &ghcmessages.MoveTaskOrder{
		ID:                 strfmt.UUID(moveTaskOrder.ID.String()),
		CreatedAt:          strfmt.Date(moveTaskOrder.CreatedAt),
		IsAvailableToPrime: &moveTaskOrder.IsAvailableToPrime,
		IsCanceled:         &moveTaskOrder.IsCanceled,
		MoveOrderID:        strfmt.UUID(moveTaskOrder.MoveOrderID.String()),
		ReferenceID:        moveTaskOrder.ReferenceID,
		UpdatedAt:          strfmt.Date(moveTaskOrder.UpdatedAt),
	}
	return payload
}

// Customer payload
func Customer(customer *models.Customer) *ghcmessages.Customer {
	if customer == nil {
		return nil
	}
	payload := ghcmessages.Customer{
		Agency:             customer.Agency,
		CurrentAddress:     Address(&customer.CurrentAddress),
		DestinationAddress: Address(&customer.DestinationAddress),
		DodID:              customer.DODID,
		Email:              customer.Email,
		FirstName:          customer.FirstName,
		ID:                 strfmt.UUID(customer.ID.String()),
		LastName:           customer.LastName,
		Phone:              customer.PhoneNumber,
		UserID:             strfmt.UUID(customer.UserID.String()),
	}
	return &payload
}

// MoveOrder payload
func MoveOrder(moveOrder *models.MoveOrder) *ghcmessages.MoveOrder {
	if moveOrder == nil {
		return nil
	}
	destinationDutyStation := DutyStation(moveOrder.DestinationDutyStation)
	originDutyStation := DutyStation(moveOrder.OriginDutyStation)
	if moveOrder.Grade != nil {
		moveOrder.Entitlement.SetWeightAllotment(*moveOrder.Grade)
	}
	entitlements := Entitlement(moveOrder.Entitlement)
	payload := ghcmessages.MoveOrder{
		DestinationDutyStation: destinationDutyStation,
		Entitlement:            entitlements,
		OrderNumber:            moveOrder.OrderNumber,
		OrderTypeDetail:        moveOrder.OrderTypeDetail,
		ID:                     strfmt.UUID(moveOrder.ID.String()),
		OriginDutyStation:      originDutyStation,
	}

	if moveOrder.Customer != nil {
		payload.Agency = moveOrder.Customer.Agency
		payload.CustomerID = strfmt.UUID(moveOrder.CustomerID.String())
		payload.FirstName = moveOrder.Customer.FirstName
		payload.LastName = moveOrder.Customer.LastName
	}
	if moveOrder.ReportByDate != nil {
		payload.ReportByDate = strfmt.Date(*moveOrder.ReportByDate)
	}
	if moveOrder.DateIssued != nil {
		payload.DateIssued = strfmt.Date(*moveOrder.DateIssued)
	}
	if moveOrder.Grade != nil {
		payload.Grade = *moveOrder.Grade
	}
	if moveOrder.ConfirmationNumber != nil {
		payload.ConfirmationNumber = *moveOrder.ConfirmationNumber
	}
	if moveOrder.OrderType != nil {
		payload.OrderType = *moveOrder.OrderType
	}

	return &payload
}

// Entitlement payload
func Entitlement(entitlement *models.Entitlement) *ghcmessages.Entitlements {
	if entitlement == nil {
		return nil
	}
	var proGearWeight, proGearWeightSpouse, totalWeight int64
	if entitlement.WeightAllotment() != nil {
		proGearWeight = int64(entitlement.WeightAllotment().ProGearWeight)
		proGearWeightSpouse = int64(entitlement.WeightAllotment().ProGearWeightSpouse)
		totalWeight = int64(entitlement.WeightAllotment().TotalWeightSelf)
	}
	var authorizedWeight *int64
	if entitlement.AuthorizedWeight() != nil {
		aw := int64(*entitlement.AuthorizedWeight())
		authorizedWeight = &aw
	}
	var sit int64
	if entitlement.StorageInTransit != nil {
		sit = int64(*entitlement.StorageInTransit)
	}
	var totalDependents int64
	if entitlement.TotalDependents != nil {
		totalDependents = int64(*entitlement.TotalDependents)
	}
	return &ghcmessages.Entitlements{
		ID:                    strfmt.UUID(entitlement.ID.String()),
		AuthorizedWeight:      authorizedWeight,
		DependentsAuthorized:  entitlement.DependentsAuthorized,
		NonTemporaryStorage:   entitlement.NonTemporaryStorage,
		PrivatelyOwnedVehicle: entitlement.PrivatelyOwnedVehicle,
		ProGearWeight:         proGearWeight,
		ProGearWeightSpouse:   proGearWeightSpouse,
		StorageInTransit:      sit,
		TotalDependents:       totalDependents,
		TotalWeight:           totalWeight,
	}
}

// DutyStation payload
func DutyStation(dutyStation *models.DutyStation) *ghcmessages.DutyStation {
	if dutyStation == nil {
		return nil
	}
	address := Address(&dutyStation.Address)
	payload := ghcmessages.DutyStation{
		Address:   address,
		AddressID: address.ID,
		ID:        strfmt.UUID(dutyStation.ID.String()),
		Name:      dutyStation.Name,
	}
	return &payload
}

// Address payload
func Address(address *models.Address) *ghcmessages.Address {
	if address == nil {
		return nil
	}
	return &ghcmessages.Address{
		ID:             strfmt.UUID(address.ID.String()),
		StreetAddress1: &address.StreetAddress1,
		StreetAddress2: address.StreetAddress2,
		StreetAddress3: address.StreetAddress3,
		City:           &address.City,
		State:          &address.State,
		PostalCode:     &address.PostalCode,
		Country:        address.Country,
	}
}

// MTOShipment payload
func MTOShipment(mtoShipment *models.MTOShipment) *ghcmessages.MTOShipment {
	strfmt.MarshalFormat = strfmt.RFC3339Micro

	payload := &ghcmessages.MTOShipment{
		ID:                       strfmt.UUID(mtoShipment.ID.String()),
		MoveTaskOrderID:          strfmt.UUID(mtoShipment.MoveTaskOrderID.String()),
		ShipmentType:             mtoShipment.ShipmentType,
		Status:                   string(mtoShipment.Status),
		CustomerRemarks:          mtoShipment.CustomerRemarks,
		RejectionReason:          mtoShipment.RejectionReason,
		PickupAddress:            Address(&mtoShipment.PickupAddress),
		SecondaryDeliveryAddress: Address(mtoShipment.SecondaryDeliveryAddress),
		SecondaryPickupAddress:   Address(mtoShipment.SecondaryPickupAddress),
		DestinationAddress:       Address(&mtoShipment.DestinationAddress),
		CreatedAt:                strfmt.DateTime(mtoShipment.CreatedAt),
		UpdatedAt:                strfmt.DateTime(mtoShipment.UpdatedAt),
	}

	if mtoShipment.RequestedPickupDate != nil {
		payload.RequestedPickupDate = *handlers.FmtDatePtr(mtoShipment.RequestedPickupDate)
	}

	if mtoShipment.ApprovedDate != nil {
		payload.ApprovedDate = strfmt.Date(*mtoShipment.ApprovedDate)
	}

	return payload
}

// MTOShipmentWithEtag payload
func MTOShipmentWithEtag(mtoShipment *models.MTOShipment) *ghcmessages.MTOShipmentWithEtag {
	return &ghcmessages.MTOShipmentWithEtag{
		MTOShipment: *MTOShipment(mtoShipment),
		ETag:        base64.StdEncoding.EncodeToString([]byte(mtoShipment.UpdatedAt.Format(time.RFC3339Nano))),
	}
}

// MTOShipments payload
func MTOShipments(mtoShipments *models.MTOShipments) *ghcmessages.MTOShipments {
	payload := make(ghcmessages.MTOShipments, len(*mtoShipments))

	for i, m := range *mtoShipments {
		payload[i] = MTOShipmentWithEtag(&m)
	}
	return &payload
}
