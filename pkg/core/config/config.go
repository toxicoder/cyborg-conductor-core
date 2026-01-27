package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config represents the application configuration
type Config struct {
	// MaxContextBytes defines the maximum context size in bytes per job
	MaxContextBytes int64 `json:"max_context_bytes"`
}

// LoadConfig loads a Config message from JSON/YAML or any key-value map
func LoadConfig(source interface{}) (*Config, error) {
	var config Config
	
	switch v := source.(type) {
	case string:
		// Assume it's JSON
		err := json.Unmarshal([]byte(v), &config)
		if err != nil {
			return nil, err
		}
	case []byte:
		// Assume it's JSON
		err := json.Unmarshal(v, &config)
		if err != nil {
			return nil, err
		}
	case map[string]interface{}:
		// Convert map to JSON then unmarshal
		jsonData, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonData, &config)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported config source type: %T", v)
	}
	
	return &config, nil
}

// LoadFromFile loads configuration from a file
func LoadFromFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	return LoadConfig(data)
}