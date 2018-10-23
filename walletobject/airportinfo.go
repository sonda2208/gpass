package walletobject

type AirportInfo struct {
	Kind                string           `json:"kind,omitempty"`
	AirportIataCode     string           `json:"airportIataCode,omitempty"`
	AirportNameOverride *LocalizedString `json:"airportNameOverride,omitempty"`
	Terminal            string           `json:"terminal,omitempty"`
	Gate                string           `json:"gate,omitempty"`
}
