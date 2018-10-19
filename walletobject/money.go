package walletobject

type Money struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Micros       int64  `json:"micros,omitempty"`
}
