package walletobject

type EventVenue struct {
	Kind    string           `json:"kind,omitempty"`
	Name    *LocalizedString `json:"name,omitempty"`
	Address *LocalizedString `json:"address,omitempty"`
}
