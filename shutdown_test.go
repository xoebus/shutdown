package shutdown

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	ctx := WithShutdown(context.Background())

	pid := os.Getpid()
	p, err := os.FindProcess(pid)
	if err != nil {
		t.Fatalf("failed to find process (pid: %d): %v", pid, err)
	}

	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatalf("failed to signal process: %v", err)
	}

	select {
	case <-ctx.Done():
		break
	case <-time.After(time.Second):
		t.Fatal("context was not cancelled before timeout")
	}
}
