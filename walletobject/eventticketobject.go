package walletobject

type EventTicketObject struct {
	Kind                          string                `json:"kind,omitempty"`
	ID                            string                `json:"id,omitempty"`
	ClassID                       string                `json:"classId,omitempty"`
	Version                       string                `json:"version,omitempty"`
	State                         string                `json:"state,omitempty"`
	Barcode                       *Barcode              `json:"barcode,omitempty"`
	Messages                      []Message             `json:"messages,omitempty"`
	ValidTimeInterval             *TimeInterval         `json:"validTimeInterval,omitempty"`
	Locations                     []LatLongPoint        `json:"locations,omitempty"`
	HasUsers                      bool                  `json:"hasUsers,omitempty"`
	SmartTapRedemptionValue       string                `json:"smartTapRedemptionValue,omitempty"`
	HasLinkedDevice               bool                  `json:"hasLinkedDevice,omitempty"`
	DisableExpirationNotification bool                  `json:"disableExpirationNotification,omitempty"`
	InfoModuleData                *InfoModule           `json:"infoModuleData,omitempty"`
	ImageModulesData              []ImageModule         `json:"imageModulesData,omitempty"`
	TextModulesData               []TextModule          `json:"textModulesData,omitempty"`
	LinksModuleData               *LinksModule          `json:"linksModuleData,omitempty"`
	ClassReference                *EventTicketClass     `json:"classReference,omitempty"`
	SeatInfo                      *EventSeat            `json:"seatInfo,omitempty"`
	ReservationInfo               *EventReservationInfo `json:"reservationInfo,omitempty"`
	TicketHolderName              string                `json:"ticketHolderName,omitempty"`
	TicketNumber                  string                `json:"ticketNumber,omitempty"`
	TicketType                    *LocalizedString      `json:"ticketType,omitempty"`
	FaceValue                     *Money                `json:"faceValue,omitempty"`
}
