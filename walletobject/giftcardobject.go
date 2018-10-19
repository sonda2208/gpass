package walletobject

type GiftCardObject struct {
	CardNumber              string         `json:"cardNumber,omitempty"`
	ClassID                 string         `json:"classId,omitempty"`
	EventNumber             string         `json:"eventNumber,omitempty"`
	ID                      string         `json:"id,omitempty"`
	Kind                    string         `json:"kind,omitempty"`
	Pin                     string         `json:"pin,omitempty"`
	SmartTapRedemptionValue string         `json:"smartTapRedemptionValue,omitempty"`
	State                   string         `json:"state,omitempty"`
	HasLinkedDevice         bool           `json:"hasLinkedDevice,omitempty"`
	HasUsers                bool           `json:"hasUsers,omitempty"`
	Balance                 *Money         `json:"balance,omitempty"`
	BalanceUpdateTime       *DateTime      `json:"balanceUpdateTime,omitempty"`
	ClassReference          *GiftCardClass `json:"classReference,omitempty"`
	Barcode                 *Barcode       `json:"barcode,omitempty"`
	Locations               []LatLongPoint `json:"locations,omitempty"`
	Messages                []Message      `json:"messages,omitempty"`
	TextModulesData         []TextModule   `json:"textModulesData,omitempty"`
	IssuerData              *TypedValue    `json:"issuerData,omitempty"`
	InfoModuleData          *InfoModule    `json:"infoModuleData,omitempty"`
	LinksModuleData         *LinksModule   `json:"linksModuleData,omitempty"`
	ImageModulesData        *ImageModule   `json:"imageModulesData,omitempty"`
	ValidTimeInterval       *TimeInterval  `json:"validTimeInterval,omitempty"`
}
