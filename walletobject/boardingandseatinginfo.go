package walletobject

type BoardingAndSeatingInfo struct {
	Kind                   string `json:"kind,omitempty"`
	BoardingGroup          string `json:"boardingGroup,omitempty"`
	SeatNumber             string `json:"seatNumber,omitempty"`
	BoardingPosition       string `json:"boardingPosition,omitempty"`
	SequenceNumber         string `json:"sequenceNumber,omitempty"`
	SeatClass              string `json:"seatClass,omitempty"`
	BoardingDoor           string `json:"boardingDoor,omitempty"`
	BoardingPrivilegeImage *Image `json:"boardingPrivilegeImage,omitempty"`
}
