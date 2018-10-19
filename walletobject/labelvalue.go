package walletobject

type LabelValue struct {
	Label          string           `json:"label,omitempty"`
	Value          string           `json:"value,omitempty"`
	LocalizedLabel *LocalizedString `json:"localizedLabel,omitempty"`
	LocalizedValue *LocalizedString `json:"localizedValue,omitempty"`
}
