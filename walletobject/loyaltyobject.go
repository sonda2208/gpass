package walletobject

type LoyaltyObject struct {
	AccountID               string         `json:"accountId,omitempty"`
	AccountName             string         `json:"accountName,omitempty"`
	ClassID                 string         `json:"classId,omitempty"`
	ID                      string         `json:"id,omitempty"`
	SmartTapRedemptionValue string         `json:"smartTapRedemptionValue,omitempty"`
	State                   string         `json:"state,omitempty"`
	Kind                    string         `json:"kind,omitempty"`
	LinkedOfferIds          []string       `json:"linkedOfferIds,omitempty"`
	HasLinkedDevice         bool           `json:"hasLinkedDevice,omitempty"`
	HasUsers                bool           `json:"hasUsers,omitempty"`
	Barcode                 *Barcode       `json:"barcode,omitempty"`
	IssuerData              *TypedValue    `json:"issuerData,omitempty"`
	ClassReference          *LoyaltyClass  `json:"classReference,omitempty"`
	InfoModuleData          *InfoModule    `json:"infoModuleData,omitempty"`
	LinksModuleData         *LinksModule   `json:"linksModuleData,omitempty"`
	TextModulesData         []TextModule   `json:"textModulesData,omitempty"`
	ImageModulesData        []ImageModule  `json:"imageModulesData,omitempty"`
	Locations               []LatLongPoint `json:"locations,omitempty"`
	Message                 []Message      `json:"message,omitempty"`
	LoyaltyPoints           *LoyaltyPoint  `json:"loyaltyPoints,omitempty"`
	SecondaryLoyaltyPoints  *LoyaltyPoint  `json:"secondaryLoyaltyPoints,omitempty"`
	ValidTimeInterval       *TimeInterval  `json:"validTimeInterval,omitempty"`
	Version                 int64          `json:"version,omitempty"`
}
