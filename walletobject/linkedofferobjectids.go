package walletobject

type LinkedOfferObjectIds struct {
	Kind                       string   `json:"kind,omitempty"`
	AddLinkedOfferObjectIds    []string `json:"addLinkedOfferObjectIds,omitempty"`
	RemoveLinkedOfferObjectIds []string `json:"removeLinkedOfferObjectIds,omitempty"`
}
