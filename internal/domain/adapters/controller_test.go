package adapters

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	ts := httptest.NewServer(NewAppMux())
	defer ts.Close()

	_, body := testRequest(t, ts, "GET", "/healthcheck", nil)
	if body != getExpectedJsonString("OK") {
		t.Fatalf(body)
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}

func getExpectedJsonString(body any) string {
	bytes, _ := json.Marshal(body)

	return string(bytes)
}
