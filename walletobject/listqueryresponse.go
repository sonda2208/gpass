package walletobject

type ListQueryResponse struct {
	Pagination Pagination    `json:"pagination,omitempty"`
	Resources  []interface{} `json:"resources,omitempty"`
}
