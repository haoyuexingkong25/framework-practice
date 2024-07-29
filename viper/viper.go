package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

const Config_TYPE = "yaml"

func GetViper(name, path string) {
	viper.SetConfigName(name)        // name of config file (without extension)
	viper.SetConfigType(Config_TYPE) // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)        // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
