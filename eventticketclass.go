package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	EventTicketClassResourcePath = "eventTicketClass"
)

type EventTicketClassClient struct {
	Client
}

func NewEventTicketClassClient(basePath string, client HTTPClient) *EventTicketClassClient {
	return &EventTicketClassClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: EventTicketClassResourcePath,
		},
	}
}

func (c *EventTicketClassClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.EventTicketClass, error) {
	e := &walletobject.EventTicketClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + EventTicketClassResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketClassClient) Get(id string) (*walletobject.EventTicketClass, error) {
	e := &walletobject.EventTicketClass{}
	req := &Request{
		method:      "GET",
		url:         "/" + EventTicketClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketClassClient) List(issuerID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + EventTicketClassResourcePath,
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

func (c *EventTicketClassClient) Insert(e *walletobject.EventTicketClass) (*walletobject.EventTicketClass, error) {
	ne := &walletobject.EventTicketClass{}
	req := &Request{
		method:      "POST",
		url:         "/" + EventTicketClassResourcePath,
		queryParams: nil,
		payload:     e,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(ne); err != nil {
		return nil, err
	}

	return ne, nil
}

func (c *EventTicketClassClient) Patch(id string, i interface{}) (*walletobject.EventTicketClass, error) {
	e := &walletobject.EventTicketClass{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + EventTicketClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketClassClient) Update(id string, e *walletobject.EventTicketClass) (*walletobject.EventTicketClass, error) {
	ne := &walletobject.EventTicketClass{}
	req := &Request{
		method:      "PUT",
		url:         "/" + EventTicketClassResourcePath + "/" + id,
		queryParams: nil,
		payload:     e,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(ne); err != nil {
		return nil, err
	}

	return ne, nil
}
