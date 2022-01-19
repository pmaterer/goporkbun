package goporkbun

type pingResponse struct {
	YourIP string `json:"yourIp"`
}

func (c *Client) Ping() (bool, error) {
	req, err := c.NewRequest("POST", "/ping", c.credentials)
	if err != nil {
		return false, err
	}

	resp := &pingResponse{}

	if err := c.Do(req, resp); err != nil {
		return false, err
	}

	return true, nil
}
