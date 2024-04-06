package rt

import (
	"context"
	"time"
)

// How long should we wait for the SLOW statement ?
var SlowDuration time.Duration = 1 * time.Second

// Wait for SlowDuration or for the context to be cancelled.
func Slow(ctx context.Context) {
	select {
	case <-time.After(SlowDuration):
	case <-ctx.Done():
	}
}
