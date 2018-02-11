package sbankenSDK

import (
	"github.com/larwef/sbankenSDK/authentication"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Returns a configured client for testing http calls. Handler function is set in test function
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	server := httptest.NewServer(apiHandler)

	config := Config{
		AccountsEndpoint:     server.URL + "/",
		TransactionsEndpoint: server.URL + "/",
		TransfersEndpoint:    server.URL + "/",
	}
	client = NewClient(&http.Client{}, config, authentication.SbankenToken{})

	return client, mux, server.Close
}

func getTestFileAsString(t *testing.T, filepath string) string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if actual := r.Method; actual != expected {
		t.Errorf("Request method: %v, expected %v", actual, expected)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, expected string) {
	if actual := r.Header.Get(header); actual != expected {
		t.Errorf("Header.Get(%q) returned %q, expected %q", header, actual, expected)
	}
}

func assertNotError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Got unexpected error: %s", err)
	}
}

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Expected %v %v to be equal to %v %v", reflect.TypeOf(actual).Name(), actual, reflect.TypeOf(expected).Name(), expected)
	}
}
