package ports

import "context"

type FeatureFragmentProvider interface {
	IsEnabled(ctx context.Context, n string, opts ...interface{}) bool
	IsEnabledNumber(ctx context.Context, n string, opts ...interface{}) float64
	IsEnabledString(ctx context.Context, n string, opts ...interface{}) string
	IsEnabledByStrategy(ctx context.Context, n string, opts ...interface{}) interface{}
	// IsJSONEnabled(ctx context.Context, n string, opts ...interface{}) []byte
	// IsEnabledString(ctx context.Context, n string, opts ...interface{}) string
}
