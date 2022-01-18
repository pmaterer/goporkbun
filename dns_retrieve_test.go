package goporkbun

import (
	"net/http"
	"reflect"
	"testing"
)

var dnsRetrieveResponse = `{"status":"SUCCESS","records":[{"id":"106926652","name":"borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""},{"id":"106926659","name":"www.borseth.ink","type":"A","content":"1.1.1.1","ttl":"300","prio":"0","notes":""}]}`

func TestRetrieveDNSRecords(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/dns/retrieve/borseth.ink", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Content-Type", "application/json")
		_, err := rw.Write([]byte(dnsRetrieveResponse))
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
