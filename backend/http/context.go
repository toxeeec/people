package http

import "context"

type contextKey int

const (
	userIDKey contextKey = iota
)

func newContext(ctx context.Context, key contextKey, value uint) context.Context {
	return context.WithValue(ctx, key, value)
}

func fromContext(ctx context.Context, key contextKey) (uint, bool) {
	v, ok := ctx.Value(key).(uint)
	return v, ok
}
