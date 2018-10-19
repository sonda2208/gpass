package walletobject

type LocalizedString struct {
	Kind             string             `json:"kind,omitempty"`
	DefaultValue     *TranslatedString  `json:"defaultValue,omitempty"`
	TranslatedValues []TranslatedString `json:"translatedValues,omitempty"`
}
