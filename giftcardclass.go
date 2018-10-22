package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	GiftcardClassResourcePath = "giftCardClass"
)

type GiftcardClassClient struct {
	Client
}

func NewGiftcardClassClient(basePath string, client HttpClient) *GiftcardClassClient {
	return &GiftcardClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GiftcardClassResourcePath,
		},
	}
}

func (c *GiftcardClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GiftcardClass, error) {
	o := &walletobject.GiftcardClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftcardClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardClassClient) Get(id string) (*walletobject.GiftcardClass, error) {
	o := &walletobject.GiftcardClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftcardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftcardClassResourcePath,
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

func (c *GiftcardClassClient) Insert(o *walletobject.GiftcardClass) (*walletobject.GiftcardClass, error) {
	no := &walletobject.GiftcardClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftcardClassResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GiftcardClassClient) Patch(id string, i interface{}) (*walletobject.GiftcardClass, error) {
	o := &walletobject.GiftcardClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GiftcardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardClassClient) Update(id string, o *walletobject.GiftcardClass) (*walletobject.GiftcardClass, error) {
	no := &walletobject.GiftcardClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GiftcardClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
