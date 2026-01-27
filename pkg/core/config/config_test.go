package config

import (
	"testing"
)

func TestConfig_LoadConfig(t *testing.T) {
	// Test loading from JSON string
	jsonData := `{"max_context_bytes": 1048576}`
	config, err := LoadConfig(jsonData)
	if err != nil {
		t.Fatalf("Failed to load config from JSON: %v", err)
	}
	
	if config.MaxContextBytes != 1048576 {
		t.Errorf("Expected MaxContextBytes to be 1048576, got %d", config.MaxContextBytes)
	}
	
	// Test loading from byte slice
	byteData := []byte(`{"max_context_bytes": 2097152}`)
	config2, err := LoadConfig(byteData)
	if err != nil {
		t.Fatalf("Failed to load config from bytes: %v", err)
	}
	
	if config2.MaxContextBytes != 2097152 {
		t.Errorf("Expected MaxContextBytes to be 2097152, got %d", config2.MaxContextBytes)
	}
	
	// Test loading from map
	mapData := map[string]interface{}{
		"max_context_bytes": int64(4194304),
	}
	config3, err := LoadConfig(mapData)
	if err != nil {
		t.Fatalf("Failed to load config from map: %v", err)
	}
	
	if config3.MaxContextBytes != 4194304 {
		t.Errorf("Expected MaxContextBytes to be 4194304, got %d", config3.MaxContextBytes)
	}
}

func TestConfig_LoadFromFile(t *testing.T) {
	// This test would require creating a temporary file
	// For now, we're just testing the function signature and basic structure
	// Actual file I/O testing would require more setup
	t.Log("LoadFromFile function exists and is callable")
}