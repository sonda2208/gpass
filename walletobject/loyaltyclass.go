package walletobject

type LoyaltyClass struct {
	AccountIDLabel                         string             `json:"accountIdLabel,omitempty"`
	AccountNameLabel                       string             `json:"accountNameLabel,omitempty"`
	CountryCode                            string             `json:"countryCode,omitempty"`
	HexBackgroundColor                     string             `json:"hexBackgroundColor,omitempty"`
	ID                                     string             `json:"id,omitempty"`
	IssuerName                             string             `json:"issuerName,omitempty"`
	Kind                                   string             `json:"kind,omitempty"`
	MultipleDevicesAndHoldersAllowedStatus string             `json:"multipleDevicesAndHoldersAllowedStatus,omitempty"`
	ProgramName                            string             `json:"programName,omitempty"`
	ReviewStatus                           string             `json:"reviewStatus,omitempty"`
	RewardsTier                            string             `json:"rewardsTier,omitempty"`
	RewardsTierLabel                       string             `json:"rewardsTierLabel,omitempty"`
	EnableSmartTap                         bool               `json:"enableSmartTap,omitempty"`
	AllowMultipleUsersPerObject            bool               `json:"allowMultipleUsersPerObject,omitempty"`
	HideBarcode                            bool               `json:"hideBarcode,omitempty"`
	HomepageURI                            *URI               `json:"homepageUri,omitempty"`
	HeroImage                              *Image             `json:"heroImage,omitempty"`
	IssuerData                             *TypedValue        `json:"issuerData,omitempty"`
	LinksModuleData                        *LinksModule       `json:"linksModuleData,omitempty"`
	LocalizedAccountIDLabel                *LocalizedString   `json:"localizedAccountIdLabel,omitempty"`
	LocalizedAccountNameLabel              *LocalizedString   `json:"localizedAccountNameLabel,omitempty"`
	LocalizedIssuerName                    *LocalizedString   `json:"localizedIssuerName,omitempty"`
	LocalizedProgramName                   *LocalizedString   `json:"localizedProgramName,omitempty"`
	LocalizedRewardsTier                   *LocalizedString   `json:"localizedRewardsTier,omitempty"`
	LocalizedRewardsTierLabel              *LocalizedString   `json:"localizedRewardsTierLabel,omitempty"`
	TextModulesData                        []TextModule       `json:"textModulesData,omitempty"`
	ImageModulesData                       []ImageModule      `json:"imageModulesData,omitempty"`
	Locations                              []LatLongPoint     `json:"locations,omitempty"`
	Messages                               []Message          `json:"messages,omitempty"`
	ProgramLogo                            *Image             `json:"programLogo,omitempty"`
	InfoModuleData                         *ImageModule       `json:"infoModuleData,omitempty"`
	Review                                 *CommonClassReview `json:"review,omitempty"`
	Version                                string             `json:"version,omitempty"`
}
