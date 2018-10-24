package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	FlightObjectResourcePath = "flightObject"
)

type FlightObjectClient struct {
	Client
}

func NewFlightObjectClient(basePath string, client HTTPClient) *FlightObjectClient {
	return &FlightObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: FlightObjectResourcePath,
		},
	}
}

func (c *FlightObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.FlightObject, error) {
	f := &walletobject.FlightObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + FlightObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightObjectClient) Get(id string) (*walletobject.FlightObject, error) {
	f := &walletobject.FlightObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + FlightObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + FlightObjectResourcePath,
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

func (c *FlightObjectClient) Insert(f *walletobject.FlightObject) (*walletobject.FlightObject, error) {
	nf := &walletobject.FlightObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + FlightObjectResourcePath,
		queryParams: nil,
		payload:     f,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nf); err != nil {
		return nil, err
	}

	return nf, nil
}

func (c *FlightObjectClient) Patch(id string, i interface{}) (*walletobject.FlightObject, error) {
	f := &walletobject.FlightObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + FlightObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightObjectClient) Update(id string, f *walletobject.FlightObject) (*walletobject.FlightObject, error) {
	nf := &walletobject.FlightObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + FlightObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     f,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nf); err != nil {
		return nil, err
	}

	return nf, nil
}
