package walletobject

type URI struct {
	Description          string           `json:"description,omitempty"`
	Kind                 string           `json:"kind,omitempty"`
	URI                  string           `json:"uri,omitempty"`
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
}
