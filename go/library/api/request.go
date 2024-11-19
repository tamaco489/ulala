package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	baseURL   string
	apiSecret string
)

func init() {
	baseURL = "http://localhost:8080"
	apiSecret = "ABCDEFG123456789"
}

type Request struct {
	Method string
	Path   string
	Data   interface{}
}

func NewRequest(method, path string, data interface{}) *Request {
	return &Request{
		Method: method,
		Path:   path,
		Data:   data,
	}
}

func NewRequestWithContextAndIncludeBody(ctx context.Context, r *Request) (*http.Request, error) {
	b, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, r.Path, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func SendRequest(ctx context.Context, r *Request) ([]byte, error) {
	var req *http.Request
	var err error
	var url = baseURL + r.Path
	r.Path = url

	switch r.Method {
	case http.MethodGet:
		req, err = http.NewRequestWithContext(ctx, r.Method, url, nil)
		if err != nil {
			return nil, err
		}

	case http.MethodPost:
		req, err = NewRequestWithContextAndIncludeBody(ctx, r)
		if err != nil {
			return nil, err
		}

	case http.MethodPatch:
		req, err = NewRequestWithContextAndIncludeBody(ctx, r)
		if err != nil {
			return nil, err
		}

	case http.MethodDelete:
		req, err = NewRequestWithContextAndIncludeBody(ctx, r)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Invalid request parameter")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiSecret)

	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
