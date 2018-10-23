package walletobject

type EventDateTime struct {
	Kind                 string           `json:"kind,omitempty"`
	DoorsOpenLabel       string           `json:"doorsOpenLabel,omitempty"`
	CustomDoorsOpenLabel *LocalizedString `json:"customDoorsOpenLabel,omitempty"`
	DoorsOpen            string           `json:"doorsOpen,omitempty"`
	Start                string           `json:"start,omitempty"`
	End                  string           `json:"end,omitempty"`
}
