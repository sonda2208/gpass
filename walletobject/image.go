package walletobject

type Image struct {
	Kind      string `json:"kind,omitempty"`
	SourceURI *URI   `json:"sourceUri,omitempty"`
}
