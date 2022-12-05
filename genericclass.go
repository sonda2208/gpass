package googlepasses

import (
	"strconv"

	"github.com/Hutchison-Technologies/gpass/walletobjects"
)

const (
	GenericClassResourcePath = "genericClass"
)

type GenericClassClient struct {
	Client
}

func NewGenericClassClient(basePath string, client HTTPClient) *GenericClassClient {
	return &GenericClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: GenericClassResourcePath,
		},
	}
}

func (c *GenericClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.GenericClass, error) {
	o := &walletobject.GenericClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GenericClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericClassClient) Get(id string) (*walletobject.GenericClass, error) {
	o := &walletobject.GenericClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + GenericClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + GenericClassResourcePath,
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

func (c *GenericClassClient) Insert(o *walletobject.GenericClass) (*walletobject.GenericClass, error) {
	no := &walletobject.GenericClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + GenericClassResourcePath,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}

func (c *GenericClassClient) Patch(id string, i interface{}) (*walletobject.GenericClass, error) {
	o := &walletobject.GenericClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + GenericClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(o); err != nil {
		return nil, err
	}

	return o, nil
}

func (c *GenericClassClient) Update(id string, o *walletobject.GenericClass) (*walletobject.GenericClass, error) {
	no := &walletobject.GenericClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + GenericClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     o,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(no); err != nil {
		return nil, err
	}

	return no, nil
}
