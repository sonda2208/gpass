package walletobject

type TranslatedString struct {
	Kind     string `json:"kind,omitempty"`
	Language string `json:"language,omitempty"`
	Value    string `json:"value,omitempty"`
}
