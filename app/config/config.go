package config

import (
	"strconv"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DBHost       string
	DBPort       int
	DBUser       string
	DBPassword   string
	DBName     	 string
	MaxDBConns   int
}

func New() *Config {
	viper.AutomaticEnv()
	config := &Config{}
	var ok bool
	var err error
	// config values stored as environment values
	config.DBHost, ok = viper.Get("DB_HOST").(string)
	if !ok {
		log.Fatalf("Invalid type assertion DB_HOST")
	}
	config.DBName, ok = viper.Get("DB_NAME").(string)
	if !ok {
		log.Fatalf("Invalid type assertion DB_NAME")
	}
	config.DBUser, ok = viper.Get("DB_USER").(string)
	if !ok {
		log.Fatalf("Invalid type assertion DB_USER")
	}
	config.DBPassword, ok = viper.Get("DB_PW").(string)
	if !ok {
		log.Fatalf("Invalid type assertion DB_PW")
	}
	DBPortString, ok := viper.Get("DB_PORT").(string)
	if !ok {
		log.Fatalf("Invalid type assertion DB_PORT")
	}
	config.DBPort, err = strconv.Atoi(DBPortString)
	if err != nil {
		log.Fatalf("Failed to convert DBPortString to integer")
	}
	MaxDBConnsString, ok := viper.Get("MAX_DB_CONNECTIONS").(string)
	if !ok {
		log.Fatalf("Invalid type assertion MAX_DB_CONNECTIONS")
	}
	config.MaxDBConns, err = strconv.Atoi(MaxDBConnsString)
	if err != nil {
		log.Fatalf("Failed to convert MaxDBConnsString to integer")
	}
	return config
}