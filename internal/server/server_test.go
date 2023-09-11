package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/service"
)

func TestIntegrationAPIHandler_success(t *testing.T) {
	// set up any necessary dependencies
	os.WriteFile(".env", []byte("GO_ENV='development'\nAPP_HOST=''\nAPP_PORT='8082'\nAPP_LOG_LEVEL='DEBUG'\nX='1'"), 0755)

	srv, err := NewServer()
	if err != nil {
		t.Error("unable to get the API handler")
	}
	go func() {
		StartServer(srv)
	}()

	fileHelper(t)
	// Create a request to test the handler
	req := httptest.NewRequest("GET", "/hosts/inefficient", nil)
	rr := httptest.NewRecorder()

	t.Setenv("X", "1")
	// Send the request to your test server
	srv.Handler.ServeHTTP(rr, req)

	// Check the response status code (example: expect 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
	if !strings.EqualFold(rr.Body.String(), "{\"hostnames\":[\"mta-prod-7\"]}") {
		t.Errorf("\nexpected: %v\ngot: %v", "{\"hostnames\":[\"mta-prod-7\"]}", rr.Body.String())
	}

	// Clean up the storage directory(for log purposes) used in the integration test
	os.RemoveAll("ipconfig")
	os.Remove(".env")
	os.RemoveAll("storage")
}

func fileHelper(t *testing.T) {
	t.Helper()
	data := []service.IpConfig{
		{IP: "127.0.0.1", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-8", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.6", Hostname: "mta-prod-7", Active: true},
		{IP: "127.0.0.7", Hostname: "mta-prod-7", Active: false},
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		t.Error()
	}

	os.MkdirAll("ipconfig", os.ModePerm)
	err = os.WriteFile("ipconfig/data.json", dataBytes, 0755)
	if err != nil {
		t.Error("unable to write file", err)
	}
}
