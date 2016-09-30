package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type values map[string]string

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setUp() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func tearDown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testContentType(t *testing.T, r *http.Request, want string) {
	contentType, ok := r.Header["Content-Type"]
	if !ok || len(contentType) < 1 {
		t.Errorf("Content-Type header not present, want: %v", want)
		return
	}
	if got := r.Header["Content-Type"][0]; got != want {
		t.Errorf("Content-type: %v, want %v", got, want)
	}
}

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}
