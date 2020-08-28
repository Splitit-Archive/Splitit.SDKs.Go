package splitit

import "context"

func WithCulture(ctx context.Context, culture string) context.Context {
	return context.WithValue(ctx, cultureCtxKey{}, culture)
}

type cultureCtxKey struct{}
type noApiKeyCtxKey struct{}
type noSessionIdKey struct{}
