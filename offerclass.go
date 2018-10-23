package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	OfferClassResourcePath = "offerClass"
)

type OfferClassClient struct {
	Client
}

func NewOfferClassClient(basePath string, client HTTPClient) *OfferClassClient {
	return &OfferClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: OfferClassResourcePath,
		},
	}
}

func (c *OfferClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.OfferClass, error) {
	o := &walletobject.OfferClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + OfferClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferClassClient) Get(id string) (*walletobject.OfferClass, error) {
	o := &walletobject.OfferClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + OfferClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + OfferClassResourcePath,
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

func (c *OfferClassClient) Insert(o *walletobject.OfferClass) (*walletobject.OfferClass, error) {
	no := &walletobject.OfferClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + OfferClassResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *OfferClassClient) Patch(id string, i interface{}) (*walletobject.OfferClass, error) {
	o := &walletobject.OfferClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + OfferClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferClassClient) Update(id string, o *walletobject.OfferClass) (*walletobject.OfferClass, error) {
	no := &walletobject.OfferClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + OfferClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
