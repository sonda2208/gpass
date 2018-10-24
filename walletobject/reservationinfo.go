package walletobject

type ReservationInfo struct {
	Kind              string             `json:"kind,omitempty"`
	ConfirmationCode  string             `json:"confirmationCode,omitempty"`
	EticketNumber     string             `json:"eticketNumber,omitempty"`
	FrequentFlyerInfo *FrequentFlyerInfo `json:"frequentFlyerInfo,omitempty"`
}
