package logger

import (
	"context"
)

var (
	ctxKey = ctxKeyType{}
)

type ctxKeyType struct{}

func ToContext(ctx context.Context, logger *logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}

func FromContext(ctx context.Context) *logger {
	v := ctx.Value(ctxKey)
	l, ok := v.(*logger)
	if !ok {
		return Default
	}

	return l
}
