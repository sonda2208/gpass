package walletobject

type Message struct {
	Body            string           `json:"body,omitempty"`
	Header          string           `json:"header,omitempty"`
	ID              string           `json:"id,omitempty"`
	Kind            string           `json:"kind,omitempty"`
	MessageType     string           `json:"messageType,omitempty"`
	ActionURI       *URI             `json:"actionUri,omitempty"`
	Image           *Image           `json:"image,omitempty"`
	DisplayInterval *TimeInterval    `json:"displayInterval,omitempty"`
	LocalizedBody   *LocalizedString `json:"localizedBody,omitempty"`
	LocalizedHeader *LocalizedString `json:"localizedHeader,omitempty"`
}
