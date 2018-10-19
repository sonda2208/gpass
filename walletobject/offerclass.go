package walletobject

type OfferClass struct {
	CountryCode                            string             `json:"countryCode,omitempty"`
	Details                                string             `json:"details,omitempty"`
	FinePrint                              string             `json:"finePrint,omitempty"`
	HexBackgroundColor                     string             `json:"hexBackgroundColor,omitempty"`
	ID                                     string             `json:"id,omitempty"`
	IssuerName                             string             `json:"issuerName,omitempty"`
	Kind                                   string             `json:"kind,omitempty"`
	AllowMultipleUsersPerObject            bool               `json:"allowMultipleUsersPerObject,omitempty"`
	HideBarcode                            bool               `json:"hideBarcode,omitempty"`
	EnableSmartTap                         bool               `json:"enableSmartTap,omitempty"`
	HeroImage                              *Image             `json:"heroImage,omitempty"`
	InfoModuleData                         *InfoModule        `json:"infoModuleData,omitempty"`
	IssuerData                             *TypedValue        `json:"issuerData,omitempty"`
	HomepageURI                            *URI               `json:"homepageUri,omitempty"`
	HelpURI                                *URI               `json:"helpUri,omitempty"`
	DistributionTimeInterval               *TimeInterval      `json:"distributionTimeInterval,omitempty"`
	LinksModuleData                        *LinksModule       `json:"linksModuleData,omitempty"`
	LocalizedDetails                       *LocalizedString   `json:"localizedDetails,omitempty"`
	LocalizedFinePrint                     *LocalizedString   `json:"localizedFinePrint,omitempty"`
	LocalizedIssuerName                    *LocalizedString   `json:"localizedIssuerName,omitempty"`
	LocalizedProvider                      *LocalizedString   `json:"localizedProvider,omitempty"`
	LocalizedTitle                         *LocalizedString   `json:"localizedTitle,omitempty"`
	ImageModulesData                       []ImageModule      `json:"imageModulesData,omitempty"`
	Locations                              []LatLongPoint     `json:"locations,omitempty"`
	Messages                               []Message          `json:"messages,omitempty"`
	MultipleDevicesAndHoldersAllowedStatus string             `json:"multipleDevicesAndHoldersAllowedStatus,omitempty"`
	Provider                               string             `json:"provider,omitempty"`
	RedemptionChannel                      string             `json:"redemptionChannel,omitempty"`
	Review                                 *CommonClassReview `json:"review,omitempty"`
	ReviewStatus                           string             `json:"reviewStatus,omitempty"`
	TextModulesData                        []TextModule       `json:"textModulesData,omitempty"`
	Title                                  string             `json:"title,omitempty"`
	TitleImage                             *Image             `json:"titleImage,omitempty"`
	Version                                string             `json:"version,omitempty"`
}
