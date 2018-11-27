package googlepasses

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	GooglePayAPIBasePath = "https://www.googleapis.com/walletobjects/v1"
	GooglePayAPIScope    = "https://www.googleapis.com/auth/wallet_object.issuer"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	basePath     string
	resourcePath string
	client       HTTPClient
}

type Request struct {
	method      string
	url         string
	queryParams *QueryParams
	payload     interface{}
	service     *Client
}

func (r *Request) Do() *Response {
	var body io.Reader
	if r.payload != nil {
		data, err := json.Marshal(r.payload)
		if err != nil {
			return &Response{
				err: err,
			}
		}

		body = bytes.NewReader(data)
	}

	if r.queryParams != nil {
		r.url += "?" + r.queryParams.Encode()
	}

	r.url = r.service.basePath + r.url
	req, err := http.NewRequest(r.method, r.url, body)
	if err != nil {
		return &Response{
			err: err,
		}
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := r.service.client.Do(req)
	if err != nil {
		return &Response{
			res: res,
			err: err,
		}
	}

	return &Response{
		res: res,
	}
}

type Response struct {
	res *http.Response
	err error
}

func (r *Response) DecodeResponse(target interface{}) error {
	if r.err != nil {
		return r.err
	}

	if r.res != nil && r.res.StatusCode != http.StatusOK {
		if r.res.ContentLength <= 0 {
			if r.res.StatusCode != http.StatusOK {
				return &Error{
					Code: r.res.StatusCode,
				}
			}
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.res.Body)

		resError := &Error{}
		_ = json.Unmarshal(body.Bytes(), resError)
		return resError
	}

	if target == nil {
		return nil
	}

	if r.res.StatusCode == http.StatusNoContent || r.res.Body == nil {
		return nil
	}

	return json.NewDecoder(r.res.Body).Decode(target)
	// hello world
}
