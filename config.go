package megaprem_bot

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Token string `json:"token"`
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
