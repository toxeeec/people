package people

import "context"

type contextKey int

const (
	UserIDKey contextKey = iota // uint
)

func NewContext(ctx context.Context, key contextKey, value uint) context.Context {
	return context.WithValue(ctx, key, value)
}

func FromContext(ctx context.Context, key contextKey) (uint, bool) {
	v, ok := ctx.Value(key).(uint)
	return v, ok
}
