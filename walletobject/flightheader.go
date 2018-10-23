package walletobject

type FlightHeader struct {
	Kind                  string         `json:"kind,omitempty"`
	Carrier               *FlightCarrier `json:"carrier,omitempty"`
	FlightNumber          string         `json:"flightNumber,omitempty"`
	OperatingCarrier      *FlightCarrier `json:"operatingCarrier,omitempty"`
	OperatingFlightNumber string         `json:"operatingFlightNumber,omitempty"`
}
