package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	GiftCardObjectResourcePath = "giftCardObject"
)

type GiftCardObjectClient struct {
	Client
}

func NewGiftCardObjectClient(basePath string, client HttpClient) *GiftCardObjectClient {
	return &GiftCardObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GiftCardObjectResourcePath,
		},
	}
}

func (c *GiftCardObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GiftCardObject, error) {
	o := &walletobject.GiftCardObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftCardObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardObjectClient) Get(id string) (*walletobject.GiftCardObject, error) {
	o := &walletobject.GiftCardObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftCardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GiftCardObjectResourcePath,
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

func (c *GiftCardObjectClient) Insert(o *walletobject.GiftCardObject) (*walletobject.GiftCardObject, error) {
	no := &walletobject.GiftCardObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GiftCardObjectResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GiftCardObjectClient) Patch(id string, i interface{}) (*walletobject.GiftCardObject, error) {
	o := &walletobject.GiftCardObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GiftCardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GiftCardObjectClient) Update(id string, o *walletobject.GiftCardObject) (*walletobject.GiftCardObject, error) {
	no := &walletobject.GiftCardObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GiftCardObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
