package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token           string `json:"token"`
	ImgurID         string `json:"imgur-id"`
	GiphyKey        string `json:"giphy-key"`
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
