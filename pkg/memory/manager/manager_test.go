package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryManager(t *testing.T) {
	// Test creating a new memory manager
	manager, err := New("/tmp/test")
	assert.NoError(t, err)
	assert.NotNil(t, manager)
	assert.Equal(t, "/tmp/test", manager.root)
	assert.NotNil(t, manager.cache)
}

func TestMemoryManagerApplyPolicy(t *testing.T) {
	// Test applying policy
	manager, err := New("/tmp/test")
	assert.NoError(t, err)
	assert.NotNil(t, manager)

	// Test with empty policy (should not error)
	err = manager.ApplyPolicy(Config{})
	assert.NoError(t, err)
}

func TestMemoryManagerEvidence(t *testing.T) {
	// Test getting evidence
	manager, err := New("/tmp/test")
	assert.NoError(t, err)
	assert.NotNil(t, manager)

	// Test that evidence returns nil (since it's a placeholder)
	evidence := manager.Evidence()
	assert.Nil(t, evidence)
}