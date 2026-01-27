package context

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContextKeyConstants(t *testing.T) {
	assert.Equal(t, ContextKey("cyborg_id"), ContextKeyCyborgID)
	assert.Equal(t, ContextKey("job_id"), ContextKeyJobID)
	assert.Equal(t, ContextKey("task_id"), ContextKeyTaskID)
	assert.Equal(t, ContextKey("request_id"), ContextKeyRequestID)
}

func TestContextWithCyborgID(t *testing.T) {
	ctx := context.Background()
	result := ContextWithCyborgID(ctx, "test-cyborg-id")
	
	cyborgID := CyborgIDFromContext(result)
	assert.Equal(t, "test-cyborg-id", cyborgID)
}

func TestContextWithJobID(t *testing.T) {
	ctx := context.Background()
	result := ContextWithJobID(ctx, "test-job-id")
	
	jobID := JobIDFromContext(result)
	assert.Equal(t, "test-job-id", jobID)
}

func TestContextWithTaskID(t *testing.T) {
	ctx := context.Background()
	result := ContextWithTaskID(ctx, "test-task-id")
	
	taskID := TaskIDFromContext(result)
	assert.Equal(t, "test-task-id", taskID)
}

func TestContextWithRequestID(t *testing.T) {
	ctx := context.Background()
	result := ContextWithRequestID(ctx, "test-request-id")
	
	requestID := RequestIDFromContext(result)
	assert.Equal(t, "test-request-id", requestID)
}

func TestContextWithTimeout(t *testing.T) {
	ctx := context.Background()
	result, cancel := ContextWithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	assert.NotNil(t, result)
	assert.NotNil(t, cancel)
}

func TestContextWithDeadline(t *testing.T) {
	ctx := context.Background()
	deadline := time.Now().Add(5 * time.Second)
	result, cancel := ContextWithDeadline(ctx, deadline)
	defer cancel()
	
	assert.NotNil(t, result)
	assert.NotNil(t, cancel)
}

func TestContextWithCancel(t *testing.T) {
	ctx := context.Background()
	result, cancel := ContextWithCancel(ctx)
	defer cancel()
	
	assert.NotNil(t, result)
	assert.NotNil(t, cancel)
}

func TestCyborgIDFromContext(t *testing.T) {
	// Test with context that has cyborg ID
	ctx := ContextWithCyborgID(context.Background(), "test-cyborg")
	result := CyborgIDFromContext(ctx)
	assert.Equal(t, "test-cyborg", result)
	
	// Test with empty context
	emptyCtx := context.Background()
	result = CyborgIDFromContext(emptyCtx)
	assert.Equal(t, "", result)
	
	// Test with context that has wrong type value
	wrongCtx := context.WithValue(context.Background(), ContextKeyCyborgID, 123)
	result = CyborgIDFromContext(wrongCtx)
	assert.Equal(t, "", result)
}

func TestJobIDFromContext(t *testing.T) {
	// Test with context that has job ID
	ctx := ContextWithJobID(context.Background(), "test-job")
	result := JobIDFromContext(ctx)
	assert.Equal(t, "test-job", result)
	
	// Test with empty context
	emptyCtx := context.Background()
	result = JobIDFromContext(emptyCtx)
	assert.Equal(t, "", result)
	
	// Test with context that has wrong type value
	wrongCtx := context.WithValue(context.Background(), ContextKeyJobID, 123)
	result = JobIDFromContext(wrongCtx)
	assert.Equal(t, "", result)
}

func TestTaskIDFromContext(t *testing.T) {
	// Test with context that has task ID
	ctx := ContextWithTaskID(context.Background(), "test-task")
	result := TaskIDFromContext(ctx)
	assert.Equal(t, "test-task", result)
	
	// Test with empty context
	emptyCtx := context.Background()
	result = TaskIDFromContext(emptyCtx)
	assert.Equal(t, "", result)
	
	// Test with context that has wrong type value
	wrongCtx := context.WithValue(context.Background(), ContextKeyTaskID, 123)
	result = TaskIDFromContext(wrongCtx)
	assert.Equal(t, "", result)
}

func TestRequestIDFromContext(t *testing.T) {
	// Test with context that has request ID
	ctx := ContextWithRequestID(context.Background(), "test-request")
	result := RequestIDFromContext(ctx)
	assert.Equal(t, "test-request", result)
	
	// Test with empty context
	emptyCtx := context.Background()
	result = RequestIDFromContext(emptyCtx)
	assert.Equal(t, "", result)
	
	// Test with context that has wrong type value
	wrongCtx := context.WithValue(context.Background(), ContextKeyRequestID, 123)
	result = RequestIDFromContext(wrongCtx)
	assert.Equal(t, "", result)
}