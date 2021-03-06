package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Prefix          string `json:"prefix"`
	Token           string `json:"token"`
	ImgurID         string `json:"imgur-id"`
	GiphyKey        string `json:"giphy-key"`
	GoogleApi       string `json:"google-search-api"`
	GoogleSearch    string `json:"google-custom-search-engine"`
	MessageLifetime int    `json:"message-lifetime"`
}

func loadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
