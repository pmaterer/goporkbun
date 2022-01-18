package goporkbun

import (
	"net/http"
	"testing"
)

var createRecordResponse = `{"status":"SUCCESS","id":106926652}`

func TestCreateRecord(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dns/create/borseth.ink", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Content-Type", "application/json")
		_, err := rw.Write([]byte(createRecordResponse))
		testErrorNil(t, err)
	})

	newRecord := CreateRecord{
		Name:    "www",
		Type:    "A",
		Content: "127.0.0.1",
		TTL:     "300",
	}

	expectedResponse := 106926652
	gotResponse, err := client.CreateRecord("borseth.ink", &newRecord)
	testErrorNil(t, err)

	if expectedResponse != gotResponse {
		testFailNotEqual(t, expectedResponse, gotResponse)
	}
}
