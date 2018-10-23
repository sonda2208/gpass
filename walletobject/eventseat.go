package walletobject

type EventSeat struct {
	Kind    string           `json:"kind,omitempty"`
	Seat    *LocalizedString `json:"seat,omitempty"`
	Row     *LocalizedString `json:"row,omitempty"`
	Section *LocalizedString `json:"section,omitempty"`
	Gate    *LocalizedString `json:"gate,omitempty"`
}
