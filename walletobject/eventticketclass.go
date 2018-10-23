package walletobject

type EventTicketClass struct {
	Kind                                   string             `json:"kind,omitempty"`
	ID                                     string             `json:"id,omitempty"`
	Version                                string             `json:"version,omitempty"`
	IssuerName                             string             `json:"issuerName,omitempty"`
	LocalizedIssuerName                    *LocalizedString   `json:"localizedIssuerName,omitempty"`
	Messages                               []Message          `json:"messages,omitempty"`
	AllowMultipleUsersPerObject            bool               `json:"allowMultipleUsersPerObject,omitempty"`
	HomepageURI                            *URI               `json:"homepageUri,omitempty"`
	Locations                              []LatLongPoint     `json:"locations,omitempty"`
	ReviewStatus                           string             `json:"reviewStatus,omitempty"`
	Review                                 *CommonClassReview `json:"review,omitempty"`
	InfoModuleData                         *InfoModule        `json:"infoModuleData,omitempty"`
	ImageModulesData                       []ImageModule      `json:"imageModulesData,omitempty"`
	TextModulesData                        []TextModule       `json:"textModulesData,omitempty"`
	LinksModuleData                        []LinksModule      `json:"linksModuleData,omitempty"`
	RedemptionIssuers                      []int              `json:"redemptionIssuers,omitempty"`
	CountryCode                            string             `json:"countryCode,omitempty"`
	HeroImage                              *Image             `json:"heroImage,omitempty"`
	EnableSmartTap                         bool               `json:"enableSmartTap,omitempty"`
	HexBackgroundColor                     string             `json:"hexBackgroundColor,omitempty"`
	MultipleDevicesAndHoldersAllowedStatus string             `json:"multipleDevicesAndHoldersAllowedStatus,omitempty"`
	EventName                              *LocalizedString   `json:"eventName,omitempty"`
	EventID                                string             `json:"eventId,omitempty"`
	Logo                                   *Image             `json:"logo,omitempty"`
	Venue                                  *EventVenue        `json:"venue,omitempty"`
	DateTime                               *EventDateTime     `json:"dateTime,omitempty"`
	FinePrint                              *LocalizedString   `json:"finePrint,omitempty"`
	ConfirmationCodeLabel                  string             `json:"confirmationCodeLabel,omitempty"`
	CustomConfirmationCodeLabel            *LocalizedString   `json:"customConfirmationCodeLabel,omitempty"`
	SeatLabel                              string             `json:"seatLabel,omitempty"`
	CustomSeatLabel                        *LocalizedString   `json:"customSeatLabel,omitempty"`
	RowLabel                               string             `json:"rowLabel,omitempty"`
	CustomRowLabel                         *LocalizedString   `json:"customRowLabel,omitempty"`
	SectionLabel                           string             `json:"sectionLabel,omitempty"`
	CustomSectionLabel                     *LocalizedString   `json:"customSectionLabel,omitempty"`
	GateLabel                              string             `json:"gateLabel,omitempty"`
	CustomGateLabel                        *LocalizedString   `json:"customGateLabel,omitempty"`
}
