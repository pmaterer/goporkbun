package goporkbun

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func testFailNotEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	t.Fatalf("Not equal: \n"+
		"expected : %v\n"+
		"actual   : %v", expected, actual)
}

func testErrorNil(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("abc123", "xzy789")
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request) {
	t.Helper()

	if got := r.Method; got != http.MethodPost {
		t.Errorf("Request method: %v, want %v", got, http.MethodPost)
	}
}

func testHeader(t *testing.T, r *http.Request, header, want string) {
	t.Helper()

	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testCredentials(t *testing.T, r *http.Request, want credentials) {
	t.Helper()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Could not read request body: %v", err)
	}

	got := &credentials{}

	err = json.Unmarshal(body, got)
	if err != nil {
		t.Fatalf("Could not unmarshal request credentials: %v", err)
	}

	if want.Key != got.Key {
		t.Fatalf("Got wrong key: %v, want: %v", got.Key, want.Key)
	}

	if want.SecretKey != got.SecretKey {
		t.Fatalf("Got wrong secret key: %v, want: %v", got.SecretKey, want.SecretKey)
	}
}

func TestClient_SetUserAgent(t *testing.T) {
	setup()

	defer teardown()

	expected := "terraform-provider-porkbun/1.0.0"
	client.SetUserAgent(expected)
	got := client.userAgent

	if client.userAgent != expected {
		testFailNotEqual(t, expected, got)
	}
}
