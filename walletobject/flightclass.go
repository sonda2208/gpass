package walletobject

type FlightClass struct {
	Kind                                    string                    `json:"kind,omitempty"`
	ID                                      string                    `json:"id,omitempty"`
	Version                                 string                    `json:"version,omitempty"`
	IssuerName                              string                    `json:"issuerName,omitempty"`
	LocalizedIssuerName                     *LocalizedString          `json:"localizedIssuerName,omitempty"`
	Messages                                []Message                 `json:"messages,omitempty"`
	AllowMultipleUsersPerObject             bool                      `json:"allowMultipleUsersPerObject,omitempty"`
	HomepageURI                             *URI                      `json:"homepageUri,omitempty"`
	Locations                               []LatLongPoint            `json:"locations,omitempty"`
	ReviewStatus                            string                    `json:"reviewStatus,omitempty"`
	Review                                  *CommonClassReview        `json:"review,omitempty"`
	InfoModuleData                          *InfoModule               `json:"infoModuleData,omitempty"`
	ImageModulesData                        []ImageModule             `json:"imageModulesData,omitempty"`
	TextModulesData                         []TextModule              `json:"textModulesData,omitempty"`
	LinksModuleData                         *LinksModule              `json:"linksModuleData,omitempty"`
	RedemptionIssuers                       []int                     `json:"redemptionIssuers,omitempty"`
	CountryCode                             string                    `json:"countryCode,omitempty"`
	HeroImage                               *Image                    `json:"heroImage,omitempty"`
	EnableSmartTap                          bool                      `json:"enableSmartTap,omitempty"`
	HexBackgroundColor                      string                    `json:"hexBackgroundColor,omitempty"`
	MultipleDevicesAndHoldersAllowedStatus  string                    `json:"multipleDevicesAndHoldersAllowedStatus,omitempty"`
	LocalScheduledDepartureDateTime         string                    `json:"localScheduledDepartureDateTime,omitempty"`
	LocalEstimatedOrActualDepartureDateTime string                    `json:"localEstimatedOrActualDepartureDateTime,omitempty"`
	LocalBoardingDateTime                   string                    `json:"localBoardingDateTime,omitempty"`
	LocalGateClosingDateTime                string                    `json:"localGateClosingDateTime,omitempty"`
	LocalScheduledArrivalDateTime           string                    `json:"localScheduledArrivalDateTime,omitempty"`
	LocalEstimatedOrActualArrivalDateTime   string                    `json:"localEstimatedOrActualArrivalDateTime,omitempty"`
	FlightHeader                            *FlightHeader             `json:"flightHeader,omitempty"`
	Origin                                  *AirportInfo              `json:"origin,omitempty"`
	Destination                             *AirportInfo              `json:"destination,omitempty"`
	FlightStatus                            string                    `json:"flightStatus,omitempty"`
	BoardingAndSeatingPolicy                *BoardingAndSeatingPolicy `json:"boardingAndSeatingPolicy,omitempty"`
}
