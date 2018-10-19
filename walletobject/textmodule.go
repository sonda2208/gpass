package walletobject

type TextModule struct {
	Header          string           `json:"header,omitempty"`
	Body            string           `json:"body,omitempty"`
	IconBodies      *IconBody        `json:"iconBodies,omitempty"`
	LocalizedBody   *LocalizedString `json:"localizedBody,omitempty"`
	LocalizedHeader *LocalizedString `json:"localizedHeader,omitempty"`
}
