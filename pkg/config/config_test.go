package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Save original environment variables
	originalServerHost := os.Getenv("SERVER_HOST")
	originalServerPort := os.Getenv("SERVER_PORT")
	originalDBHost := os.Getenv("DB_HOST")
	originalDBPort := os.Getenv("DB_PORT")
	originalLogLevel := os.Getenv("LOG_LEVEL")
	originalJobsMatrixPath := os.Getenv("JOBS_MATRIX_PATH")

	// Clean up environment variables
	os.Unsetenv("SERVER_HOST")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("JOBS_MATRIX_PATH")

	defer func() {
		// Restore original environment variables
		os.Setenv("SERVER_HOST", originalServerHost)
		os.Setenv("SERVER_PORT", originalServerPort)
		os.Setenv("DB_HOST", originalDBHost)
		os.Setenv("DB_PORT", originalDBPort)
		os.Setenv("LOG_LEVEL", originalLogLevel)
		os.Setenv("JOBS_MATRIX_PATH", originalJobsMatrixPath)
	}()

	// Test default configuration
	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Test default values
	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Equal(t, "cyborg_conductor", cfg.Database.Name)
	assert.Equal(t, "postgres", cfg.Database.Username)
	assert.Equal(t, "postgres", cfg.Database.Password)
	assert.Equal(t, "info", cfg.Logging.Level)
	assert.Equal(t, "json", cfg.Logging.Format)
	assert.Equal(t, "cyborgs/jobs_matrix.csv", cfg.Cyborg.JobsMatrixPath)
	assert.Equal(t, "default", cfg.Cyborg.DefaultNamespace)
	assert.Equal(t, 100, cfg.Runtime.MaxConcurrentStreams)
	assert.Equal(t, int64(300), cfg.Runtime.Timeout.JobTimeout)
	assert.Equal(t, int64(60), cfg.Runtime.Timeout.TaskTimeout)
	assert.Equal(t, int64(30), cfg.Runtime.Timeout.Connection)

	// Test environment variable overrides
	os.Setenv("SERVER_HOST", "test-host")
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("DB_HOST", "test-db-host")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("JOBS_MATRIX_PATH", "/custom/path")

	cfg2, err := LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg2)

	assert.Equal(t, "test-host", cfg2.Server.Host)
	assert.Equal(t, 9090, cfg2.Server.Port)
	assert.Equal(t, "test-db-host", cfg2.Database.Host)
	assert.Equal(t, 5433, cfg2.Database.Port)
	assert.Equal(t, "debug", cfg2.Logging.Level)
	assert.Equal(t, "/custom/path", cfg2.Cyborg.JobsMatrixPath)
}

func TestGetEnv(t *testing.T) {
	// Test with default value
	result := GetEnv("NONEXISTENT_VAR", "default")
	assert.Equal(t, "default", result)

	// Test with actual environment variable
	os.Setenv("TEST_VAR", "test-value")
	result = GetEnv("TEST_VAR", "default")
	assert.Equal(t, "test-value", result)
}

func TestGetEnvInt(t *testing.T) {
	// Test with default value
	result := GetEnvInt("NONEXISTENT_INT_VAR", 42)
	assert.Equal(t, 42, result)

	// Test with actual environment variable
	os.Setenv("TEST_INT_VAR", "123")
	result = GetEnvInt("TEST_INT_VAR", 42)
	assert.Equal(t, 123, result)

	// Test with invalid environment variable (should use default)
	os.Setenv("TEST_INT_VAR", "invalid")
	result = GetEnvInt("TEST_INT_VAR", 42)
	assert.Equal(t, 42, result)
}

func TestGetEnvInt64(t *testing.T) {
	// Test with default value
	result := GetEnvInt64("NONEXISTENT_INT64_VAR", 42)
	assert.Equal(t, int64(42), result)

	// Test with actual environment variable
	os.Setenv("TEST_INT64_VAR", "123456")
	result = GetEnvInt64("TEST_INT64_VAR", 42)
	assert.Equal(t, int64(123456), result)

	// Test with invalid environment variable (should use default)
	os.Setenv("TEST_INT64_VAR", "invalid")
	result = GetEnvInt64("TEST_INT64_VAR", 42)
	assert.Equal(t, int64(42), result)
}

func TestGetEnvBool(t *testing.T) {
	// Test with default value
	result := GetEnvBool("NONEXISTENT_BOOL_VAR", true)
	assert.Equal(t, true, result)

	// Test with actual environment variable - true case
	os.Setenv("TEST_BOOL_VAR", "true")
	result = GetEnvBool("TEST_BOOL_VAR", false)
	assert.Equal(t, true, result)

	// Test with actual environment variable - false case  
	os.Setenv("TEST_BOOL_VAR", "false")
	result = GetEnvBool("TEST_BOOL_VAR", true)
	assert.Equal(t, false, result)

	// Test with invalid environment variable (should use default)
	os.Unsetenv("TEST_BOOL_VAR") // Clear the variable
	result = GetEnvBool("TEST_BOOL_VAR", true)
	assert.Equal(t, true, result)
}