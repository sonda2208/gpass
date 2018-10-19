package googlepasses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

type ClientMock struct {
	req *http.Request
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return &http.Response{
		StatusCode: 200,
	}, nil
}

func TestServiceRequestAndResponse(t *testing.T) {
	t.Run("With payload", func(t *testing.T) {
		jsonData := `{"message":{"body":"this is a body","header":"this is a header"}}`
		msg := &walletobject.MessagePayload{}
		err := json.Unmarshal([]byte(jsonData), msg)
		if err != nil {
			t.Error(errors.WithStack(err))
		}

		client := &ClientMock{}
		svc := Client{
			client: client,
		}
		req := Request{
			method:      "GET",
			url:         "",
			queryParams: nil,
			payload:     msg,
			service:     &svc,
		}

		err = req.Do().DecodeResponse(msg)
		if err != nil {
			t.Error(errors.WithStack(err))
		}

		body, err := client.req.GetBody()
		if err != nil {
			t.Error(errors.WithStack(err))
		}

		reqPayload := new(bytes.Buffer)
		reqPayload.ReadFrom(body)
		assert.Equal(t, jsonData, reqPayload.String())
	})

	t.Run("Without payload", func(t *testing.T) {
		client := &ClientMock{}
		svc := Client{
			client: client,
		}
		req := Request{
			method:      "GET",
			url:         "",
			queryParams: nil,
			payload:     nil,
			service:     &svc,
		}

		err := req.Do().DecodeResponse(nil)
		if err != nil {
			t.Error(errors.WithStack(err))
		}
	})
}
