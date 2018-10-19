package walletobject

type OfferObject struct {
	Barcode                 *Barcode       `json:"barcode,omitempty"`
	ClassID                 string         `json:"classId,omitempty"`
	ClassReference          *OfferClass    `json:"classReference,omitempty"`
	HasLinkedDevice         bool           `json:"hasLinkedDevice,omitempty"`
	HasUsers                bool           `json:"hasUsers,omitempty"`
	ID                      string         `json:"id,omitempty"`
	ImageModulesData        []ImageModule  `json:"imageModulesData,omitempty"`
	InfoModuleData          *InfoModule    `json:"infoModuleData,omitempty"`
	IssuerData              *TypedValue    `json:"issuerData,omitempty"`
	Kind                    string         `json:"kind,omitempty"`
	LinksModuleData         *LinksModule   `json:"linksModuleData,omitempty"`
	Locations               []LatLongPoint `json:"locations,omitempty"`
	Messages                []Message      `json:"messages,omitempty"`
	SmartTapRedemptionValue string         `json:"smartTapRedemptionValue,omitempty"`
	State                   string         `json:"state,omitempty"`
	TextModulesData         []TextModule   `json:"textModulesData,omitempty"`
	ValidTimeInterval       *TimeInterval  `json:"validTimeInterval,omitempty"`
	Version                 int64          `json:"version,omitempty"`
}
