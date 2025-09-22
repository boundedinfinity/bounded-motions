package model

import (
	"encoding/json"
	"os"
)

type ConfigJson struct {
	Keybindings KeyBindingJson   `json:"key-binding"`
	Style       ConfigJsonStyles `json:"style"`
}

type ConfigJsonStyles struct {
	NavigationBox    ConfigJsonStyle `json:"navigation-box"`
	NavigationPath   ConfigJsonStyle `json:"navigation-path"`
	NavigationOption ConfigJsonStyle `json:"navigation-option"`
}

type ConfigJsonStyle struct {
	ForegroundColor string `json:"foreground-color"`
	Padding         int    `json:"padding"`
	Margin          int    `json:"margin"`
}

func LoadConfig(path string, config *ConfigJson) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}
