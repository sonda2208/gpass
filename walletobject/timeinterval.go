package walletobject

type TimeInterval struct {
	Kind  string    `json:"kind,omitempty"`
	Start *DateTime `json:"start,omitempty"`
	End   *DateTime `json:"end,omitempty"`
}
