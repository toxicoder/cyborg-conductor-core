package context

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/toxicoder/cyborg-conductor-core/pkg/proto/pb"
)

// ContextOverlayEngine loads immutable evidence snapshots and presents them as read-only buffers
type ContextOverlayEngine struct {
	evidenceRoot string
}

// NewContextOverlayEngine creates a new context overlay engine
func NewContextOverlayEngine(evidenceRoot string) *ContextOverlayEngine {
	return &ContextOverlayEngine{
		evidenceRoot: evidenceRoot,
	}
}

// GetSnapshot loads an immutable evidence snapshot for a cyborg
func (c *ContextOverlayEngine) GetSnapshot(ctx context.Context, cyborgID string) ([]byte, error) {
	// Construct the path to the evidence file
	evidenceFile := filepath.Join(c.evidenceRoot, fmt.Sprintf("%s.merklelog.bin", cyborgID))
	
	// Check if file exists
	_, err := os.Stat(evidenceFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("evidence file not found for cyborg %s: %w", cyborgID, err)
	} else if err != nil {
		return nil, fmt.Errorf("error checking evidence file for cyborg %s: %w", cyborgID, err)
	}
	
	// Open the file in read-only mode
	file, err := os.Open(evidenceFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open evidence file for cyborg %s: %w", cyborgID, err)
	}
	defer file.Close()
	
	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read evidence file for cyborg %s: %w", cyborgID, err)
	}
	
	return content, nil
}

// GetSnapshotWithBuffer loads an immutable evidence snapshot with a specific buffer size
func (c *ContextOverlayEngine) GetSnapshotWithBuffer(ctx context.Context, cyborgID string, bufferSize int) ([]byte, error) {
	// Construct the path to the evidence file
	evidenceFile := filepath.Join(c.evidenceRoot, fmt.Sprintf("%s.merklelog.bin", cyborgID))
	
	// Check if file exists
	_, err := os.Stat(evidenceFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("evidence file not found for cyborg %s: %w", cyborgID, err)
	} else if err != nil {
		return nil, fmt.Errorf("error checking evidence file for cyborg %s: %w", cyborgID, err)
	}
	
	// Open the file in read-only mode
	file, err := os.Open(evidenceFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open evidence file for cyborg %s: %w", cyborgID, err)
	}
	defer file.Close()
	
	// Create buffer of specified size
	buffer := make([]byte, bufferSize)
	
	// Read up to bufferSize bytes
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read evidence file for cyborg %s: %w", cyborgID, err)
	}
	
	// Return the actual bytes read
	return buffer[:n], nil
}

// ValidateEvidencePath checks if the evidence path is valid
func (c *ContextOverlayEngine) ValidateEvidencePath(ctx context.Context, cyborgID string) (bool, error) {
	evidenceFile := filepath.Join(c.evidenceRoot, fmt.Sprintf("%s.merklelog.bin", cyborgID))
	
	_, err := os.Stat(evidenceFile)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error validating evidence path for cyborg %s: %w", cyborgID, err)
	}
	
	return true, nil
}

// ListAvailableSnapshots returns a list of available evidence snapshots
func (c *ContextOverlayEngine) ListAvailableSnapshots(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(c.evidenceRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to read evidence root directory: %w", err)
	}
	
	var snapshots []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".merklelog.bin" {
			// Remove the extension to get just the cyborg ID
			cyborgID := entry.Name()[:len(entry.Name())-len(".merklelog.bin")]
			snapshots = append(snapshots, cyborgID)
		}
	}
	
	return snapshots, nil
}