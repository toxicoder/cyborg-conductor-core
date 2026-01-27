package orchestrator

import (
	"context"
	"sync"
	"time"

	"github.com/toxicoder/cyborg-conductor-core/pkg/core/pb"
	"github.com/toxicoder/cyborg-conductor-core/pkg/memory/manager"
	"github.com/toxicoder/cyborg-conductor-core/internal/runner"
)

// Scheduler holds a back-pressure aware worker pool and selects a cyborg based on capability match, latency budget, and current load
type Scheduler struct {
	// Mutex to protect access to internal state
	mu sync.RWMutex
	
	// Cyborg registry - maps cyborg IDs to their descriptors
	registry map[string]*pb.CyborgDescriptor
	
	// Memory manager for context size tracking
	memoryManager *manager.MemoryCacheManager
	
	// Execution manager for running tasks
	execManager *runner.SubprocessRunner
	
	// Worker pool channels
	workerPool chan chan *Task
	taskQueue  chan *Task
	
	// Shutdown signal
	shutdown chan struct{}
	
	// Current worker count
	workerCount int32
}

// Task represents a unit of work to be scheduled
type Task struct {
	// The cyborg descriptor to run the task on
	Cyborg *pb.CyborgDescriptor
	
	// The command/script to execute
	Command string
	
	// Arguments for the command
	Args []string
	
	// Context for cancellation/timeout
	Context context.Context
	
	// Result channel to send results back
	Result chan *TaskResult
	
	// Timeout duration for the task
	Timeout time.Duration
}

// TaskResult represents the result of a scheduled task
type TaskResult struct {
	// Output from stdout
	Stdout []byte
	
	// Output from stderr
	Stderr []byte
	
	// Error if execution failed
	Err error
	
	// Duration of execution
	Duration time.Duration
	
	// Cyborg used for execution
	CyborgID string
}

// NewScheduler creates a new scheduler with the given memory manager and execution manager
func NewScheduler(memoryManager *manager.MemoryCacheManager, execManager *runner.SubprocessRunner) *Scheduler {
	s := &Scheduler{
		registry:      make(map[string]*pb.CyborgDescriptor),
		memoryManager: memoryManager,
		execManager:   execManager,
		workerPool:    make(chan chan *Task, 10), // Buffered channel for worker pool
		taskQueue:     make(chan *Task, 100),     // Buffered channel for tasks
		shutdown:      make(chan struct{}),
		workerCount:   0,
	}
	
	// Start the scheduler loop
	go s.run()
	
	return s
}

// RegisterCyborg adds a cyborg to the registry
func (s *Scheduler) RegisterCyborg(descriptor *pb.CyborgDescriptor) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.registry[descriptor.Id] = descriptor
	return nil
}

// GetCyborg retrieves a cyborg descriptor by ID
func (s *Scheduler) GetCyborg(id string) (*pb.CyborgDescriptor, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	descriptor, exists := s.registry[id]
	return descriptor, exists
}

// ListCyborgs returns all registered cyborgs
func (s *Scheduler) ListCyborgs() []*pb.CyborgDescriptor {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	cyborgs := make([]*pb.CyborgDescriptor, 0, len(s.registry))
	for _, descriptor := range s.registry {
		cyborgs = append(cyborgs, descriptor)
	}
	
	return cyborgs
}

// SelectCyborg selects an appropriate cyborg based on capability match, latency budget, and current load
func (s *Scheduler) SelectCyborg(capabilities []string, latencyBudget time.Duration) *pb.CyborgDescriptor {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	// In a real implementation, this would use more sophisticated logic:
	// - Match capabilities
	// - Consider latency budgets
	// - Check current load on cyborgs
	// - Consider resource quotas
	// For now, we'll select the first available cyborg
	
	for _, descriptor := range s.registry {
		// Simple capability matching logic
		if s.hasAllCapabilities(descriptor, capabilities) {
			// In a real implementation, we'd also check:
			// - Current load (number of concurrent tasks)
			// - Available resources
			// - Latency constraints
			return descriptor
		}
	}
	
	return nil
}

