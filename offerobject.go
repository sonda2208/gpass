package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	OfferObjectResourcePath = "offerObject"
)

type OfferObjectClient struct {
	Client
}

func NewOfferObjectClient(basePath string, client HTTPClient) *OfferObjectClient {
	return &OfferObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: OfferObjectResourcePath,
		},
	}
}

func (c *OfferObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.OfferObject, error) {
	o := &walletobject.OfferObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + OfferObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferObjectClient) Get(id string) (*walletobject.OfferObject, error) {
	o := &walletobject.OfferObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + OfferObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + OfferObjectResourcePath,
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

func (c *OfferObjectClient) Insert(o *walletobject.OfferObject) (*walletobject.OfferObject, error) {
	no := &walletobject.OfferObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + OfferObjectResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *OfferObjectClient) Patch(id string, i interface{}) (*walletobject.OfferObject, error) {
	o := &walletobject.OfferObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + OfferObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *OfferObjectClient) Update(id string, o *walletobject.OfferObject) (*walletobject.OfferObject, error) {
	no := &walletobject.OfferObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + OfferObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
