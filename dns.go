package goporkbun

import (
	"fmt"
	"strconv"
)

const (
	MinimumTTLValue = 300
)

type DNSRecordType string

const (
	RecordTypeA     DNSRecordType = "A"
	RecordTypeMX    DNSRecordType = "MX"
	RecordTypeCNAME DNSRecordType = "CNAME"
	RecordTypeALIAS DNSRecordType = "ALIAS"
	RecordTypeTXT   DNSRecordType = "TXT"
	RecordTypeNS    DNSRecordType = "NS"
	RecordTypeAAAA  DNSRecordType = "AAAA"
	RecordTypeSRV   DNSRecordType = "SRV"
	RecordTypeTLSA  DNSRecordType = "TLSA"
	RecordTypeCAA   DNSRecordType = "CAA"
)

// DNSRecords represents a collection of DNS records
type DNSRecords struct {
	Records []DNSRecord `json:"records"`
}

// DNSRecord represents a DNS record
type DNSRecord struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Type     DNSRecordType `json:"type"`
	Content  string        `json:"content"`
	TTL      string        `json:"ttl"`
	Priority string        `json:"prio"`
	Notes    string        `json:"notes"`
}

func (c *Client) RetrieveRecords(domain string) (*DNSRecords, error) {
	records := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieve/%s", domain)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (c *Client) RetrieveRecordsByID(domain, id string) (*DNSRecords, error) {
	records := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieve/%s/%s", domain, id)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (c *Client) RetrieveRecordsBySubdomain(domain, subdomainType, subdomain string) (*DNSRecords, error) {
	records := &DNSRecords{}

	path := fmt.Sprintf("/dns/retrieveByNameType/%s/%s/%s", domain, subdomainType, subdomain)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return nil, err
	}

	err = c.Do(req, records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// DNSRecordOptions is used with CreateRecord and UpdateRecord
// to create and update DNS records
type DNSRecordOptions struct {
	credentials
	Name     string        `json:"name"`
	Type     DNSRecordType `json:"type"`
	Content  string        `json:"content"`
	TTL      string        `json:"ttl"`
	Priority string        `json:"prio"`
}

func (o *DNSRecordOptions) Validate() error {
	ttl, err := strconv.Atoi(o.TTL)
	if err != nil {
		return fmt.Errorf("invalid ttl: value must be an integer")
	}

	if ttl < MinimumTTLValue {
		return fmt.Errorf("invalid ttl: minimum value is 300, got: %d", ttl)
	}

	return nil
}

func (c *Client) CreateRecord(domain string, newRecord *DNSRecordOptions) (string, error) {
	newRecord.credentials = c.credentials

	if err := newRecord.Validate(); err != nil {
		return "", err
	}

	path := fmt.Sprintf("/dns/create/%s", domain)

	req, err := c.NewRequest("POST", path, newRecord)
	if err != nil {
		return "", err
	}

	resp := &struct {
		ID int `json:"id"`
	}{}

	err = c.Do(req, resp)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(resp.ID), nil
}

func (c *Client) EditRecord(domain, id string, recordEdit *DNSRecordOptions) error {
	recordEdit.credentials = c.credentials

	if err := recordEdit.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf("/dns/edit/%s/%s", domain, id)

	req, err := c.NewRequest("POST", path, recordEdit)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// DNSSubdomainOptions is used with EditRecordsBySubdomain
// to update all records of a particular subdomain
type DNSSubdomainOptions struct {
	credentials
	Content string `json:"content"`
	TTL     string `json:"ttl"`
}

func (c *Client) EditRecordsBySubdomain(domain, subdomain string, recordType DNSRecordType, subdomainEditOptions *DNSSubdomainOptions) error {
	subdomainEditOptions.credentials = c.credentials

	path := fmt.Sprintf("/dns/editByNameType/%s/%s/%s", domain, recordType, subdomain)

	req, err := c.NewRequest("POST", path, subdomainEditOptions)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteRecord(domain, id string) error {
	path := fmt.Sprintf("/dns/delete/%s/%s", domain, id)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteRecordBySubdomain(domain, subdomain string, recordType DNSRecordType) error {
	path := fmt.Sprintf("/dns/deleteByNameType/%s/%s/%s", domain, recordType, subdomain)

	req, err := c.NewRequest("POST", path, c.credentials)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
