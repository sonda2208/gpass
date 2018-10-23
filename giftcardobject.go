package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	GiftcardObjectResourcePath = "giftCardObject"
)

type GiftcardObjectClient struct {
	Client
}

func NewGiftcardObjectClient(basePath string, client HTTPClient) *GiftcardObjectClient {
	return &GiftcardObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GiftcardObjectResourcePath,
		},
	}
}

func (c *GiftcardObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GiftcardObject, error) {
	o := &walletobject.GiftcardObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftcardObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardObjectClient) Get(id string) (*walletobject.GiftcardObject, error) {
	o := &walletobject.GiftcardObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftcardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftcardObjectResourcePath,
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

func (c *GiftcardObjectClient) Insert(o *walletobject.GiftcardObject) (*walletobject.GiftcardObject, error) {
	no := &walletobject.GiftcardObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftcardObjectResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GiftcardObjectClient) Patch(id string, i interface{}) (*walletobject.GiftcardObject, error) {
	o := &walletobject.GiftcardObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GiftcardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftcardObjectClient) Update(id string, o *walletobject.GiftcardObject) (*walletobject.GiftcardObject, error) {
	no := &walletobject.GiftcardObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GiftcardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
