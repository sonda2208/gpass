package walletobject

type FrequentFlyerInfo struct {
	Kind                     string           `json:"kind,omitempty"`
	FrequentFlyerNumber      string           `json:"frequentFlyerNumber,omitempty"`
	FrequentFlyerProgramName *LocalizedString `json:"frequentFlyerProgramName,omitempty"`
}
