package walletobject

type EventReservationInfo struct {
	Kind             string `json:"kind,omitempty"`
	ConfirmationCode string `json:"confirmationCode,omitempty"`
}
