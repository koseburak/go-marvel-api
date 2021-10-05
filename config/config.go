package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Configuration keeps the os env variables data structure
type Configuration struct {
	MarvelPrivateKey string `mapstructure:"marvel_private_key"`
	MarvelPublicKey  string `mapstructure:"marvel_public_key"`
	MarvelAPIBaseURL string `mapstructure:"marvel_api_base_url"`
}

var envVars = []string{
	"MARVEL_PUBLIC_KEY",
	"MARVEL_PRIVATE_KEY",
	"MARVEL_API_BASE_URL",
}

// Config init and load the os env variables using viper package
func Config() (*Configuration, error) {

	// viper config/env key replacer configuration
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Bind viper keys to env vars
	for _, v := range envVars {
		if err := viper.BindEnv(v); err != nil {
			return nil, err
		}
	}

	// unmarshall the configuration structure
	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// set env config. file name, type and path
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// set every one of env variables into os
	for _, v := range envVars {
		os.Setenv(v, viper.GetString(v))
	}

	return &config, nil
}
