package googlepasses

import (
	"strconv"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
)

const (
	EventTicketObjectResourcePath = "eventTicketObject"
)

type EventTicketObjectClient struct {
	Client
}

func NewEventTicketObjectClient(basePath string, client HTTPClient) *EventTicketObjectClient {
	return &EventTicketObjectClient{
		Client: Client{
			basePath:     basePath,
			client:       client,
			resourcePath: EventTicketObjectResourcePath,
		},
	}
}

func (c *EventTicketObjectClient) AddMessage(id string, m *walletobject.MessagePayload) (*walletobject.EventTicketObject, error) {
	e := &walletobject.EventTicketObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + EventTicketObjectResourcePath + "/" + id + "/addMessage",
		queryParams: nil,
		payload:     m,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketObjectClient) Get(id string) (*walletobject.EventTicketObject, error) {
	e := &walletobject.EventTicketObject{}
	req := &Request{
		method:      "GET",
		url:         "/" + EventTicketObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     nil,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketObjectClient) List(classID string, maxResults int, paginationToken string) (*walletobject.ListQueryResponse, error) {
	r := &walletobject.ListQueryResponse{}
	req := &Request{
		method:      "GET",
		url:         "/" + EventTicketObjectResourcePath,
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

func (c *EventTicketObjectClient) Insert(e *walletobject.EventTicketObject) (*walletobject.EventTicketObject, error) {
	ne := &walletobject.EventTicketObject{}
	req := &Request{
		method:      "POST",
		url:         "/" + EventTicketObjectResourcePath,
		queryParams: nil,
		payload:     e,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(ne); err != nil {
		return nil, err
	}

	return ne, nil
}

func (c *EventTicketObjectClient) Patch(id string, i interface{}) (*walletobject.EventTicketObject, error) {
	e := &walletobject.EventTicketObject{}
	req := &Request{
		method:      "PATCH",
		url:         "/" + EventTicketObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     i,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(e); err != nil {
		return nil, err
	}

	return e, nil
}

func (c *EventTicketObjectClient) Update(id string, e *walletobject.EventTicketObject) (*walletobject.EventTicketObject, error) {
	ne := &walletobject.EventTicketObject{}
	req := &Request{
		method:      "PUT",
		url:         "/" + EventTicketObjectResourcePath + "/" + id,
		queryParams: nil,
		payload:     e,
		service:     &c.Client,
	}

	if err := req.Do().DecodeResponse(ne); err != nil {
		return nil, err
	}

	return ne, nil
}
