package core

import (
	"context"
	"time"
)

// Based on: https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	Context context.Context

	TraceID    string
	StartTime  time.Time
	StatusCode int

	CallerID uint64
	ClientID uint64
}

const (
	ContextKey string = "app-context"
)

// Alias.
type Context = Values

func (c Context) AsTags() []string {
	return []string{
		"trace_id:" + c.TraceID,
	}
}
