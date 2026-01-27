package context

import (
	"context"
	"time"
)

// ContextKey represents a key for context values
type ContextKey string

const (
	// ContextKeyCyborgID is the context key for cyborg ID
	ContextKeyCyborgID ContextKey = "cyborg_id"
	
	// ContextKeyJobID is the context key for job ID
	ContextKeyJobID ContextKey = "job_id"
	
	// ContextKeyTaskID is the context key for task ID
	ContextKeyTaskID ContextKey = "task_id"
	
	// ContextKeyRequestID is the context key for request ID
	ContextKeyRequestID ContextKey = "request_id"
)

// ContextWithCyborgID adds cyborg ID to context
func ContextWithCyborgID(ctx context.Context, cyborgID string) context.Context {
	return context.WithValue(ctx, ContextKeyCyborgID, cyborgID)
}

// ContextWithJobID adds job ID to context
func ContextWithJobID(ctx context.Context, jobID string) context.Context {
	return context.WithValue(ctx, ContextKeyJobID, jobID)
}

// ContextWithTaskID adds task ID to context
func ContextWithTaskID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, ContextKeyTaskID, taskID)
}

// ContextWithRequestID adds request ID to context
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

// CyborgIDFromContext extracts cyborg ID from context
func CyborgIDFromContext(ctx context.Context) string {
	if v := ctx.Value(ContextKeyCyborgID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// JobIDFromContext extracts job ID from context
func JobIDFromContext(ctx context.Context) string {
	if v := ctx.Value(ContextKeyJobID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// TaskIDFromContext extracts task ID from context
func TaskIDFromContext(ctx context.Context) string {
	if v := ctx.Value(ContextKeyTaskID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// RequestIDFromContext extracts request ID from context
func RequestIDFromContext(ctx context.Context) string {
	if v := ctx.Value(ContextKeyRequestID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// ContextWithTimeout creates a context with timeout
func ContextWithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

// ContextWithDeadline creates a context with deadline
func ContextWithDeadline(ctx context.Context, deadline time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(ctx, deadline)
}

// ContextWithCancel creates a cancellable context
func ContextWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}