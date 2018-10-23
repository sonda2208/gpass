package walletobject

type BoardingAndSeatingPolicy struct {
	Kind            string `json:"kind,omitempty"`
	BoardingPolicy  string `json:"boardingPolicy,omitempty"`
	SeatClassPolicy string `json:"seatClassPolicy,omitempty"`
}
