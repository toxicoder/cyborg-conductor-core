package pb

import (
	"fmt"
	"sync"
)

// Registry stores cyborg descriptors by ID
type Registry struct {
	mu    sync.RWMutex
	items map[string]*CyborgDescriptor
}

// NewRegistry creates a new registry
func NewRegistry() *Registry {
	return &Registry{
		items: make(map[string]*CyborgDescriptor),
	}
}

// Register adds a cyborg descriptor to the registry
func (r *Registry) Register(descriptor *CyborgDescriptor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if descriptor == nil {
		return fmt.Errorf("descriptor cannot be nil")
	}
	
	if descriptor.CyborgID == "" {
		return fmt.Errorf("cyborg ID cannot be empty")
	}
	
	r.items[descriptor.CyborgID] = descriptor
	return nil
}

// Get retrieves a cyborg descriptor by ID
func (r *Registry) Get(id string) (*CyborgDescriptor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	descriptor, exists := r.items[id]
	return descriptor, exists
}

// List returns all cyborg descriptors
func (r *Registry) List() []*CyborgDescriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	result := make([]*CyborgDescriptor, 0, len(r.items))
	for _, descriptor := range r.items {
		result = append(result, descriptor)
	}
	return result
}