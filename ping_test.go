package goporkbun

import (
	"net/http"
	"reflect"
	"testing"
)

var pingResponseJSON = `{"status":"SUCCESS","yourIp":"127.0.0.1"}`

func TestPing(t *testing.T) {
	setup()

	defer teardown()

	mux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		testMethod(t, r)
		testHeader(t, r, "Content-Type", "application/json")
		_, err := rw.Write([]byte(pingResponseJSON))
		testErrorNil(t, err)
	})

	expectedReponse := true

	gotResponse, err := client.Ping()

	testErrorNil(t, err)

	if !reflect.DeepEqual(gotResponse, expectedReponse) {
		testFailNotEqual(t, expectedReponse, gotResponse)
	}
}
