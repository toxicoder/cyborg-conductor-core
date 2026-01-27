package adapt

import (
	"context"
	"fmt"
	"time"
)

// Runner represents a cyborg runner that adapts to different tasks
type Runner struct {
	ID          string
	Name        string
	Health      string
	LastSeen    time.Time
	ActiveTasks int
}

// NewRunner creates a new runner instance
func NewRunner(id, name string) *Runner {
	return &Runner{
		ID:       id,
		Name:     name,
		Health:   "healthy",
		LastSeen: time.Now(),
	}
}

// Run executes a task with the runner
func (r *Runner) Run(ctx context.Context, taskID string) error {
	fmt.Printf("Runner %s starting task %s\n", r.ID, taskID)
	
	// Simulate task execution
	select {
	case <-ctx.Done():
		fmt.Printf("Runner %s cancelled task %s\n", r.ID, taskID)
		return ctx.Err()
	case <-time.After(2 * time.Second):
		fmt.Printf("Runner %s completed task %s\n", r.ID, taskID)
		return nil
	}
}

// HealthCheck checks the runner's health status
func (r *Runner) HealthCheck() string {
	// Simulate health check logic
	if time.Since(r.LastSeen) > 5*time.Minute {
		r.Health = "unhealthy"
	} else {
		r.Health = "healthy"
	}
	return r.Health
}

// UpdateLastSeen updates the last seen timestamp
func (r *Runner) UpdateLastSeen() {
	r.LastSeen = time.Now()
}

// GetStats returns runner statistics
func (r *Runner) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"id":           r.ID,
		"name":         r.Name,
		"health":       r.Health,
		"last_seen":    r.LastSeen,
		"active_tasks": r.ActiveTasks,
	}
}