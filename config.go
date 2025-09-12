package main

import (
	"encoding/json"
	"os"
)

type config struct {
	SourceDir    string   `json:"source-dir"`
	ProcessedDir string   `json:"processed-dir"`
	AllowedExts  []string `json:"allowed-exts"`
	Labels       []string `json:"labels"`
}

func loadConfig(path string) (config, error) {
	var config config

	data, err := os.ReadFile(path)

	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
