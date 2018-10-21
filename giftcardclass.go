package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	GiftCardClassResourcePath = "giftCardClass"
)

type GiftCardClassClient struct {
	Client
}

func NewGiftCardClassClient(basePath string, client HttpClient) *GiftCardClassClient {
	return &GiftCardClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GiftCardClassResourcePath,
		},
	}
}

func (c *GiftCardClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GiftCardClass, error) {
	o := &walletobject.GiftCardClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftCardClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardClassClient) Get(id string) (*walletobject.GiftCardClass, error) {
	o := &walletobject.GiftCardClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftCardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftCardClassResourcePath,
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

func (c *GiftCardClassClient) Insert(o *walletobject.GiftCardClass) (*walletobject.GiftCardClass, error) {
	no := &walletobject.GiftCardClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftCardClassResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GiftCardClassClient) Patch(id string, i interface{}) (*walletobject.GiftCardClass, error) {
	o := &walletobject.GiftCardClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GiftCardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardClassClient) Update(id string, o *walletobject.GiftCardClass) (*walletobject.GiftCardClass, error) {
	no := &walletobject.GiftCardClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GiftCardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
