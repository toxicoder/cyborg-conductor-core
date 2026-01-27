package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config represents the application configuration
type Config struct {
	// Server configuration
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		GrpcPort int `json:"grpc_port"`
	} `json:"server"`
	
	// TLS configuration
	TLS struct {
		CertFile string `json:"cert_file"`
		KeyFile  string `json:"key_file"`
	} `json:"tls"`
	
	// Database configuration
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	
	// Logging configuration
	Logging struct {
		Level  string `json:"level"`
		Format string `json:"format"`
	} `json:"logging"`
	
	// Cyborg configuration
	Cyborg struct {
		// Path to the jobs matrix CSV file
		JobsMatrixPath string `json:"jobs_matrix_path"`
		// Default namespace for cyborg deployments
		DefaultNamespace string `json:"default_namespace"`
	} `json:"cyborg"`
	
	// Runtime configuration
	Runtime struct {
		// Maximum concurrent streams
		MaxConcurrentStreams int `json:"max_concurrent_streams"`
		// Timeout settings
		Timeout struct {
			JobTimeout     int64 `json:"job_timeout"`
			TaskTimeout    int64 `json:"task_timeout"`
			Connection     int64 `json:"connection"`
		} `json:"timeout"`
	} `json:"runtime"`
}

// LoadConfig loads configuration from environment variables and defaults
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: struct {
			Host string `json:"host"`
			Port int    `json:"port"`
			GrpcPort int `json:"grpc_port"`
		}{
			Host: "localhost",
			Port: 8080,
			GrpcPort: 9090,
		},
		Database: struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Host:     "localhost",
			Port:     5432,
			Name:     "cyborg_conductor",
			Username: "postgres",
			Password: "postgres",
		},
		Logging: struct {
			Level  string `json:"level"`
			Format string `json:"format"`
		}{
			Level:  "info",
			Format: "json",
		},
		Cyborg: struct {
			JobsMatrixPath     string `json:"jobs_matrix_path"`
			DefaultNamespace   string `json:"default_namespace"`
		}{
			JobsMatrixPath:   "cyborgs/jobs_matrix.csv",
			DefaultNamespace: "default",
		},
		Runtime: struct {
			MaxConcurrentStreams int `json:"max_concurrent_streams"`
			Timeout              struct {
				JobTimeout     int64 `json:"job_timeout"`
				TaskTimeout    int64 `json:"task_timeout"`
				Connection     int64 `json:"connection"`
			} `json:"timeout"`
		}{
			MaxConcurrentStreams: 100,
			Timeout: struct {
				JobTimeout     int64 `json:"job_timeout"`
				TaskTimeout    int64 `json:"task_timeout"`
				Connection     int64 `json:"connection"`
			}{
				JobTimeout:  300, // 5 minutes
				TaskTimeout: 60,  // 1 minute
				Connection:  30,  // 30 seconds
			},
		},
	}
	
	// Load from environment variables
	loadFromEnv(cfg)
	
	// Validate required configuration
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}
	
	return cfg, nil
}

// validateConfig validates the required configuration parameters
func validateConfig(cfg *Config) error {
	// Validate database configuration
	if cfg.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if cfg.Database.Port == 0 {
		return fmt.Errorf("database port is required")
	}
	if cfg.Database.Name == "" {
		return fmt.Errorf("database name is required")
	}
	if cfg.Database.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if cfg.Database.Password == "" {
		return fmt.Errorf("database password is required")
	}
	
	// Validate server configuration
	if cfg.Server.Host == "" {
		return fmt.Errorf("server host is required")
	}
	if cfg.Server.Port == 0 {
		return fmt.Errorf("server port is required")
	}
	
	return nil
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(cfg *Config) {
	// Server config
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Host = host
	}
	if portStr := os.Getenv("SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Server.Port = port
		}
	}
	if grpcPortStr := os.Getenv("GRPC_PORT"); grpcPortStr != "" {
		if port, err := strconv.Atoi(grpcPortStr); err == nil {
			cfg.Server.GrpcPort = port
		}
	}
	
	// Database config
	if host := os.Getenv("DB_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if portStr := os.Getenv("DB_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Database.Port = port
		}
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		cfg.Database.Name = name
	}
	if user := os.Getenv("DB_USER"); user != "" {
		cfg.Database.Username = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		cfg.Database.Password = pass
	}
	
	// Logging config
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		cfg.Logging.Level = level
	}
	if format := os.Getenv("LOG_FORMAT"); format != "" {
		cfg.Logging.Format = format
	}
	
	// Cyborg config
	if jobsPath := os.Getenv("JOBS_MATRIX_PATH"); jobsPath != "" {
		cfg.Cyborg.JobsMatrixPath = jobsPath
	}
	if namespace := os.Getenv("DEFAULT_NAMESPACE"); namespace != "" {
		cfg.Cyborg.DefaultNamespace = namespace
	}
	
	// Runtime config
	if maxStreamsStr := os.Getenv("MAX_CONCURRENT_STREAMS"); maxStreamsStr != "" {
		if maxStreams, err := strconv.Atoi(maxStreamsStr); err == nil {
			cfg.Runtime.MaxConcurrentStreams = maxStreams
		}
	}
	
	// Timeout config
	if jobTimeoutStr := os.Getenv("JOB_TIMEOUT"); jobTimeoutStr != "" {
		if jobTimeout, err := strconv.ParseInt(jobTimeoutStr, 10, 64); err == nil {
			cfg.Runtime.Timeout.JobTimeout = jobTimeout
		}
	}
	if taskTimeoutStr := os.Getenv("TASK_TIMEOUT"); taskTimeoutStr != "" {
		if taskTimeout, err := strconv.ParseInt(taskTimeoutStr, 10, 64); err == nil {
			cfg.Runtime.Timeout.TaskTimeout = taskTimeout
		}
	}
	if connTimeoutStr := os.Getenv("CONNECTION_TIMEOUT"); connTimeoutStr != "" {
		if connTimeout, err := strconv.ParseInt(connTimeoutStr, 10, 64); err == nil {
			cfg.Runtime.Timeout.Connection = connTimeout
		}
	}
}

// GetEnv gets an environment variable value with a default fallback
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvInt gets an environment variable integer value with a default fallback
func GetEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

// GetEnvInt64 gets an environment variable int64 value with a default fallback
func GetEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return i
		}
	}
	return defaultValue
}

// GetEnvBool gets an environment variable boolean value with a default fallback
func GetEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return strings.ToLower(value) == "true"
	}
	return defaultValue
}