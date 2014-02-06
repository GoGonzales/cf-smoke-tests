package config

import (
	"encoding/json"
	"os"
)

type IntegrationConfig struct {
	AppsDomain string `json:"apps_domain"`
}

func Load() (config IntegrationConfig) {
	path := os.Getenv("CONFIG")
	if path == "" {
		panic("Must set $CONFIG to point to an integration config .json file.")
	}

	return loadPath(path)
}

func loadPath(path string) (config IntegrationConfig) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return
}
