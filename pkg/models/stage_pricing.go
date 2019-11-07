package models

type StageDomesticServiceArea struct {
	BasePointCity     string `db:"base_point_city" csv:"base_point_city"`
	State             string `db:"state" csv:"state"`
	ServiceAreaNumber string `db:"service_area_number" csv:"service_area_number"`
	Zip3s             string `db:"zip3s" csv:"zip3s"`
}

type StageInternationalServiceArea struct {
	RateArea   string `db:"rate_area" csv:"rate_area"`
	RateAreaID string `db:"rate_area_id" csv:"rate_area_id"`
}

type StageDomesticLinehaulPrice struct {
	ServiceAreaNumber string `db:"service_area_number" csv:"service_area_number"`
	OriginServiceArea string `db:"origin_service_area" csv:"origin_service_area"`
	ServicesSchedule  string `db:"services_schedule" csv:"services_schedule"`
	Season            string `db:"season" csv:"season"`
	WeightLower       string `db:"weight_lower" csv:"weight_lower"`
	WeightUpper       string `db:"weight_upper" csv:"weight_upper"`
	MilesLower        string `db:"miles_lower" csv:"miles_lower"`
	MilesUpper        string `db:"miles_upper" csv:"miles_upper"`
	EscalationNumber  string `db:"escalation_number" csv:"escalation_number"`
	Rate              string `db:"rate" csv:"rate"`
}

type StageDomesticServiceAreaPrice struct {
	ServiceAreaNumber                     string `db:"service_area_number" csv:"service_area_number"`
	ServiceAreaName                       string `db:"service_area_name" csv:"service_area_name"`
	ServicesSchedule                      string `db:"services_schedule" csv:"services_schedule"`
	SITPickupDeliverySchedule             string `db:"sit_pickup_delivery_schedule" csv:"sit_pickup_delivery_schedule"`
	Season                                string `db:"season" csv:"season"`
	ShorthaulPrice                        string `db:"shorthaul_price" csv:"shorthaul_price"`
	OriginDestinationPrice                string `db:"origin_destination_price" csv:"origin_destination_price"`
	OriginDestinationSITFirstDayWarehouse string `db:"origin_destination_sit_first_day_warehouse" csv:"origin_destination_sit_first_day_warehouse"`
	OriginDestinationSITAddlDays          string `db:"origin_destination_sit_addl_days" csv:"origin_destination_sit_addl_days"`
}

type StageShipmentManagementServicesPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_per_task_order" csv:"price_per_task_order"`
}

type StageCounselingServicesPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_per_task_order" csv:"price_per_task_order"`
}

type StageTransitionPrice struct {
	ContractYear      string `db:"contract_year" csv:"contract_year"`
	PricePerTaskOrder string `db:"price_total_cost" csv:"price_total_cost"`
}

type StageDomesticMoveAccessorialPrices struct {
	ServicesSchedule string `db:"services_schedule" csv:"services_schedule"`
	ServiceProvided  string `db:"service_provided" csv:"service_provided"`
	PricePerUnit     string `db:"price_per_unit" csv:"price_per_unit"`
}

type StageInternationalMoveAccessorialPrices struct {
	Market          string `db:"market" csv:"market"`
	ServiceProvided string `db:"service_provided" csv:"service_provided"`
	PricePerUnit    string `db:"price_per_unit" csv:"price_per_unit"`
}

type StageDomesticInternationalAdditionalPrices struct {
	Market       string `db:"market" csv:"market"`
	ShipmentType string `db:"shipment_type" csv:"shipment_type"`
	Factor       string `db:"factor" csv:"factor"`
}
