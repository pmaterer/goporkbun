package goporkbun

import (
	"net/http"
	"reflect"
	"testing"
)

var (
	createRecordResponsePayload              = `{"status":"SUCCESS","id":106926652}`
	retrieveRecordResponsePayload            = `{"status":"SUCCESS","records":[{"id":"106926652","name":"borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""},{"id":"106926659","name":"www.borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""}]}`
	retrieveRecordResponseByIDPayload        = `{"status":"SUCCESS","records":[{"id":"106926652","name":"borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""}]}`
	retrieveRecordResponseBySubdomainPayload = `{"status":"SUCCESS","records":[{"id":"106926659","name":"www.borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""}]}`
	editRecordResponsePayload                = `{"status":"SUCCESS"}`
	deleteRecordResponsePayload              = `{"status":"SUCCESS"}`
)

func TestRetrieveDNSRecords(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/retrieve/borseth.ink", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(retrieveRecordResponsePayload))
		testErrorNil(t, err)
	})

	expectedResponse := &DNSRecords{
		Records: []DNSRecord{
			{
				ID:       "106926652",
				Name:     "borseth.ink",
				Type:     "A",
				Content:  "1.1.1.1",
				TTL:      "300",
				Priority: "0",
				Notes:    "",
			},
			{
				ID:       "106926659",
				Name:     "www.borseth.ink",
				Type:     "A",
				Content:  "1.1.1.1",
				TTL:      "300",
				Priority: "0",
				Notes:    "",
			},
		},
	}

	gotResponse, err := client.RetrieveRecords("borseth.ink")
	testErrorNil(t, err)

	if !reflect.DeepEqual(gotResponse, expectedResponse) {
		testFailNotEqual(t, expectedResponse, gotResponse)
	}
}

func TestRetrieveDNSRecordsByID(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/retrieve/borseth.ink/106926652", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		_, err := rw.Write([]byte(retrieveRecordResponseByIDPayload))
		testErrorNil(t, err)
	})

	expectedResponse := &DNSRecords{
		Records: []DNSRecord{
			{
				ID:       "106926652",
				Name:     "borseth.ink",
				Type:     "A",
				Content:  "1.1.1.1",
				TTL:      "300",
				Priority: "0",
				Notes:    "",
			},
		},
	}

	gotResponse, err := client.RetrieveRecordsByID("borseth.ink", "106926652")
	testErrorNil(t, err)

	if !reflect.DeepEqual(gotResponse, expectedResponse) {
		testFailNotEqual(t, expectedResponse, gotResponse)
	}
}

func TestRetrieveDNSRecordsBySubdomain(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/retrieveByNameType/borseth.ink/A/www", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		_, err := rw.Write([]byte(retrieveRecordResponseBySubdomainPayload))
		testErrorNil(t, err)
	})

	expectedResponse := &DNSRecords{
		Records: []DNSRecord{
			{
				ID:       "106926659",
				Name:     "www.borseth.ink",
				Type:     "A",
				Content:  "1.1.1.1",
				TTL:      "300",
				Priority: "0",
				Notes:    "",
			},
		},
	}

	gotResponse, err := client.RetrieveRecordsBySubdomain("borseth.ink", "A", "www")
	testErrorNil(t, err)

	if !reflect.DeepEqual(gotResponse, expectedResponse) {
		testFailNotEqual(t, expectedResponse, gotResponse)
	}
}

func TestCreateRecord(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/create/borseth.ink", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(createRecordResponsePayload))
		testErrorNil(t, err)
	})

	newRecord := &DNSRecordOptions{
		Name:    "www",
		Type:    "A",
		Content: "127.0.0.1",
		TTL:     "300",
	}

	expectedResponse := "106926652"
	gotResponse, err := client.CreateRecord("borseth.ink", newRecord)
	testErrorNil(t, err)

	if expectedResponse != gotResponse {
		testFailNotEqual(t, expectedResponse, gotResponse)
	}
}

func TestEditRecord(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/edit/borseth.ink/106926652", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "UserAgent", "")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(editRecordResponsePayload))
		testErrorNil(t, err)
	})

	editRecord := &DNSRecordOptions{
		Name:    "www",
		Type:    "A",
		Content: "127.0.0.1",
		TTL:     "300",
	}

	err := client.EditRecord("borseth.ink", "106926652", editRecord)
	testErrorNil(t, err)
}

func TestEditRecordsBySubdomain(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/editByNameType/borseth.ink/A/www", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(editRecordResponsePayload))
		testErrorNil(t, err)
	})

	editRecord := &DNSSubdomainOptions{
		Content: "127.0.0.1",
		TTL:     "600",
	}

	err := client.EditRecordsBySubdomain("borseth.ink", "www", "A", editRecord)
	testErrorNil(t, err)
}

func TestDeleteRecord(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/delete/borseth.ink/106926652", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(deleteRecordResponsePayload))
		testErrorNil(t, err)
	})

	err := client.DeleteRecord("borseth.ink", "106926652")
	testErrorNil(t, err)
}

func TestDeleteRecordsBySubdomain(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/deleteByNameType/borseth.ink/A/www", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		testCredentials(t, r, client.credentials)
		_, err := rw.Write([]byte(deleteRecordResponsePayload))
		testErrorNil(t, err)
	})

	err := client.DeleteRecordBySubdomain("borseth.ink", "www", "A")
	testErrorNil(t, err)
}
