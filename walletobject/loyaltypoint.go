package walletobject

type LoyaltyPoint struct {
	Label               string               `json:"label,omitempty"`
	PointsType          string               `json:"pointsType,omitempty"`
	Balance             *LoyaltyPointBalance `json:"balance,omitempty"`
	LocalizedLabel      *LocalizedString     `json:"localizedLabel,omitempty"`
	LocalizedPointsType *LocalizedString     `json:"localizedPointsType,omitempty"`
	PointsValidInterval *TimeInterval        `json:"pointsValidInterval,omitempty"`
}
