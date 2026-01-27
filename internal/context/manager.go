package context

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/simplelru"
	"github.com/toxicoder/cyborg-conductor-core/pkg/proto/generated"
)

// MemoryCacheManager enforces context tolerance policy and manages memory
type MemoryCacheManager struct {
	cache        *simplelru.LRU[string, []byte]
	maxSizeBytes int64
	mu           sync.RWMutex
	compression  bool
}

// NewMemoryCacheManager creates a new memory cache manager
func NewMemoryCacheManager(maxSizeBytes int64) (*MemoryCacheManager, error) {
	// Create LRU cache with a reasonable size (1000 items)
	cache, err := simplelru.NewLRU[string, []byte](1000, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create LRU cache: %w", err)
	}

	return &MemoryCacheManager{
		cache:        cache,
		maxSizeBytes: maxSizeBytes,
		compression:  true, // Enable compression by default
	}, nil
}

// ApplyPolicy applies the context tolerance policy
func (m *MemoryCacheManager) ApplyPolicy(ctx context.Context, policy interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// In a real implementation, this would:
	// 1. Check the policy against current cache state
	// 2. Trim or replicate cache entries as needed
	// 3. Enforce the policy constraints
	
	// For now, just validate that we're within limits
	currentSize := m.getCurrentSize()
	if currentSize > m.maxSizeBytes {
		// Trim cache to make room
		m.trimCache()
	}

	return nil
}

// Set sets a key-value pair in the cache
func (m *MemoryCacheManager) Set(ctx context.Context, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if adding this would exceed limits
	newSize := m.getCurrentSize() + int64(len(value))
	if newSize > m.maxSizeBytes {
		// Trim cache before adding
		m.trimCache()
	}

	// Add to cache
	m.cache.Add(key, value)
	return nil
}

// Get retrieves a value from the cache
func (m *MemoryCacheManager) Get(ctx context.Context, key string) ([]byte, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.cache.Get(key)
	return value, exists
}

// Delete removes a key from the cache
func (m *MemoryCacheManager) Delete(ctx context.Context, key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Remove(key)
}

// Clear clears all entries from the cache
func (m *MemoryCacheManager) Clear(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Purge()
}

// Size returns the current cache size
func (m *MemoryCacheManager) Size(ctx context.Context) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.Len()
}

// getCurrentSize returns the current size of all cached data in bytes
func (m *MemoryCacheManager) getCurrentSize() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var totalSize int64
	for _, key := range m.cache.Keys() {
		if value, exists := m.cache.Get(key); exists {
			totalSize += int64(len(value))
		}
	}
	return totalSize
}

// trimCache trims the cache to make room for new entries
func (m *MemoryCacheManager) trimCache() {
	// Remove oldest entries until we're under the size limit
	// This is a simple implementation - a more sophisticated one might
	// prioritize based on recency, importance, etc.
	
	currentSize := m.getCurrentSize()
	for currentSize > m.maxSizeBytes/2 && m.cache.Len() > 0 {
		// Remove the oldest entry
		oldestKey := m.cache.Keys()[0]
		m.cache.Remove(oldestKey)
		
		// Recalculate size
		currentSize = m.getCurrentSize()
	}
}

// Compress compresses data using LZ4 when size exceeds threshold
func (m *MemoryCacheManager) Compress(ctx context.Context, data []byte, threshold int) ([]byte, error) {
	if !m.compression || len(data) <= threshold {
		return data, nil
	}

	// In a real implementation, this would use LZ4 compression
	// For now, just return the data as-is
	return data, nil
}

// GetStats returns cache statistics
func (m *MemoryCacheManager) GetStats(ctx context.Context) map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	stats := make(map[string]interface{})
	stats["size"] = m.cache.Len()
	stats["max_size_bytes"] = m.maxSizeBytes
	stats["current_size_bytes"] = m.getCurrentSize()
	stats["compression_enabled"] = m.compression
	
	return stats
}

// IsFull checks if the cache is at maximum capacity
func (m *MemoryCacheManager) IsFull(ctx context.Context) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.cache.Len() >= m.cache.MaxSize()
}

// SetCompression enables or disables compression
func (m *MemoryCacheManager) SetCompression(ctx context.Context, enabled bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.compression = enabled
}

// SetMaxSize sets the maximum size for the cache
func (m *MemoryCacheManager) SetMaxSize(ctx context.Context, maxSizeBytes int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.maxSizeBytes = maxSizeBytes
}

// EvictOldest evicts the oldest entry from cache
func (m *MemoryCacheManager) EvictOldest(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.cache.Len() > 0 {
		oldestKey := m.cache.Keys()[0]
		m.cache.Remove(oldestKey)
	}
}

// EvictAll evicts all entries from cache
func (m *MemoryCacheManager) EvictAll(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Purge()
}