package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/service"
	"github.com/gorilla/mux"
)

func TestGetInefficientHosts_success_OK(t *testing.T) {
	router := mux.NewRouter()

	handler := NewAPIHandler()
	handler.IPConfigs = []service.IpConfig{
		{IP: "127.0.0.1", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-8", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.6", Hostname: "mta-prod-7", Active: true},
		{IP: "127.0.0.7", Hostname: "mta-prod-7", Active: false},
	}

	router.HandleFunc("/hosts/inefficient", handler.getInefficientHosts).Methods("GET")

	// Create a request to test the handler
	req := httptest.NewRequest("GET", "/hosts/inefficient", nil)
	rr := httptest.NewRecorder()

	t.Setenv("X", "1")
	// Send the request to your test server
	router.ServeHTTP(rr, req)

	// Check the response status code (example: expect 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
	if !strings.EqualFold(rr.Body.String(), "{\"hostnames\":[\"mta-prod-7\"]}") {
		t.Error(rr.Body.String())
	}

	os.RemoveAll("storage")
}

func TestGetInefficientHosts_incorrectX(t *testing.T) {
	router := mux.NewRouter()

	handler := NewAPIHandler()
	handler.IPConfigs = []service.IpConfig{
		{IP: "127.0.0.1", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-8", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-5", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-8", Active: true},
		{IP: "127.0.0.6", Hostname: "mta-prod-7", Active: true},
	}

	router.HandleFunc("/hosts/inefficient", handler.getInefficientHosts).Methods("GET")

	t.Setenv("X", "C")
	// Create a request to test the handler
	req := httptest.NewRequest("GET", "/hosts/inefficient", nil)
	rr := httptest.NewRecorder()

	// Send the request to your test server
	router.ServeHTTP(rr, req)

	// Check the response status code (example: expect 200 OK)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
	}

	os.RemoveAll("storage")
}
func TestGetInefficientHosts_with_env_file(t *testing.T) {
	// set up any necessary dependencies
	fileHelper(t)
	router, err := New()
	if err != nil {
		t.Error(err)
	}

	// Create a request to test the handler
	req := httptest.NewRequest("GET", "/hosts/inefficient", nil)
	rr := httptest.NewRecorder()

	t.Setenv("X", "1")
	// Send the request to your test server
	router.ServeHTTP(rr, req)

	// Check the response status code (example: expect 200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
	if !strings.EqualFold(rr.Body.String(), "{\"hostnames\":[\"mta-prod-7\"]}") {
		t.Error(rr.Body.String())
	}

	os.RemoveAll("storage")
	os.RemoveAll("ipconfig")
}

// func TestGetInefficientHosts_with_incorrect_marshal(t *testing.T) {
// 	router := mux.NewRouter()

// 	handler := NewAPIHandler()
// 	handler.IPConfigs = []service.IpConfig{
// 		{IP: "127.0.0.1", Hostname: "mta-prod-5", Active: true},
// 		{IP: "127.0.0.2", Hostname: "mta-prod-8", Active: false},
// 		{IP: "127.0.0.3", Hostname: "mta-prod-8", Active: true},
// 		{IP: "127.0.0.4", Hostname: "mta-prod-5", Active: true},
// 		{IP: "127.0.0.5", Hostname: "mta-prod-8", Active: true},
// 		{IP: "127.0.0.6", Hostname: "mta-prod-7", Active: true},
// 	}

// 	router.HandleFunc("/hosts/inefficient", handler.getInefficientHosts).Methods("GET")

// 	t.Setenv("X", "1")
// 	// Create a request to test the handler
// 	req := httptest.NewRequest("GET", "/hosts/inefficient", nil)
// 	rr := httptest.NewRecorder()

// 	// Send the request to your test server
// 	router.ServeHTTP(rr, req)

// 	// Check the response status code (example: expect 200 OK)
// 	if rr.Code != http.StatusInternalServerError {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
// 	}

// 	os.RemoveAll("storage")
// }

func TestPanicRecoveryMiddleware_panic_happens(t *testing.T) {
	handler := PanicRecoveryMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("panic occurred")
	}))

	// Create a request to trigger the panic
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Run the request through the test server
	handler.ServeHTTP(rr, req)

	// Check the response status code (expect 500 Internal Server Error)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
	}

	os.RemoveAll("storage")
}

// Creates the sample data.json file in folder ipconfig for mock tests
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
