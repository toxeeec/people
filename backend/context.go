package people

import "context"

type ContextKey int

const (
	UserIDKey ContextKey = iota
)

func NewContext(ctx context.Context, key ContextKey, value uint) context.Context {
	return context.WithValue(ctx, key, value)
}

func FromContext(ctx context.Context, key ContextKey) (uint, bool) {
	v, ok := ctx.Value(key).(uint)
	return v, ok
}
