package service

import (
	"encoding/json"
	"os"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
)

type IpConfig struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Active   bool   `json:"active"`
}

// Fetch all the IPconfig data from sample json file
func GetIPConfigData() ([]IpConfig, error) {
	// Read the JSON file containing sample data
	jsonData, err := os.ReadFile("ipconfig/data.json")
	if err != nil {
		utilities.ErrorLog("error reading JSON file ", err)
		return nil, err
	}

	// A list of IPConfigs to hold the data
	var configurations []IpConfig

	// Unmarshal the JSON data into the IPConfigs list
	err = json.Unmarshal(jsonData, &configurations)
	if err != nil {
		utilities.ErrorLog("error unmarshalling JSON:", err)
		return nil, err
	}

	return configurations, nil
}
