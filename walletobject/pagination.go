package walletobject

type Pagination struct {
	Kind           string `json:"kind,omitempty"`
	NextPageToken  string `json:"nextPageToken,omitempty"`
	ResultsPerPage int    `json:"resultsPerPage,omitempty"`
}
