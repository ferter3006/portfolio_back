package config

import (
	"encoding/json"
	"os"
)

type RedsysConfig struct {
	MerchantCode string `json:"merchant_code"`
	Terminal    string `json:"terminal"`
	SecretKey   string `json:"secret_key"`
	Currency    string `json:"currency"`
	TestURL     string `json:"test_url"`
}

func LoadRedsysConfig(path string) (*RedsysConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg RedsysConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
