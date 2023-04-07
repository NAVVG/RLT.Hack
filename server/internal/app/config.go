package app

import (
	"encoding/json"
	"os"
)

type Config struct {
	Postgres struct {
		URL string `json:"url"`
	} `json:"postgres"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := new(Config)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
