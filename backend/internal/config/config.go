/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Environment variables to be set
type Conf struct {
	DatabaseUsername string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
}

func SetSecretFromEnvironmentVariable(fileVariable string) error {
	file, exists := os.LookupEnv(fileVariable)
	if !exists {
		message := fmt.Sprintf("required environment variable %s not set\n", fileVariable)
		return errors.New(message)
	}
	if file == "" {
		message := fmt.Sprintf("required environment variable %s must not be empty\n", fileVariable)
		return errors.New(message)
	}

	secret, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	secretVariable := strings.TrimSuffix(fileVariable, "_FILE")
	err = os.Setenv(secretVariable, string(secret))

	return err
}

func SetSecrets() error {
	secrets := [1]string{"DATABASE_PASSWORD_FILE"}

	for _, secret := range secrets {
		if err := SetSecretFromEnvironmentVariable(secret); err != nil {
			return err
		}
	}

	return nil
}

func Load(path string) (Conf, error) {
	conf := Conf{}
	err := SetSecrets()
	if err != nil {
		return conf, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		message := fmt.Sprintf("Failed to read in Viper configuration: %v", err)
		log.Fatal(message)
		return conf, err
	}

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&conf); err != nil {
		message := fmt.Sprintf("Failed to unmasrhal Viper configuration: %v", err)
		log.Fatal(message)
		return conf, err
	}

	return conf, nil
}
