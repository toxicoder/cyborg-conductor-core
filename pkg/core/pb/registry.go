package pb

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"github.com/toxicoder/cyborg-conductor-core/pkg/core/types"
	"google.golang.org/protobuf/proto"
)

// Registry manages all registered cyborgs with their capabilities and configurations
type Registry struct {
	mu    sync.RWMutex
	items map[string]*types.CyborgDescriptor
}

// NewRegistry creates a new thread-safe cyborg registry
func NewRegistry() *Registry {
	return &Registry{items: make(map[string]*types.CyborgDescriptor)}
}

// Register adds a cyborg descriptor to the registry
func (r *Registry) Register(descriptor *types.CyborgDescriptor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if descriptor == nil {
		return fmt.Errorf("cannot register nil cyborg descriptor")
	}
	
	if _, exists := r.items[descriptor.GetCyborgId()]; exists {
		return fmt.Errorf("cyborg %s already registered", descriptor.GetCyborgId())
	}
	
	r.items[descriptor.GetCyborgId()] = descriptor
	return nil
}

// Get retrieves a cyborg descriptor by ID
func (r *Registry) Get(id string) (*types.CyborgDescriptor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	descriptor, exists := r.items[id]
	return descriptor, exists
}

// List returns all registered cyborg descriptors
func (r *Registry) List() []*types.CyborgDescriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	result := make([]*types.CyborgDescriptor, 0, len(r.items))
	for _, descriptor := range r.items {
		result = append(result, descriptor)
	}
	return result
}

// LoadFromTxtpb loads all cyborg descriptors from .txtpb files in the specified directory
func (r *Registry) LoadFromTxtpb(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dir, err)
	}
	
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".txtpb" {
			continue
		}
		
		path := filepath.Join(dir, f.Name())
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}
		
		var descriptor types.CyborgDescriptor
		if err := proto.UnmarshalText(string(data), &descriptor); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", path, err)
		}
		
		if err := r.Register(&descriptor); err != nil {
			return fmt.Errorf("failed to register cyborg from %s: %w", path, err)
		}
	}
	
	return nil
}

// Size returns the number of registered cyborgs
func (r *Registry) Size() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.items)
}