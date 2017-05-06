package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	DenonIP   string `json:"denon_ip"`
	DenonPort int    `json:"denon_port"`
	Port      int    `json:"port"`
}

func parseConfig(configPath string) (Config, error) {
	var config Config

	confB, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, errors.New("Could not read config file")
	}

	err = json.Unmarshal(confB, &config)
	if err != nil {
		return config, errors.New("Could not parse config")
	}

	return config, nil
}
