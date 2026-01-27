package runner

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os/exec"
	"sync"
	"time"

	"github.com/toxicoder/cyborg-conductor-core/pkg/proto/generated"
)

// SubprocessRunner executes external scripts with timeout enforcement
type SubprocessRunner struct {
	// Add any necessary configuration fields
	config *config.Config
}

// RunResult contains the result of a subprocess execution
type RunResult struct {
	Stdout []byte
	Stderr []byte
	Err    error
}

// NewSubprocessRunner creates a new subprocess runner
func NewSubprocessRunner(cfg *config.Config) *SubprocessRunner {
	return &SubprocessRunner{
		config: cfg,
	}
}

// Run executes a script with the given arguments
// It uses exec.CommandContext with a global timeout
func (r *SubprocessRunner) Run(ctx context.Context, script string, args []string) (*RunResult, error) {
	// Create context with timeout using the configured task timeout
	timeout := time.Duration(r.config.Runtime.Timeout.TaskTimeout) * time.Second
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Create command
	cmd := exec.CommandContext(timeoutCtx, script, args...)

	// Capture output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	// Read output asynchronously to avoid deadlocks
	stdoutBytes, stderrBytes := make([]byte, 0), make([]byte, 0)
	done := make(chan struct{})

	go func() {
		defer close(done)
		stdoutBytes, _ = io.ReadAll(stdout)
	}()

	go func() {
		defer close(done)
		stderrBytes, _ = io.ReadAll(stderr)
	}()

	// Wait for command to complete or timeout
	err = cmd.Wait()

	// Wait for all reads to complete
	<-done
	<-done

	// Check if command timed out
	if timeoutCtx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command timed out after %d seconds", r.config.Runtime.Timeout.TaskTimeout)
	}

	// Handle command execution error
	if err != nil {
		// Decode stdout/stderr into protobuf-wrapped messages if needed
		// For now, we'll just return the raw output and error
		return &RunResult{
			Stdout: stdoutBytes,
			Stderr: stderrBytes,
			Err:    err,
		}, nil
	}

	// Success case
	return &RunResult{
		Stdout: stdoutBytes,
		Stderr: stderrBytes,
		Err:    nil,
	}, nil
}

// RunWithPolicy executes a script with the given policy
func (r *SubprocessRunner) RunWithPolicy(ctx context.Context, script string, args []string, policy interface{}) (*RunResult, error) {
	// Apply policy if needed
	if policy != nil {
		// Enforce policy constraints
		// For example, check resource limits, etc.
	}

	// Call the main Run method
	return r.Run(ctx, script, args)
}