package walletobject

type LoyaltyPointBalance struct {
	Double float64 `json:"double,omitempty"`
	Int    int     `json:"int,omitempty"`
	String string  `json:"string,omitempty"`
	Money  *Money  `json:"money,omitempty"`
}
