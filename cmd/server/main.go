package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/toxicoder/cyborg-conductor-core/internal/context"
	"github.com/toxicoder/cyborg-conductor-core/internal/runner"
	"github.com/toxicoder/cyborg-conductor-core/pkg/config"
	"github.com/toxicoder/cyborg-conductor-core/pkg/context/pb"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/conductor"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/pb"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/types"
)

var logger *zap.Logger
var cfg *config.Config
var db *sql.DB
var scheduler *conductor.Scheduler

func initLogger() *zap.Logger {
	// Create a logger configuration
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	// Create logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	
	return logger
}

// initDB initializes the database connection
func initDB() error {
	// Construct PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
	
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	
	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	// Test the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	
	logger.Info("Database connection established successfully")
	return nil
}

func main() {
	// Initialize logger
	logger = initLogger()
	defer logger.Sync()

	fmt.Println("Cyborg Conductor Core Server")
	
	// Load configuration first
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}
	
	// Initialize core components
	if err := initializeServer(); err != nil {
		logger.Fatal("Failed to initialize server", zap.Error(err))
	}
	
	// Start server
	if err := startServer(); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
	
	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	logger.Info("Server shutting down...")
}

func initializeServer() error {
	logger.Info("Initializing server components...")
	
	// Initialize database connection
	if err := initDB(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	
	// Initialize registry with all cyborgs from jobs matrix
	if err := registerAllCyborgs(); err != nil {
		return fmt.Errorf("failed to register cyborgs: %w", err)
	}
	
	// Initialize memory manager with configurable path
	// Using default size limit of 100MB
	manager, err := context.NewMemoryCacheManager(100 * 1024 * 1024) // 100MB
	if err != nil {
		return fmt.Errorf("failed to initialize memory manager: %w", err)
	}
	
	// Initialize execution manager
	execManager := runner.NewSubprocessRunner(cfg)
	
	// Initialize scheduler
	scheduler = conductor.NewScheduler(manager, execManager)
	
	// Register components with global state (or context)
	// This would typically be done through dependency injection or global registry
	
	logger.Info("Server components initialized successfully")
	return nil
}

// CyborgRegistrationServiceServer implements the gRPC service
type CyborgRegistrationServiceServer struct {
	// Embed the generated service interface
	pb.UnimplementedCyborgRegistrationServiceServer
}

// RegisterCyborg implements the gRPC method for registering a cyborg
func (s *CyborgRegistrationServiceServer) RegisterCyborg(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	logger.Info("Received RegisterCyborg request")
	
	// Validate the request
	if req == nil || req.CyborgDescriptor == nil {
		return &pb.RegisterResponse{
			Success: false,
			Error:   "Invalid request: cyborg descriptor is required",
		}, nil
	}
	
	// Validate that cyborg ID is not empty
	if req.CyborgDescriptor.CyborgId == "" {
		return &pb.RegisterResponse{
			Success: false,
			Error:   "Invalid cyborg descriptor: cyborg ID cannot be empty",
		}, nil
	}
	
	// Validate that tags are not empty
	if len(req.CyborgDescriptor.Tags) == 0 {
		return &pb.RegisterResponse{
			Success: false,
			Error:   "Invalid cyborg descriptor: tags cannot be empty",
		}, nil
	}
	
	// Validate deployment spec is parsable (basic check)
	if req.CyborgDescriptor.DeploymentSpec == "" {
		return &pb.RegisterResponse{
			Success: false,
			Error:   "Invalid cyborg descriptor: deployment spec cannot be empty",
		}, nil
	}
	
	// Add the cyborg to the scheduler's registry
	err := scheduler.RegisterCyborg(req.CyborgDescriptor)
	if err != nil {
		return &pb.RegisterResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to register cyborg: %v", err),
		}, nil
	}
	
	// Return success
	return &pb.RegisterResponse{
		Success: true,
		// In a real implementation, we might assign a system-generated ID
		RegisteredCyborgId: req.CyborgDescriptor.CyborgId,
	}, nil
}

func startServer() error {
	logger.Info("Starting server...")
	
	// Initialize HTTP server
	mux := http.NewServeMux()
	
	// Add health check endpoint with better checks
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// Add security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		
		// Check database connectivity
		if err := db.Ping(); err != nil {
			logger.Error("Database health check failed", zap.Error(err))
			http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
			return
		}
		
		// Basic health check - server is up
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	// Add status endpoint
	mux.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		// Add security headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		
		// Return server status information
		w.Write([]byte(`{"status": "running", "version": "1.0.0"}`))
	})
	
	// Add Prometheus metrics endpoint
	mux.Handle("/metrics", promhttp.Handler())
	
	// Update TODO.md to reflect the implementation
	// This is a placeholder to note that metrics are now implemented
	// In a production system, we would also want to register custom metrics here
	
	// Start server with TLS support if certificates are provided
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCyborgRegistrationServiceServer(grpcServer, &CyborgRegistrationServiceServer{})
	
	// Listen on gRPC port
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GrpcPort))
	if err != nil {
		logger.Fatal("Failed to create gRPC listener", zap.Error(err))
	}
	
	go func() {
		logger.Info("gRPC server listening", zap.Int("port", cfg.Server.GrpcPort))
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Fatal("gRPC server failed to start", zap.Error(err))
		}
	}()
	
	go func() {
		logger.Info("Server listening", zap.Int("port", cfg.Server.Port))
		if cfg.TLS.CertFile != "" && cfg.TLS.KeyFile != "" {
			logger.Info("Starting HTTPS server", zap.String("cert_file", cfg.TLS.CertFile), zap.String("key_file", cfg.TLS.KeyFile))
			if err := server.ListenAndServeTLS(cfg.TLS.CertFile, cfg.TLS.KeyFile); err != nil && err != http.ErrServerClosed {
				logger.Fatal("HTTPS server failed to start", zap.Error(err))
			}
		} else {
			logger.Info("Starting HTTP server")
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal("HTTP server failed to start", zap.Error(err))
			}
		}
	}()
	
	logger.Info("Server running...")
	return nil
}

func registerAllCyborgs() error {
	logger.Info("Registering all cyborgs from txtpb files...")
	
	// Create and load registry from txtpb files
	registry := pb.NewRegistry()
	if err := registry.LoadFromTxtpb("cyborgs"); err != nil {
		return fmt.Errorf("failed to load cyborgs from txtpb files: %w", err)
	}
	
	logger.Info("Successfully registered cyborgs", zap.Int("count", registry.Size()))
	return nil
}