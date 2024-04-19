package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// LoadSettings function that load the settings ini data
func LoadSettings() *ini.File {
	cfg, err := ini.Load("config/settings.ini")
	if err != nil {
		panic(fmt.Errorf("error to load the file setting.ini: %v", err))
	}
	return cfg
}
