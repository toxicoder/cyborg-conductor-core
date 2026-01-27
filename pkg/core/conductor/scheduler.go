package conductor

import (
	"context"
	"sync"
	"time"

	pb "github.com/toxicoder/cyborg-conductor-core/pkg/proto/generated"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/pb"
)

// Job represents a unit of work to be executed by a cyborg
type Job struct {
	ID          string
	Capabilities []string
	Payload     []byte
	TimeoutMs   int32
	// Add a reference to the cyborg that should execute this job
	CyborgID    string
}

// Pool represents a back-pressure aware worker pool
type Pool struct {
	jobChan chan *Job
	wg      sync.WaitGroup
}

// Conductor manages job distribution to cyborgs
type Conductor struct {
	pool        *Pool
	registry    *pb.Registry
	mu          sync.RWMutex
	running     bool
	workerCount int
}

// NewConductor creates a new conductor with the given registry
func NewConductor(registry *pb.Registry) *Conductor {
	return &Conductor{
		registry: registry,
		pool: &Pool{
			jobChan: make(chan *Job, 1000), // Buffered channel for back-pressure
		},
		workerCount: 10, // Default worker count
	}
}

// Start initializes and starts the conductor workers
func (c *Conductor) Start(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return nil
	}
	c.running = true
	c.mu.Unlock()

	// Start worker goroutines
	for i := 0; i < c.workerCount; i++ {
		c.pool.wg.Add(1)
		go c.worker(ctx, i)
	}

	// Start job dispatcher
	go c.dispatchJobs(ctx)

	return nil
}

// Stop shuts down the conductor
func (c *Conductor) Stop() error {
	c.mu.Lock()
	if !c.running {
		c.mu.Unlock()
		return nil
	}
	c.running = false
	c.mu.Unlock()

	close(c.pool.jobChan)
	c.pool.wg.Wait()
	return nil
}

// SubmitJob submits a job to the conductor queue
func (c *Conductor) SubmitJob(job *Job) error {
	select {
	case c.pool.jobChan <- job:
		return nil
	default:
		// Queue is full - implement back-pressure
		return &JobQueueFullError{Message: "Job queue is full"}
	}
}

// worker processes jobs from the job channel
func (c *Conductor) worker(ctx context.Context, workerID int) {
	defer c.pool.wg.Done()
	
	for {
		select {
		case job, ok := <-c.pool.jobChan:
			if !ok {
				return // Channel closed
			}
			
			// Process the job
			err := c.processJob(ctx, job)
			if err != nil {
				// Handle job processing error
				// In a real implementation, this would log and potentially retry
				continue
			}
			
		case <-ctx.Done():
			return
		}
	}
}

// dispatchJobs continuously dispatches jobs to available cyborgs
func (c *Conductor) dispatchJobs(ctx context.Context) {
	for {
		select {
		case job, ok := <-c.pool.jobChan:
			if !ok {
				return
			}
			
			// Find suitable cyborg for the job
			cyborg, err := c.findSuitableCyborg(job)
			if err != nil {
				// Handle error - could log and potentially retry
				continue
			}
			
			if cyborg != nil {
				// Dispatch to cyborg
				err := c.dispatchToCyborg(ctx, job, cyborg)
				if err != nil {
					// Handle dispatch error
					continue
				}
			} else {
				// No suitable cyborg found - could queue or fail
				// For now, we'll just log and continue
			}
			
		case <-ctx.Done():
			return
		}
	}
}

// findSuitableCyborg selects the best cyborg for a job based on capabilities
func (c *Conductor) findSuitableCyborg(job *Job) (*pb.CyborgDescriptor, error) {
	// Get all registered cyborgs
	cyborgs := c.registry.List()
	
	// Simple selection logic - in a real implementation this would be more sophisticated
	// and consider factors like load, latency budget, etc.
	// For now, we'll just return the first cyborg with all required capabilities
	for _, cyborg := range cyborgs {
		// Check if cyborg has all required capabilities
		if c.hasAllCapabilities(cyborg, job.Capabilities) {
			return cyborg, nil
		}
	}
	
	return nil, nil // No suitable cyborg found
}

// hasAllCapabilities checks if a cyborg has all the required capabilities
func (c *Conductor) hasAllCapabilities(cyborg *pb.CyborgDescriptor, requiredCaps []string) bool {
	if cyborg == nil || cyborg.Capabilities == nil {
		return false
	}
	
	// Convert cyborg capabilities to map for faster lookup
	cyborgCaps := make(map[string]bool)
	for _, cap := range cyborg.Capabilities {
		cyborgCaps[cap.Name] = true // Assuming CapabilitySpec has a Name field
	}
	
	// Check if all required capabilities are present
	for _, requiredCap := range requiredCaps {
		if !cyborgCaps[requiredCap] {
			return false
		}
	}
	
	return true
}

// processJob handles the processing of a single job
func (c *Conductor) processJob(ctx context.Context, job *Job) error {
	// In a real implementation, this would:
	// 1. Validate the job
	// 2. Determine if it's deterministic or LLM-based
	// 3. Call appropriate handler (SubprocessRunner or LLM streaming session)
	
	// For now, just simulate processing
	time.Sleep(10 * time.Millisecond)
	return nil
}

// dispatchToCyborg dispatches a job to a specific cyborg
func (c *Conductor) dispatchToCyborg(ctx context.Context, job *Job, cyborg *pb.CyborgDescriptor) error {
	// In a real implementation, this would:
	// 1. Create a SubmitJobMessage with the job details
	// 2. Send it to the cyborg via appropriate transport
	// 3. Handle the response
	
	// For now, just simulate dispatch
	time.Sleep(5 * time.Millisecond)
	return nil
}

// JobQueueFullError represents an error when job queue is full
type JobQueueFullError struct {
	Message string
}

func (e *JobQueueFullError) Error() string {
	return e.Message
}