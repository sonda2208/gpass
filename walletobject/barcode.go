package walletobject

type Barcode struct {
	AlternateText string `json:"alternateText,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Label         string `json:"label,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         string `json:"value,omitempty"`
}
