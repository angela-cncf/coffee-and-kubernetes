package env

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config defines this microservice's configuration options
type Config struct {
	REST RESTConfig `json:"rest"`
	DB   DBConfig   `json:"db"`
}

// DBConfig defines the data needed for connecting to a db
type DBConfig struct {
	Enabled         bool   `json:"enabled"`
	Name            string `json:"name"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	VisitorCountKey string `json:"key,omitempty"`
}

// RESTConfig defines a microservice's REST server configuration
type RESTConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	BasePath string `json:"basePath,omitempty"`
}

// CreateConfig parses configFile json and returns the configuration for a microservice
func CreateConfig(configFile string) *Config {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("Unable to open %v: %v", configFile, err.Error())
	} else {
		defer jsonFile.Close()
	}

	byteValue, err := io.ReadAll(io.Reader(jsonFile))
	if err != nil {
		fmt.Printf("Unable to read from %v: %v", configFile, err.Error())
	}

	var env Config
	err = json.Unmarshal([]byte(byteValue), &env)
	if err != nil {
		fmt.Printf("Unable to decode config file into Config struct: %v", err.Error())
	}

	return &env
}
