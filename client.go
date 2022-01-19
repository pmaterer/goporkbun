package goporkbun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	Version          = "0.0.1"
	defaultBaseURL   = "https://porkbun.com/api/json/v3"
	defaultUserAgent = "goporkbun/" + Version
)

type credentials struct {
	SecretKey string `json:"secretapikey"`
	Key       string `json:"apikey"`
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string

	userAgent string
	credentials
}

func NewClient(apiKey, secretAPIKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		BaseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
		credentials: credentials{
			SecretKey: secretAPIKey,
			Key:       apiKey,
		},
	}
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+path, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		return err
	}

	return nil
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func checkResponse(resp *http.Response) error {
	if statusCode := resp.StatusCode; statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
		return nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("page not found: %s", resp.Request.URL.Path)
	}

	var errResp errorResponse

	err := json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return fmt.Errorf("could not decode error response: %s", err)
	}

	return &Error{
		Message: errResp.Message,
	}
}
