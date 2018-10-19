package walletobject

type LatLongPoint struct {
	Kind      string            `json:"kind,omitempty"`
	Latitude  float64           `json:"latitude,omitempty"`
	Longitude float64           `json:"longitude,omitempty"`
	Meta      *LocationMetaData `json:"meta,omitempty"`
}
