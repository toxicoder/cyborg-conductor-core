package runner

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestSubprocessRunner tests the subprocess runner implementation
func TestSubprocessRunner(t *testing.T) {
	// Create subprocess runner
	runner := NewSubprocessRunner()
	
	// Test that runner can be created
	assert.NotNil(t, runner)
}

// TestSubprocessRunnerRun tests running a simple command
func TestSubprocessRunnerRun(t *testing.T) {
	// Create subprocess runner
	runner := NewSubprocessRunner()
	
	// Test with a simple command that should succeed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result, err := runner.Run(ctx, "echo", []string{"hello", "world"})
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Stdout)
	assert.Empty(t, result.Stderr)
	assert.NoError(t, result.Err)
}

// TestSubprocessRunnerTimeout tests timeout handling
func TestSubprocessRunnerTimeout(t *testing.T) {
	// Create subprocess runner
	runner := NewSubprocessRunner()
	
	// Test with a command that will timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// Use a command that sleeps longer than the timeout
	result, err := runner.Run(ctx, "sleep", []string{"5"})
	
	// Should timeout
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout")
	assert.Nil(t, result)
}

// TestSubprocessRunnerError tests error handling
func TestSubprocessRunnerError(t *testing.T) {
	// Create subprocess runner
	runner := NewSubprocessRunner()
	
	// Test with a command that should fail
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result, err := runner.Run(ctx, "nonexistent-command-12345", []string{})
	
	// Should return an error
	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestSubprocessRunnerWithPolicy tests running with policy
func TestSubprocessRunnerWithPolicy(t *testing.T) {
	// Create subprocess runner
	runner := NewSubprocessRunner()
	
	// Test with a simple command
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result, err := runner.RunWithPolicy(ctx, "echo", []string{"test"}, nil)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Stdout)
}