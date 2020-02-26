package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// ReServiceName is a type of service name
type ReServiceName string

const (
	// DomesticLinehaul is a Service Name for  Domestic Linehaul
	DomesticLinehaul ReServiceName = "Dom. Linehaul"
	// FuelSurcharge is a Service Name for  Fuel Surchage
	FuelSurcharge ReServiceName = "Fuel Surcharge"
	// DomesticOriginPrice is a Service Name for Domestic Origin Price
	DomesticOriginPrice ReServiceName = "Dom. Origin Price"
	// DomesticDestinationPrice  is a Service Name for Domestic Destination Price
	DomesticDestinationPrice ReServiceName = "Dom. Destination Price"
	// DomesticPacking a Service Name for Domestic Packing
	DomesticPacking ReServiceName = "Dom. Packing"
	// DomesticUnpacking is a Service Name for Domestic Unpacking
	DomesticUnpacking ReServiceName = "Dom. Unpacking"
	// DomesticShorthaul is a Service Name for Domestic Shorthaul
	DomesticShorthaul ReServiceName = "Dom. Shorthaul"
	// DomesticNTSPackingFactor is a Service Name for Domestic NTS Packing Factor
	DomesticNTSPackingFactor ReServiceName = "Dom. NTS Packing Factor"
	// DomesticMobileHomeFactor is a Service Name for Domestic Mobile Home Factor
	DomesticMobileHomeFactor ReServiceName = "Dom. Mobile Home Factor"
	// DomesticHaulAwayBoatFactor is a Service Name for Domestic Haul Away Boat Factor
	DomesticHaulAwayBoatFactor ReServiceName = "Dom. Haul Away Boat Factor"
	// DomesticTowAwayBoatFactor is a Service Name for Domestic Tow Away Boat Factor
	DomesticTowAwayBoatFactor ReServiceName = "Dom. Tow Away Boat Factor"
)

// ReService model struct
type ReService struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ReServices is not required by pop and may be deleted
type ReServices []ReService

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *ReService) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.Code, Name: "Code"},
		&validators.StringIsPresent{Field: r.Name, Name: "Name"},
	), nil
}
