package container

import (
	"os"
	"path/filepath"
	"testing"
)

func writeCFGateStateFile(t *testing.T, rootDir, sandboxID, containerID string) {
	t.Helper()
	path := buildPath(rootDir, FullID{SandboxID: sandboxID, ContainerID: containerID}, stateFileExtension)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("creating state directory: %v", err)
	}
	if err := os.WriteFile(path, nil, 0o600); err != nil {
		t.Fatalf("creating state file: %v", err)
	}
}

func TestCFGateExactParentPreferredOverChildren(t *testing.T) {
	rootDir := t.TempDir()
	const parentID = "garden-parent"
	writeCFGateStateFile(t, rootDir, parentID, parentID)
	writeCFGateStateFile(t, rootDir, parentID, parentID+"-envoy-startup-healthcheck-0")
	writeCFGateStateFile(t, rootDir, parentID, parentID+"-envoy-startup-healthcheck-1")

	got, err := findContainerID(rootDir, parentID)
	if err != nil {
		t.Fatalf("finding exact parent ID: %v", err)
	}
	if got.ContainerID != parentID {
		t.Fatalf("resolved container ID %q, want %q", got.ContainerID, parentID)
	}
}
