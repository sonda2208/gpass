package walletobject

type InfoModule struct {
	LabelValueRows     []LabelValueRow `json:"labelValueRows,omitempty"`
	HexBackgroundColor string          `json:"hexBackgroundColor,omitempty"`
	HexFontColor       string          `json:"hexFontColor,omitempty"`
	ShowLastUpdateTime bool            `json:"showLastUpdateTime,omitempty"`
}
