package goporkbun

type PingResponse struct {
	YourIP string `json:"yourIp"`
}

func (c *Client) Ping() (*PingResponse, error) {
	req, err := c.NewRequest("POST", "/ping", c.credentials)
	if err != nil {
		return nil, err
	}

	pingResp := &PingResponse{}

	if err := c.Do(req, pingResp); err != nil {
		return nil, err
	}

	return pingResp, nil
}
