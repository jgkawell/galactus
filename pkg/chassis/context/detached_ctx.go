package context

import (
	"context"
	"time"
)

type detachedContext struct {
	parent context.Context
}

// Detach returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling.
func Detach(ctx context.Context) context.Context {
	return detachedContext{ctx}
}

func (v detachedContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (v detachedContext) Done() <-chan struct{} {
	return nil
}

func (v detachedContext) Err() error {
	return nil
}

func (v detachedContext) Value(key interface{}) interface{} {
	return v.parent.Value(key)
}
