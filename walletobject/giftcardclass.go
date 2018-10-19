package walletobject

type GiftCardClass struct {
	AllowBarcodeRedemption                 bool               `json:"allowBarcodeRedemption,omitempty"`
	AllowMultipleUsersPerObject            bool               `json:"allowMultipleUsersPerObject,omitempty"`
	EnableSmartTap                         bool               `json:"enableSmartTap,omitempty"`
	HideBarcode                            bool               `json:"hideBarcode,omitempty"`
	ID                                     string             `json:"id,omitempty"`
	Kind                                   string             `json:"kind,omitempty"`
	IssuerName                             string             `json:"issuerName,omitempty"`
	MultipleDevicesAndHoldersAllowedStatus string             `json:"multipleDevicesAndHoldersAllowedStatus,omitempty"`
	MerchantName                           string             `json:"merchantName,omitempty"`
	PinLabel                               string             `json:"pinLabel,omitempty"`
	CountryCode                            string             `json:"countryCode,omitempty"`
	RedemptionMessage                      string             `json:"redemptionMessage,omitempty"`
	HexBackgroundColor                     string             `json:"hexBackgroundColor,omitempty"`
	EventNumberLabel                       string             `json:"eventNumberLabel,omitempty"`
	ReviewStatus                           string             `json:"reviewStatus,omitempty"`
	HomepageURI                            *URI               `json:"homepageURI,omitempty"`
	IssuerData                             *TypedValue        `json:"issuerData,omitempty"`
	LocalizedIssuerName                    *LocalizedString   `json:"localizedIssuerName,omitempty"`
	LocalizedMerchantName                  *LocalizedString   `json:"localizedMerchantName,omitempty"`
	LocalizedEventNumberLabel              *LocalizedString   `json:"localizedEventNumberLabel,omitempty"`
	LocalizedPinLabel                      *LocalizedString   `json:"localizedPinLabel,omitempty"`
	Locations                              []LatLongPoint     `json:"locations,omitempty"`
	Messages                               []Message          `json:"messages,omitempty"`
	InfoModuleData                         *InfoModule        `json:"infoModuleData,omitempty"`
	LinksModuleData                        *LinksModule       `json:"linksModuleData,omitempty"`
	TextModulesData                        []TextModule       `json:"textModulesData,omitempty"`
	ImageModulesData                       []ImageModule      `json:"imageModulesData,omitempty"`
	Review                                 *CommonClassReview `json:"review,omitempty"`
	HeroImage                              *Image             `json:"heroImage,omitempty"`
	ProgramLogo                            *Image             `json:"programLogo,omitempty"`
	Version                                string             `json:"version,omitempty"`
}
