package goporkbun

import "fmt"

type CreateRecord struct {
	credentials
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	TTL     string `json:"ttl"`
}

func (c *Client) CreateRecord(domain string, record *CreateRecord) (int, error) {
	record.credentials = c.credentials

	path := fmt.Sprintf("/dns/create/%s", domain)

	req, err := c.NewRequest("POST", path, record)
	if err != nil {
		return 0, err
	}

	resp := &struct {
		ID int `json:"id"`
	}{}

	err = c.Do(req, resp)
	if err != nil {
		return 0, err
	}

	return resp.ID, nil
}
