package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	FlightClassResourcePath = "flightClass"
)

type FlightClassClient struct {
	Client
}

func NewFlightClassClient(basePath string, client HTTPClient) *FlightClassClient {
	return &FlightClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: FlightClassResourcePath,
		},
	}
}

func (c *FlightClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.FlightClass, error) {
	f := &walletobject.FlightClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + FlightClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightClassClient) Get(id string) (*walletobject.FlightClass, error) {
	f := &walletobject.FlightClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + FlightClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + FlightClassResourcePath,
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

func (c *FlightClassClient) Insert(f *walletobject.FlightClass) (*walletobject.FlightClass, error) {
	nf := &walletobject.FlightClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + FlightClassResourcePath,
		queryParams: nil,
		payload:     f,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nf); err != nil {
		return nil, err
	}

	return nf, nil
}

func (c *FlightClassClient) Patch(id string, i interface{}) (*walletobject.FlightClass, error) {
	f := &walletobject.FlightClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + FlightClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *FlightClassClient) Update(id string, f *walletobject.FlightClass) (*walletobject.FlightClass, error) {
	nf := &walletobject.FlightClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + FlightClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     f,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(nf); err != nil {
		return nil, err
	}

	return nf, nil
}
