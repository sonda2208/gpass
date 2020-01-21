package gpass

import "github.com/sonda2208/gpass/walletobjects"

type Message struct {
	Body   string
	Header string
}

func (m *Message) toWO() *walletobjects.Message {
	return &walletobjects.Message{
		Body:        m.Body,
		Header:      m.Header,
		Kind:        "walletobjects#walletObjectMessage",
		MessageType: "TEXT",
	}
}

func woToMessage(m *walletobjects.Message) *Message {
	return &Message{
		Body:   m.Body,
		Header: m.Header,
	}
}

type AddMessageRequest struct {
	Message *Message
}

func (amr *AddMessageRequest) toWO() *walletobjects.AddMessageRequest {
	return &walletobjects.AddMessageRequest{
		Message: amr.Message.toWO(),
	}
}

func woToAddMessageRequest(amr *walletobjects.AddMessageRequest) *AddMessageRequest {
	return &AddMessageRequest{
		Message: woToMessage(amr.Message),
	}
}
