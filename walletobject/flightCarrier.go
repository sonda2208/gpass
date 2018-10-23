package walletobject

type FlightCarrier struct {
	Kind                string           `json:"kind,omitempty"`
	CarrierIataCode     string           `json:"carrierIataCode,omitempty"`
	AirlineName         *LocalizedString `json:"airlineName,omitempty"`
	AirlineLogo         *Image           `json:"airlineLogo,omitempty"`
	AirlineAllianceLogo *Image           `json:"airlineAllianceLogo,omitempty"`
}
