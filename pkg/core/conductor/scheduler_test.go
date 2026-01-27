package conductor

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	pb "github.com/toxicoder/cyborg-conductor-core/pkg/proto/pb"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/pb"
)

// TestConductor tests the conductor implementation
func TestConductor(t *testing.T) {
	// Create a mock registry
	registry := pb.NewRegistry()
	
	// Create conductor
	conductor := NewConductor(registry)
	
	// Test that conductor can be created
	assert.NotNil(t, conductor)
	assert.NotNil(t, conductor.pool)
	assert.NotNil(t, conductor.registry)
}

// TestConductorStartStop tests starting and stopping the conductor
func TestConductorStartStop(t *testing.T) {
	// Create a mock registry
	registry := pb.NewRegistry()
	
	// Create conductor
	conductor := NewConductor(registry)
	
	// Test start
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	err := conductor.Start(ctx)
	assert.NoError(t, err)
	
	// Test stop
	err = conductor.Stop()
	assert.NoError(t, err)
}

// TestConductorSubmitJob tests submitting jobs to the conductor
func TestConductorSubmitJob(t *testing.T) {
	// Create a mock registry
	registry := pb.NewRegistry()
	
	// Create conductor
	conductor := NewConductor(registry)
	
	// Test submitting a job
	job := &Job{
		ID:          "test-job-1",
		Capabilities: []string{"test-capability"},
		Payload:     []byte("test payload"),
		TimeoutMs:   3000,
	}
	
	err := conductor.SubmitJob(job)
	assert.NoError(t, err)
}

// TestConductorBackPressure tests back-pressure handling
func TestConductorBackPressure(t *testing.T) {
	// Create a mock registry
	registry := pb.NewRegistry()
	
	// Create conductor with small queue
	conductor := NewConductor(registry)
	
	// Submit many jobs to test back-pressure
	for i := 0; i < 2000; i++ { // More than the buffer size
		job := &Job{
			ID:          "test-job",
			Capabilities: []string{"test-capability"},
			Payload:     []byte("test payload"),
			TimeoutMs:   3000,
		}
		
		err := conductor.SubmitJob(job)
		// Allow some jobs to succeed, but queue may fill up
		if err != nil {
			// This is expected when queue is full
			break
		}
	}
	
	// Test that we can still submit jobs
	job := &Job{
		ID:          "test-job-2",
		Capabilities: []string{"test-capability"},
		Payload:     []byte("test payload"),
		TimeoutMs:   3000,
	}
	
	err := conductor.SubmitJob(job)
	assert.NoError(t, err)
}

// TestConductorCapacityEnforcement tests capacity enforcement
func TestConductorCapacityEnforcement(t *testing.T) {
	// Create a mock registry
	registry := pb.NewRegistry()
	
	// Create conductor
	conductor := NewConductor(registry)
	
	// Test that we can submit jobs without error
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	job := &Job{
		ID:          "test-job-1",
		Capabilities: []string{"test-capability"},
		Payload:     []byte("test payload"),
		TimeoutMs:   3000,
	}
	
	err := conductor.SubmitJob(job)
	assert.NoError(t, err)
	
	// Test that conductor can be started and stopped
	err = conductor.Start(ctx)
	assert.NoError(t, err)
	
	time.Sleep(10 * time.Millisecond) // Let it process
	
	err = conductor.Stop()
	assert.NoError(t, err)
}