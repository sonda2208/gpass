package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	LoyaltyClassResourcePath = "loyaltyClass"
)

type LoyaltyClassClient struct {
	Client
}

func NewLoyaltyClassClient(basePath string, client HTTPClient) *LoyaltyClassClient {
	return &LoyaltyClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: LoyaltyClassResourcePath,
		},
	}
}

func (c *LoyaltyClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.LoyaltyClass, error) {
	l := &walletobject.LoyaltyClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + LoyaltyClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyClassClient) Get(id string) (*walletobject.LoyaltyClass, error) {
	l := &walletobject.LoyaltyClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + LoyaltyClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + LoyaltyClassResourcePath,
		queryParams: &QueryParams{},
		payload:     nil,
		service:     &c.Client,
	}

	req.queryParams.Set("issuerId", issuerID)

	if maxResults > 0 {
		req.queryParams.Set("maxResults", strconv.Itoa(maxResults))
	}

	if paginationToken != "" {
		req.queryParams.Set("token", paginationToken)
	}

	if err := req.Do().DecodeResponse(r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *LoyaltyClassClient) Insert(l *walletobject.LoyaltyClass) (*walletobject.LoyaltyClass, error) {
	nl := &walletobject.LoyaltyClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + LoyaltyClassResourcePath,
		queryParams: nil,
		payload:     l,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nl); err != nil {
		return nil, err
	}

	return nl, nil
}

func (c *LoyaltyClassClient) Patch(id string, i interface{}) (*walletobject.LoyaltyClass, error) {
	l := &walletobject.LoyaltyClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + LoyaltyClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyClassClient) Update(id string, l *walletobject.LoyaltyClass) (*walletobject.LoyaltyClass, error) {
	nl := &walletobject.LoyaltyClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + LoyaltyClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     l,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nl); err != nil {
		return nil, err
	}

	return nl, nil
}
