package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/service"
	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
	"github.com/gorilla/mux"
)

var (
	SUCCESS_MSG = map[string]interface{}{
		"status": "success",
	}
	FAILURE_MSG = map[string]interface{}{
		"status": "failure",
		"error":  nil,
	}
)

type apiHandler struct {
	IPConfigs []service.IpConfig
}

func NewAPIHandler() *apiHandler {
	return &apiHandler{}
}

func optimizerHandler(r *mux.Router) {
	oh := NewAPIHandler()

	r.HandleFunc("/inefficient", oh.getInefficientHosts).Methods("GET")
}

// Get the list of inefficient hosts
func (ah *apiHandler) getInefficientHosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if ah.IPConfigs == nil || len(ah.IPConfigs) == 0 {
		ipConfigs, err := service.GetIPConfigData()
		if err != nil {
			utilities.ErrorLog("unable to get ip configurations from mock service")
			return
		}
		ah.IPConfigs = make([]service.IpConfig, len(ipConfigs))
		copy(ah.IPConfigs, ipConfigs)
	}

	// retrieve the list of inefficient hostnames
	hostsFreq := make(map[string]int)
	hostnames, envValue := make([]string, 0), utilities.GetEnvValue("X", "1")
	utilities.InfoLog("Value of X is:", envValue)

	X, err := strconv.Atoi(envValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getFailureMessage(errors.New("env X should be a digit or number")))
		return
	}

	// Maintain a frequency count for each hostname using hostsFreq map
	for _, ipConfig := range ah.IPConfigs {
		if ipConfig.Active {
			hostsFreq[ipConfig.Hostname]++
		} else if _, ok := hostsFreq[ipConfig.Hostname]; !ok {
			hostsFreq[ipConfig.Hostname] = 0
		}
	}

	// Iterate over hostsFreq map to find the list of hosts having less or X active IPs
	for hostname, freq := range hostsFreq {
		if freq <= X {
			hostnames = append(hostnames, hostname)
		}
	}

	res := map[string][]string{
		"hostnames": hostnames,
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getFailureMessage(errors.New("error while marshalling hostnames")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func getFailureMessage(err error) []byte {
	// Append error to failure message
	FAILURE_MSG["error"] = err.Error()

	data, err := json.Marshal(FAILURE_MSG)
	if err != nil {
		return nil
	}
	return data
}
