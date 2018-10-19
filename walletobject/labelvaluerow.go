package walletobject

type LabelValueRow struct {
	Columns            []LabelValue `json:"columns,omitempty"`
	HexBackgroundColor string       `json:"hexBackgroundColor,omitempty"`
	HextFontColor      string       `json:"hextFontColor,omitempty"`
}
