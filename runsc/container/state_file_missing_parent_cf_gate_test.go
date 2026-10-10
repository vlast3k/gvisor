package container

import (
	"errors"
	"os"
	"testing"
)

func TestCFGateMissingParentWithChildrenIsNotFound(t *testing.T) {
	rootDir := t.TempDir()
	const parentID = "garden-parent"
	writeCFGateStateFile(t, rootDir, parentID, parentID+"-envoy-startup-healthcheck-0")
	writeCFGateStateFile(t, rootDir, parentID, parentID+"-envoy-startup-healthcheck-1")

	if _, err := findContainerID(rootDir, parentID); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("finding absent parent returned %v, want os.ErrNotExist", err)
	}
}
