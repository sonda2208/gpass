package walletobject

type FlightObject struct {
	Kind                          string                  `json:"kind,omitempty"`
	ID                            string                  `json:"id,omitempty"`
	ClassID                       string                  `json:"classId,omitempty"`
	Version                       string                  `json:"version,omitempty"`
	State                         string                  `json:"state,omitempty"`
	Barcode                       *Barcode                `json:"barcode,omitempty"`
	Messages                      []Message               `json:"messages,omitempty"`
	ValidTimeInterval             *TimeInterval           `json:"validTimeInterval,omitempty"`
	Locations                     []LatLongPoint          `json:"locations,omitempty"`
	HasUsers                      bool                    `json:"hasUsers,omitempty"`
	SmartTapRedemptionValue       string                  `json:"smartTapRedemptionValue,omitempty"`
	HasLinkedDevice               bool                    `json:"hasLinkedDevice,omitempty"`
	DisableExpirationNotification bool                    `json:"disableExpirationNotification,omitempty"`
	InfoModuleData                InfoModule              `json:"infoModuleData,omitempty"`
	ImageModulesData              []ImageModule           `json:"imageModulesData,omitempty"`
	TextModulesData               []TextModule            `json:"textModulesData,omitempty"`
	LinksModuleData               *LinksModule            `json:"linksModuleData,omitempty"`
	ClassReference                *FlightClass            `json:"classReference,omitempty"`
	PassengerName                 string                  `json:"passengerName,omitempty"`
	BoardingAndSeatingInfo        *BoardingAndSeatingInfo `json:"boardingAndSeatingInfo,omitempty"`
	ReservationInfo               *ReservationInfo        `json:"reservationInfo,omitempty"`
	SecurityProgramLogo           *Image                  `json:"securityProgramLogo,omitempty"`
}