// hasAllCapabilities checks if a cyborg has all required capabilities
func (s *Scheduler) hasAllCapabilities(descriptor *pb.CyborgDescriptor, required []string) bool {
	// Convert the descriptor's capabilities to a set for efficient lookup
	descriptorCaps := make(map[string]bool)
	for _, cap := range descriptor.Capabilities {
		descriptorCaps[cap] = true
	}
	
	// Check if all required capabilities are present
	for _, req := range required {
		if !descriptorCaps[req] {
			return false
		}
	}
	
	return true
}

// SubmitTask submits a task to the scheduler for execution
func (s *Scheduler) SubmitTask(ctx context.Context, command string, args []string, capabilities []string, timeout time.Duration) (*TaskResult, error) {
	// Select an appropriate cyborg
	cyborg := s.SelectCyborg(capabilities, timeout)
	if cyborg == nil {
		return nil, &NoCyborgAvailableError{Capabilities: capabilities}
	}
	
	// Create task
	task := &Task{
		Cyborg:  cyborg,
		Command: command,
		Args:    args,
		Context: ctx,
		Result:  make(chan *TaskResult, 1),
		Timeout: timeout,
	}
	
	// Submit to task queue
	select {
	case s.taskQueue <- task:
		// Task submitted successfully
	default:
		// Task queue is full, return an error
		return nil, &QueueFullError{}
	}
	
	// Wait for result
	select {
	case result := <-task.Result:
		return result, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// run starts the scheduler loop
func (s *Scheduler) run() {
	for {
		select {
		case <-s.shutdown:
			return
		default:
			// Select a task from the queue
			select {
			case task := <-s.taskQueue:
				// Submit to a worker
				s.submitToWorker(task)
			default:
				// No tasks available, sleep briefly to avoid busy waiting
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
}

// submitToWorker submits a task to an available worker
func (s *Scheduler) submitToWorker(task *Task) {
	select {
	case workerChan := <-s.workerPool:
		// Send task to worker
		workerChan <- task
	default:
		// No available worker, start a new one
		go s.startWorker()
		// Try again
		go func() {
			select {
			case workerChan := <-s.workerPool:
				workerChan <- task
			default:
				// If still no worker available, put task back in queue
				s.taskQueue <- task
			}
		}()
	}
}

// startWorker creates and starts a new worker goroutine
func (s *Scheduler) startWorker() {
	s.mu.Lock()
	s.workerCount++
	s.mu.Unlock()
	
	workerChan := make(chan *Task)
	
	go func() {
		defer func() {
			s.mu.Lock()
			s.workerCount--
			s.mu.Unlock()
		}()
		
		for {
			select {
			case task := <-workerChan:
				// Execute the task
				result := s.executeTask(task)
				
				// Send result back
				select {
				case task.Result <- result:
				default:
					// Result channel is full, drop the result
				}
				
				// Return worker to pool
				select {
				case s.workerPool <- workerChan:
				default:
					// Pool is full, worker will exit
				}
			case <-s.shutdown:
				// Shutdown signal received
				return
			}
		}
	}()
	
	// Add worker to pool
	select {
	case s.workerPool <- workerChan:
	default:
		// Pool is full, worker will exit
	}
}

// executeTask executes a task on the assigned cyborg
func (s *Scheduler) executeTask(task *Task) *TaskResult {
	start := time.Now()
	
	// Run the command using the execution manager
	result, err := s.execManager.Run(task.Context, task.Command, task.Args)
	
	duration := time.Since(start)
	
	return &TaskResult{
		Stdout:   result.Stdout,
		Stderr:   result.Stderr,
		Err:      err,
		Duration: duration,
		CyborgID: task.Cyborg.Id,
	}
}

// Shutdown gracefully shuts down the scheduler
func (s *Scheduler) Shutdown() {
	close(s.shutdown)
}

// NoCyborgAvailableError is returned when no cyborg is available to handle a task
type NoCyborgAvailableError struct {
	Capabilities []string
}

func (e *NoCyborgAvailableError) Error() string {
	return "no cyborg available with required capabilities"
}

// QueueFullError is returned when the task queue is full
type QueueFullError struct{}

func (e *QueueFullError) Error() string {
	return "task queue is full"
}