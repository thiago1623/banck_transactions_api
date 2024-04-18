package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func LoadSettings() *ini.File {
	cfg, err := ini.Load("config/settings.ini")
	if err != nil {
		panic(fmt.Errorf("Error al cargar el archivo setting.ini: %v", err))
	}
	return cfg
}
