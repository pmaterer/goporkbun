package goporkbun

import "fmt"

type DNSRecords struct {
	Records []DNSRecord `json:"records"`
}

type DNSRecord struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	TTL      string `json:"ttl"`
	Priority string `json:"prio"`
	Notes    string `json:"notes"`
}

func (c *Client) RetrieveRecords(domain string) (*DNSRecords, error) {
	dnsRetrieveResp := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieve/%s", domain)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, dnsRetrieveResp)
	if err != nil {
		return nil, err
	}

	return dnsRetrieveResp, nil
}

func (c *Client) RetrieveRecordsByID(domain, id string) (*DNSRecords, error) {
	dnsRetrieveResp := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieve/%s/%s", domain, id)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, dnsRetrieveResp)
	if err != nil {
		return nil, err
	}

	return dnsRetrieveResp, nil
}

func (c *Client) RetrieveRecordsBySubdomain(domain, subdomainType, subdomain string) (*DNSRecords, error) {
	dnsRetrieveResp := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieveByNameType/%s/%s/%s", domain, subdomainType, subdomain)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, dnsRetrieveResp)
	if err != nil {
		return nil, err
	}

	return dnsRetrieveResp, nil
}
