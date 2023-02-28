package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadConfig() map[string]string {
	var config map[string]string
	bytes, err := ioutil.ReadFile("./../config/config.json")
	if err != nil {
		log.Fatalf("failed to read config file")
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("failed to unmarshall config")
	}
	return config
}
