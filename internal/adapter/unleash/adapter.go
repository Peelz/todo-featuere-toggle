package unleash

import (
	"context"
	"github.com/Unleash/unleash-client-go/v3"
)

type Adapter struct {
	client *unleash.Client
}

func (a Adapter) IsEnabled(ctx context.Context, n string, opts ...interface{}) bool {
	return a.client.IsEnabled("n")
}

func NewAdapter(client *unleash.Client) *Adapter {
	return &Adapter{client: client}
}
