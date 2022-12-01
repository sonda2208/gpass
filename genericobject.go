package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	GenericObjectResourcePath = "genericObject"
)

type GenericObjectClient struct {
	Client
}

func NewGenericObjectClient(basePath string, client HTTPClient) *GenericObjectClient {
	return &GenericObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GenericObjectResourcePath,
		},
	}
}

func (c *GenericObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GenericObject, error) {
	o := &walletobject.GenericObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GenericObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericObjectClient) Get(id string) (*walletobject.GenericObject, error) {
	o := &walletobject.GenericObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + GenericObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GenericObjectResourcePath,
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

func (c *GenericObjectClient) Insert(o *walletobject.GenericObject) (*walletobject.GenericObject, error) {
	no := &walletobject.GenericObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + GenericObjectResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GenericObjectClient) Patch(id string, i interface{}) (*walletobject.GenericObject, error) {
	o := &walletobject.GenericObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GenericObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericObjectClient) Update(id string, o *walletobject.GenericObject) (*walletobject.GenericObject, error) {
	no := &walletobject.GenericObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GenericObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
