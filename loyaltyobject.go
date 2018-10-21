package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	LoyaltyObjectResourcePath = "loyaltyObject"
)

type LoyaltyObjectClient struct {
	Client
}

func NewLoyaltyObjectClient(basePath string, client HttpClient) *LoyaltyObjectClient {
	return &LoyaltyObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: LoyaltyObjectResourcePath,
		},
	}
}

func (c *LoyaltyObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.LoyaltyObject, error) {
	l := &walletobject.LoyaltyObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + LoyaltyObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyObjectClient) Get(id string) (*walletobject.LoyaltyObject, error) {
	l := &walletobject.LoyaltyObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + LoyaltyObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + LoyaltyObjectResourcePath,
		queryParams: &QueryParams{},
		payload:     nil,
		service:     &c.Client,
	}

	req.queryParams.Set("classId", classID)

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

func (c *LoyaltyObjectClient) Insert(l *walletobject.LoyaltyObject) (*walletobject.LoyaltyObject, error) {
	nl := &walletobject.LoyaltyObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + LoyaltyObjectResourcePath,
		queryParams: nil,
		payload:     l,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nl); err != nil {
		return nil, err
	}

	return nl, nil
}

func (c *LoyaltyObjectClient) Patch(id string, i interface{}) (*walletobject.LoyaltyObject, error) {
	l := &walletobject.LoyaltyObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + LoyaltyObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(l); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *LoyaltyObjectClient) Update(id string, l *walletobject.LoyaltyObject) (*walletobject.LoyaltyObject, error) {
	nl := &walletobject.LoyaltyObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + LoyaltyObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     l,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nl); err != nil {
		return nil, err
	}

	return nl, nil
}
